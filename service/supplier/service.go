package Supplier

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
	DuplicateSupplierName = apperror.New(http.StatusBadRequest, "duplicate.Supplier.name", "contact group already exists. Please choose a different contact group name.")
	SupplierNotFound      = apperror.New(http.StatusNotFound, "Supplier.not.found", "contact group not found.")
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
	Update(ctx context.Context, entityRef *entity.Supplier) (int64, error)
	GetByID(ctx context.Context, SupplierId int) (entity.Supplier, error)
	List(ctx context.Context, pagination entity.Pagination, filter entity.SupplierFilter) (int, []entity.Supplier, error)
	Create(ctx context.Context, entitys *entity.Supplier) error
	Delete(ctx context.Context, id int) (int64, error)
}

// Get entity
func (s *Service) List(ctx context.Context, params httpentity.SupplierParams) (*httpentity.SupplierList, error) {
	pagination := params.PaginationRequest.GetLimitOffset()
	dbFilter := entity.SupplierFilter{}
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

	total, Supplier, err := s.Repository.List(ctx, pagination, dbFilter)
	if err != nil {
		return &httpentity.SupplierList{}, apperror.InteralError.Wrap(err)
	}
	return &httpentity.SupplierList{
		PaginationResponse: httpentity.NewPaginationResponse(total, pagination.Limit),
		Supplier:           Supplier,
	}, nil
}

// Update entity item
func (s *Service) Update(ctx context.Context, data httpentity.UpdateSupplierRequest) (*entity.Supplier, error) {
	Supplier := entity.Supplier{
		Id:                 data.Id,
		Name:               data.Name,
		Email:              data.Email,
		Phone:              data.Phone,
		StatusId:           data.StatusId,
		IsVerifiedSupplier: data.IsVerifiedSupplier,
	}
	affected, err := s.Repository.Update(ctx, &Supplier)
	if err != nil {
		return nil, apperror.InteralError.Wrap(err)
	}
	if affected == 0 {
		return nil, SupplierNotFound
	}
	return &Supplier, nil
}

// Get entity by id
func (s *Service) GetByID(ctx context.Context, SupplierId int) (entity.Supplier, error) {
	Supplier, err := s.Repository.GetByID(ctx, SupplierId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Supplier, SupplierNotFound
		}
		return Supplier, apperror.InteralError.Wrap(err)
	}

	return Supplier, nil
}

// Create entity Bulk
func (s *Service) Create(ctx context.Context, data httpentity.CreateSupplierRequest) (*entity.Supplier, error) {
	Supplier := entity.Supplier{}

	Supplier = entity.Supplier{
		Name:               data.Name,
		Email:              data.Email,
		Phone:              data.Phone,
		IsVerifiedSupplier: data.IsVerifiedSupplier,
	}

	err := s.Repository.Create(ctx, &Supplier)

	if err != nil {

		return nil, apperror.InteralError.Wrap(err)
	}
	return &Supplier, nil
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
