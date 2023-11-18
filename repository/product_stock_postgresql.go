package repository

import (
	"context"

	"github.com/go-contact-service/entity"

	"github.com/uptrace/bun"
)

type ProductStockRepo struct {
	db *bun.DB
}

func NewProductStockRepo(db *bun.DB) *ProductStockRepo {
	return &ProductStockRepo{
		db: db,
	}
}

func (repo *ProductStockRepo) Update(ctx context.Context, ProductStock *entity.ProductStock) (int64, error) {
	res, err := repo.db.NewUpdate().Model(ProductStock).
		WherePK().
		ExcludeColumn("created_at", "created_by").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()
	return affected, nil
}

// Get single entity by id

func (repo *ProductStockRepo) GetByID(ctx context.Context, ProductStockId int) (entity.ProductStock, error) {
	entityDB := entity.ProductStock{}
	err := repo.db.NewSelect().Model(&entityDB).
		Where("id = ?", ProductStockId).
		Scan(ctx)
	return entityDB, err
}

// Get All Entity

func (repo *ProductStockRepo) List(ctx context.Context, pagination entity.Pagination, filter entity.ProductStockFilter) (int, []entity.ProductStock, error) {
	ProductStock := []entity.ProductStock{}

	query := repo.db.NewSelect().Model(&ProductStock)

	if pagination.Limit != 0 {
		query.Limit(pagination.Limit).
			Offset(pagination.Offset)
	}

	//query.Where("contact_id = ?", filter.ContactID)

	totalCount, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, ProductStock, nil
}

// Create entity

func (repo *ProductStockRepo) Create(ctx context.Context, entity *entity.ProductStock) error {
	_, err := repo.db.NewInsert().Model(entity).
		ExcludeColumn("created_at", "updated_at", "deleted_at", "updated_by").
		Returning("*").Exec(ctx)
	return err
}

func (repo *ProductStockRepo) Delete(ctx context.Context, id int) (int64, error) {

	res, err := repo.db.NewUpdate().
		Model((*entity.ProductStock)(nil)).
		Set("deleted_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()

	return affected, nil
}
