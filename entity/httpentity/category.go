package httpentity

import "github.com/techno/entity"

type CreateCategoryRequest struct {
	Id       int    `json:"id" bun:"id,pk,autoincrement"`
	Name     string `json:"name" validate:"required"`
	ParentId int    `json:"parent_id" validate:"required"`
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
	StatusId int `query:"status_id"`
	PaginationRequest
}
