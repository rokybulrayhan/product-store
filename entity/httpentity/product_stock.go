package httpentity

import "github.com/go-contact-service/entity"

type CreateProductStockRequest struct {
	ProductId     int `json:"product_id" validate:"product_id"`
	StockQuantity int `json:"stock_quantity" validate:"stock_quantity"`
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
	PaginationRequest
}
