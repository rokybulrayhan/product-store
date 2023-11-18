package httpentity

import "github.com/techno/entity"

type CreateSupplierRequest struct {
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"required"`
	Phone              string `json:"phone" validate:"required"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" validate:"required"`
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
	StatusId int `query:"status_id"`
	PaginationRequest
}
