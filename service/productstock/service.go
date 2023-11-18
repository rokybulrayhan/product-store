package ProductStock

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
	DuplicateProductStockName = apperror.New(http.StatusBadRequest, "duplicate.ProductStock.name", "contact group already exists. Please choose a different contact group name.")
	ProductStockNotFound      = apperror.New(http.StatusNotFound, "ProductStock.not.found", "contact group not found.")
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
	Update(ctx context.Context, entityRef *entity.ProductStock) (int64, error)
	GetByID(ctx context.Context, ProductStockId int) (entity.ProductStock, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.ProductStockFilter) (int, []entity.ProductStock, error)
	Create(ctx context.Context, entitys *entity.ProductStock) error
	Delete(ctx context.Context, id int) (int64, error)
}

// Get entity
func (s *Service) List(ctx context.Context, params httpentity.ProductStockParams) (*httpentity.ProductStockList, error) {
	pagination := params.PaginationRequest.GetLimitOffset()
	dbFilter := entity.ProductStockFilter{}
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

	total, ProductStock, err := s.Repository.List(ctx, pagination, dbFilter)
	if err != nil {
		return &httpentity.ProductStockList{}, apperror.InteralError.Wrap(err)
	}
	return &httpentity.ProductStockList{
		PaginationResponse: httpentity.NewPaginationResponse(total, pagination.Limit),
		ProductStock:       ProductStock,
	}, nil
}

// Update entity item
func (s *Service) Update(ctx context.Context, data httpentity.UpdateProductStockRequest) (*entity.ProductStock, error) {
	ProductStock := entity.ProductStock{
		Id:            data.Id,
		ProductId:     data.ProductId,
		StockQuantity: data.StockQuantity,
	}
	affected, err := s.Repository.Update(ctx, &ProductStock)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return nil, ProductStockNotFound
	}
	return &ProductStock, nil
}

// Get entity by id
func (s *Service) GetByID(ctx context.Context, ProductStockId int) (entity.ProductStock, error) {
	ProductStock, err := s.Repository.GetByID(ctx, ProductStockId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductStock, ProductStockNotFound
		}
		return ProductStock, apperror.InteralError.Wrap(err)
	}

	return ProductStock, nil
}

// Create entity Bulk
func (s *Service) Create(ctx context.Context, data httpentity.CreateProductStockRequest) (*entity.ProductStock, error) {
	ProductStock := entity.ProductStock{}

	ProductStock = entity.ProductStock{
		ProductId:     data.ProductId,
		StockQuantity: data.StockQuantity,
	}

	err := s.Repository.Create(ctx, &ProductStock)

	if err != nil {

		return nil, apperror.InteralError.Wrap(err)
	}
	return &ProductStock, nil
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
