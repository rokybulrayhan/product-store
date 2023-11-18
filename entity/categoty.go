package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"categories"`
	Id            int       `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name" bun:"name"`
	ParentId      int       `json:"parent_id" bun:"parent_id"`
	StatusId      bool      `json:"status_id" bun:"status_id"`
	Sequence      int       `json:"sequence" bun:"sequence"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`

	UpdatedAt bun.NullTime `json:"updated_at" bun:"updated_at"`
	DeletedAt time.Time    `json:"-" bun:"deleted_at,soft_delete,nullzero"`
	CreatedBy string       `json:"created_by" bun:"created_by,nullzero"`
	UpdatedBy string       `json:"updated_by" bun:"updated_by,nullzero"`
}
type CategoryFilter struct {
	StatusId *bool
}
