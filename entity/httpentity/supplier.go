package httpentity

import "github.com/go-contact-service/entity"

type CreateSupplierRequest struct {
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"email"`
	Phone              string `json:"phone" validate:"phone"`
	StatusId           bool   `json:"status_id" validate:"status_id"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" validate:"is_verified_supplier"`
}

func (input *CreateSupplierRequest) Validate() []FieldError {
	return validate(input)
}

type UpdateSupplierRequest struct {
	Id       int  `json:"id"`
	StatusId bool `json:"status_id"`
	CreateSupplierRequest
}

func (input *UpdateSupplierRequest) Validate() []FieldError {
	return validate(input)
}

type SupplierList struct {
	PaginationResponse
	Supplier []entity.Supplier `json:"Suppliers"`
}

type SupplierParams struct {
	PaginationRequest
}
