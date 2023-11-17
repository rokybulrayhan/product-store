package httpentity

import "github.com/go-contact-service/entity"

type CreateCategoryRequest struct {
	Id       int    `json:"id" bun:"id,pk,autoincrement"`
	Name     string `json:"name" validate:"name"`
	ParentId string `json:"parent_id" validate:"parent_id"`
}

func (input *CreateCategoryRequest) Validate() []FieldError {
	return validate(input)
}

type UpdateCategoryRequest struct {
	Id       int  `json:"id"`
	StatusId bool `json:"status_id"`
	CreateCategoryRequest
}

func (input *UpdateCategoryRequest) Validate() []FieldError {
	return validate(input)
}

type CategoryList struct {
	PaginationResponse
	Category []entity.Category `json:"Categorys"`
}

type CategoryParams struct {
	PaginationRequest
}
