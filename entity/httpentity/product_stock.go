package httpentity

import "github.com/go-contact-service/entity"

type CreateProductStockRequest struct {
	ProductId     int `json:"product_id" validate:"required"`
	StockQuantity int `json:"stock_quantity" validate:"required"`
}

func (input *CreateProductStockRequest) Validate() []FieldError {
	return validate(input)
}

type UpdateProductStockRequest struct {
	Id int `json:"id"`
	CreateProductStockRequest
}

func (input *UpdateProductStockRequest) Validate() []FieldError {
	return validate(input)
}

type ProductStockList struct {
	PaginationResponse
	ProductStock []entity.ProductStock `json:"ProductStocks"`
}

type ProductStockParams struct {
	StatusId int `query:"status_id"`
	PaginationRequest
}
