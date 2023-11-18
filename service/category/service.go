package Category

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
	DuplicateCategoryName = apperror.New(http.StatusBadRequest, "duplicate.Category.name", "contact group already exists. Please choose a different contact group name.")
	CategoryNotFound      = apperror.New(http.StatusNotFound, "Category.not.found", "contact group not found.")
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
	Update(ctx context.Context, entityRef *entity.Category) (int64, error)
	GetByID(ctx context.Context, CategoryId int) (entity.Category, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.CategoryFilter) (int, []entity.Category, error)
	Create(ctx context.Context, entitys *entity.Category) error
	Delete(ctx context.Context, id int) (int64, error)
}

// Get entity
func (s *Service) List(ctx context.Context, params httpentity.CategoryParams) (*httpentity.CategoryList, error) {
	pagination := params.PaginationRequest.GetLimitOffset()
	dbFilter := entity.CategoryFilter{}

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

	total, Category, err := s.Repository.List(ctx, pagination, dbFilter)
	if err != nil {
		return &httpentity.CategoryList{}, apperror.InteralError.Wrap(err)
	}
	return &httpentity.CategoryList{
		PaginationResponse: httpentity.NewPaginationResponse(total, pagination.Limit),
		Category:           Category,
	}, nil
}

// Update entity item
func (s *Service) Update(ctx context.Context, data httpentity.UpdateCategoryRequest) (*entity.Category, error) {
	Category := entity.Category{
		Id:       data.Id,
		Name:     data.Name,
		ParentId: data.ParentId,
		StatusId: data.StatusId,
	}
	affected, err := s.Repository.Update(ctx, &Category)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return nil, CategoryNotFound
	}
	return &Category, nil
}

// Get entity by id
func (s *Service) GetByID(ctx context.Context, CategoryId int) (entity.Category, error) {
	Category, err := s.Repository.GetByID(ctx, CategoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Category, CategoryNotFound
		}
		return Category, apperror.InteralError.Wrap(err)
	}

	return Category, nil
}

// Create entity Bulk
func (s *Service) Create(ctx context.Context, data httpentity.CreateCategoryRequest) (*entity.Category, error) {
	Category := entity.Category{}

	Category = entity.Category{
		Name:     data.Name,
		ParentId: data.ParentId,
	}

	err := s.Repository.Create(ctx, &Category)

	if err != nil {

		return nil, apperror.InteralError.Wrap(err)
	}
	return &Category, nil
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
