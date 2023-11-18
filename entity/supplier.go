package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Supplier struct {
	bun.BaseModel      `bun:"suppliers"`
	Id                 int    `json:"id" bun:"id,pk,autoincrement"`
	Name               string `json:"name" bun:"name"`
	Email              string `json:"email" bun:"email"`
	Phone              string `json:"phone" bun:"phone"`
	StatusId           bool   `json:"status_id" bun:"status_id"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" bun:"is_verified_supplier"`

	CreatedAt time.Time    `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime `json:"updated_at" bun:"updated_at"`
	DeletedAt time.Time    `json:"-" bun:"deleted_at,soft_delete,nullzero"`
	CreatedBy string       `json:"created_by" bun:"created_by,nullzero"`
	UpdatedBy string       `json:"updated_by" bun:"updated_by,nullzero"`
}
type SupplierFilter struct {
	StatusId *bool
}
