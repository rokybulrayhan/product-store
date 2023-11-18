package httpentity

import "github.com/techno/entity"

type CreateProductRequest struct {
	Name           string                    `json:"name" validate:"required"`
	Description    string                    `json:"description" `
	Specifications string                    `json:"specifications"`
	BrandId        int                       `json:"brand_id" validate:"required"`
	CategoryId     int                       `json:"category_id" validate:"required"`
	SupplierId     int                       `json:"supplier_id" validate:"required"`
	UnitPrice      int                       `json:"unit_price" validate:"required"`
	DiscountPrice  int                       `json:"discount_price" validate:"required"`
	Tags           string                    `json:"tags" validate:"required"`
	ProductStock   CreateProductStockRequest `json:"product_stock"`
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
	Name       string `query:"name"`
	MinPrice   int    `query:"min_price"`
	MaxPrice   int    `query:"max_price"`
	BrandId    string `query:"brand_id"`
	CategoryId int    `query:"category_id"`
	SupplierId int    `query:"supplier_id"`
	StatusId   int    `query:"status_id"`
	PaginationRequest
}
