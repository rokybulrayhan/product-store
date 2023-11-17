package Brand

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
	DuplicateBrandName = apperror.New(http.StatusBadRequest, "duplicate.Brand.name", "contact group already exists. Please choose a different contact group name.")
	BrandNotFound      = apperror.New(http.StatusNotFound, "Brand.not.found", "contact group not found.")
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
	Update(ctx context.Context, entityRef *entity.Brand) (int64, error)
	GetByID(ctx context.Context, BrandId int) (entity.Brand, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.BrandFilter) (int, []entity.Brand, error)
	Create(ctx context.Context, entitys *entity.Brand) error
}

// Get entity
func (s *Service) List(ctx context.Context, params httpentity.BrandParams) (*httpentity.BrandList, error) {
	pagination := params.PaginationRequest.GetLimitOffset()
	dbFilter := entity.BrandFilter{}

	total, Brand, err := s.Repository.List(ctx, pagination, dbFilter)
	if err != nil {
		return &httpentity.BrandList{}, apperror.InteralError.Wrap(err)
	}
	return &httpentity.BrandList{
		PaginationResponse: httpentity.NewPaginationResponse(total, pagination.Limit),
		Brand:              Brand,
	}, nil
}

// Update entity item
func (s *Service) Update(ctx context.Context, data httpentity.UpdateBrandRequest) (*entity.Brand, error) {
	Brand := entity.Brand{
		Id:       data.Id,
		Name:     data.Name,
		StatusId: data.StatusId,
	}
	affected, err := s.Repository.Update(ctx, &Brand)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return nil, BrandNotFound
	}
	return &Brand, nil
}

// Get entity by id
func (s *Service) GetByID(ctx context.Context, BrandId int) (entity.Brand, error) {
	Brand, err := s.Repository.GetByID(ctx, BrandId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Brand, BrandNotFound
		}
		return Brand, apperror.InteralError.Wrap(err)
	}

	return Brand, nil
}

// Create entity Bulk
func (s *Service) Create(ctx context.Context, data httpentity.CreateBrandRequest) (*entity.Brand, error) {
	Brand := entity.Brand{}

	Brand = entity.Brand{
		Name: data.Name,
	}

	err := s.Repository.Create(ctx, &Brand)

	if err != nil {

		return nil, apperror.InteralError.Wrap(err)
	}
	return &Brand, nil
}
