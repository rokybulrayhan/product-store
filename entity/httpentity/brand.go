package httpentity

import "github.com/go-contact-service/entity"

type CreateBrandRequest struct {
	Id   int    `json:"id" bun:"id,pk,autoincrement"`
	Name string `json:"name" validate:"required"`
}

func (input *CreateBrandRequest) Validate() []FieldError {
	return validate(input)
}

type UpdateBrandRequest struct {
	Id       int  `json:"id"`
	StatusId bool `json:"status_id"`
	CreateBrandRequest
}

func (input *UpdateBrandRequest) Validate() []FieldError {
	return validate(input)
}

type BrandList struct {
	PaginationResponse
	Brand []entity.Brand `json:"Brands"`
}

type BrandParams struct {
	PaginationRequest
}
