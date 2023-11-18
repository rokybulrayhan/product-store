package Product

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-contact-service/entity"
	"github.com/go-contact-service/entity/apperror"
	"github.com/go-contact-service/entity/httpentity"
	"github.com/go-contact-service/lib/logger"
)

var (
	DuplicateProductName = apperror.New(http.StatusBadRequest, "duplicate.Product.name", "contact group already exists. Please choose a different contact group name.")
	ProductNotFound      = apperror.New(http.StatusNotFound, "Product.not.found", "contact group not found.")
)

type Service struct {
	Repository Repository
	Logger     logger.Logger
}

func NewService(repository Repository, logger logger.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

type Repository interface {
	Update(ctx context.Context, entityRef *entity.Product) (int64, error)
	GetByID(ctx context.Context, ProductId int) (entity.Product, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.ProductFilter) (int, []entity.Product, error)
	Create(ctx context.Context, entitys *entity.Product) error
	Delete(ctx context.Context, id int) (int64, error)
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

	err := s.Repository.Create(ctx, &Product)

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
