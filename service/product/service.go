package Product

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/techno/entity"
	"github.com/techno/entity/apperror"
	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	"github.com/uptrace/bun"

	productStockService "github.com/techno/service/productstock"
)

var (
	DuplicateProductName = apperror.New(http.StatusBadRequest, "duplicate.Product.name", "already exists. Please choose a different name.")
	ProductNotFound      = apperror.New(http.StatusNotFound, "Product.not.found", "not found.")
	QuantityError        = apperror.New(http.StatusNotFound, "please update your quantity", "please update your quantity")
)

type Service struct {
	Repository          Repository
	ProductStockService productStockService.Service
	Logger              logger.Logger
}

func NewService(repository Repository, logger logger.Logger, productStockService productStockService.Service) *Service {
	return &Service{
		Repository:          repository,
		ProductStockService: productStockService,
		Logger:              logger,
	}
}

type Repository interface {
	Update(ctx context.Context, entityRef *entity.Product) (int64, error)
	GetByID(ctx context.Context, ProductId int) (entity.Product, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.ProductFilter) (int, []entity.Product, error)
	Create(ctx context.Context, entitys *entity.Product, tx *bun.Tx) error
	Delete(ctx context.Context, id int) (int64, error)
	GetTx(ctx context.Context) (*bun.Tx, error)
}

// Get entity
func (s *Service) List(ctx context.Context, params httpentity.ProductParams) (*httpentity.ProductList, error) {
	pagination := params.PaginationRequest.GetLimitOffset()
	dbFilter := entity.ProductFilter{}

	if params.StatusId != 0 {
		var active bool
		if params.StatusId == 1 {
			active = true
		}
		if params.StatusId == 2 {
			active = false
		}
		dbFilter.StatusId = &active

	}
	dbFilter.Name = params.Name
	dbFilter.MinPrice = params.MinPrice
	dbFilter.MaxPrice = params.MaxPrice
	dbFilter.CategoryId = params.CategoryId
	if params.BrandId != "" {
		dbFilter.BrandId = strings.Split(params.BrandId, ",")
	}

	dbFilter.SupplierId = params.SupplierId

	total, Product, err := s.Repository.List(ctx, pagination, dbFilter)
	if err != nil {
		return &httpentity.ProductList{}, apperror.InteralError.Wrap(err)
	}
	return &httpentity.ProductList{
		PaginationResponse: httpentity.NewPaginationResponse(total, pagination.Limit),
		Product:            Product,
	}, nil
}

// Update entity item
func (s *Service) Update(ctx context.Context, data httpentity.UpdateProductRequest) (*entity.Product, error) {
	Product := entity.Product{
		Id:             data.Id,
		Name:           data.Name,
		Description:    data.Description,
		Specifications: data.Specifications,
		BrandId:        data.BrandId,
		CategoryId:     data.CategoryId,
		SupplierId:     data.SupplierId,
		UnitPrice:      data.UnitPrice,
		DiscountPrice:  data.DiscountPrice,
		Tags:           data.Tags,
	}
	affected, err := s.Repository.Update(ctx, &Product)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return nil, ProductNotFound
	}
	return &Product, nil
}

// Get entity by id
func (s *Service) GetByID(ctx context.Context, ProductId int) (entity.Product, error) {
	Product, err := s.Repository.GetByID(ctx, ProductId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Product, ProductNotFound
		}
		return Product, apperror.InteralError.Wrap(err)
	}

	return Product, nil
}

// Create entity Bulk
func (s *Service) Create(ctx context.Context, data httpentity.CreateProductRequest) (*entity.Product, error) {

	tx, err := s.Repository.GetTx(ctx)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	defer tx.Rollback()

	Product := entity.Product{}

	Product = entity.Product{
		Name:           data.Name,
		Description:    data.Description,
		Specifications: data.Specifications,
		BrandId:        data.BrandId,
		CategoryId:     data.CategoryId,
		SupplierId:     data.SupplierId,
		UnitPrice:      data.UnitPrice,
		DiscountPrice:  data.DiscountPrice,
		Tags:           data.Tags,
	}

	err = s.Repository.Create(ctx, &Product, tx)

	if err != nil {

		return nil, apperror.InteralError.Wrap(err)
	}

	if data.ProductStock.StockQuantity <= 0 {
		return nil, QuantityError

	}

	//used transaction
	err = s.ProductStockService.Repository.Create(ctx, &entity.ProductStock{
		ProductId:     Product.Id,
		StockQuantity: data.ProductStock.StockQuantity,
	}, tx)

	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}

	return &Product, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	affected, err := s.Repository.Delete(ctx, id)
	if err != nil {
		return apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return apperror.InteralError.Wrap(err)
	}
	return nil
}
