package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type ProductStock struct {
	bun.BaseModel `bun:"product_stocks"`
	Id            int `json:"id" bun:"id,pk,autoincrement"`
	ProductId     int `json:"product_id" bun:"product_id"`
	StockQuantity int `json:"stock_quantity" bun:"stock_quantity"`

	CreatedAt time.Time    `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime `json:"updated_at" bun:"updated_at"`
	DeletedAt time.Time    `json:"-" bun:"deleted_at,soft_delete,nullzero"`
	CreatedBy string       `json:"created_by" bun:"created_by,nullzero"`
	UpdatedBy string       `json:"updated_by" bun:"updated_by,nullzero"`
}
type ProductStockFilter struct {
	StatusId *bool
}
