package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel  `bun:"products"`
	Id             int    `json:"id" bun:"id,pk,autoincrement"`
	Name           string `json:"name" bun:"name"`
	Description    string `json:"description" bun:"description"`
	Specifications string `json:"specifications" bun:"specifications"`
	BrandId        int    `json:"brand_id" bun:"brand_id"`
	CategoryId     int    `json:"category_id" bun:"category_id"`
	SupplierId     int    `json:"supplier_id" bun:"supplier_id"`
	UnitPrice      int    `json:"unit_price" bun:"unit_price"`
	DiscountPrice  int    `json:"discount_price" bun:"discount_price"`
	Tags           string `json:"tags" bun:"tags"`
	StatusId       bool   `json:"status_id" bun:"status_id"`

	CreatedAt time.Time    `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime `json:"updated_at" bun:"updated_at"`
	DeletedAt time.Time    `json:"-" bun:"deleted_at,soft_delete,nullzero"`
	CreatedBy string       `json:"created_by" bun:"created_by,nullzero"`
	UpdatedBy string       `json:"updated_by" bun:"updated_by,nullzero"`
}
type ProductFilter struct {
	Name       string
	MinPrice   int
	MaxPrice   int
	BrandId    int
	CategoryId int
	SupplierId int
	StatusId   *bool
}
