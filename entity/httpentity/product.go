package httpentity

import "github.com/go-contact-service/entity"

type CreateProductRequest struct {
	Name           string `json:"name" validate:"name"`
	Description    string `json:"description" validate:"description"`
	Specifications string `json:"specifications" validate:"specifications"`
	BrandId        string `json:"brand_id" validate:"brand_id"`
	CategoryId     string `json:"category_id" validate:"category_id"`
	SupplierId     string `json:"supplier_id" validate:"supplier_id"`
	UnitPrice      string `json:"unit_price" validate:"unit_price"`
	DiscountPrice  string `json:"discount_price" validate:"discount_price"`
	Tags           string `json:"tags" validate:"tags"`
}

func (input *CreateProductRequest) Validate() []FieldError {
	return validate(input)
}

type UpdateProductRequest struct {
	Id       int  `json:"id"`
	StatusId bool `json:"status_id"`
	CreateProductRequest
}

func (input *UpdateProductRequest) Validate() []FieldError {
	return validate(input)
}

type ProductList struct {
	PaginationResponse
	Product []entity.Product `json:"Products"`
}

type ProductParams struct {
	PaginationRequest
}
