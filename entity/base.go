package entity

import (
	"time"
)

const (
	ProjectName               string = "gmv_inventory_service"
	PromotionalOfferTableName string = "promotional_offer"
	CacheDuration             int    = 3600
)

type BaseEntity struct {
	ID        int        `json:"id" bun:",pk,autoincrement" validate:"omitempty"`
	CreatedBy string     `json:"created_by" bun:"type:uuid" validate:"omitempty"`
	UpdatedBy string     `json:"updated_by" bun:"type:uuid" validate:"omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-" bun:",soft_delete"`
}

type BaseEntityList struct {
	ID        int        `json:"id" bun:",pk"`
	DeletedAt *time.Time `json:"-" bun:",soft_delete"`
}

type SlugCheckerEntity struct {
	Available bool `json:"available"`
}
