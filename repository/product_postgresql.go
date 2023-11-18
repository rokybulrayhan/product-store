package repository

import (
	"context"
	"fmt"

	"github.com/techno/entity"

	"github.com/uptrace/bun"
)

type ProductRepo struct {
	db *bun.DB
}

func NewProductRepo(db *bun.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (repo *ProductRepo) Update(ctx context.Context, Product *entity.Product) (int64, error) {
	res, err := repo.db.NewUpdate().Model(Product).
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

func (repo *ProductRepo) GetByID(ctx context.Context, ProductId int) (entity.Product, error) {
	entityDB := entity.Product{}
	err := repo.db.NewSelect().Model(&entityDB).
		Where("id = ?", ProductId).
		Scan(ctx)
	return entityDB, err
}

// Get All Entity

func (repo *ProductRepo) List(ctx context.Context, pagination entity.Pagination, filter entity.ProductFilter) (int, []entity.Product, error) {
	Product := []entity.Product{}

	query := repo.db.NewSelect().Model(&Product)

	if filter.MaxPrice != 0 {
		query.Where("unit_price >= ?", filter.MaxPrice)
	}
	if filter.MinPrice != 0 {
		query.Where("unit_price <= ?", filter.MinPrice)
	}

	if filter.StatusId != nil {
		query.Where("status_id = ?", filter.StatusId)
	}

	if filter.Name != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.CategoryId != 0 {
		query.Where("category_id = ?", filter.CategoryId)
	}

	if filter.SupplierId != 0 {
		query.Where("supplier_id = ?", filter.SupplierId)
	}

	if len(filter.BrandId) != 0 {
		query.Where("brand_id IN (?)", bun.In(filter.BrandId))

	}

	if pagination.Limit != 0 {
		query.Limit(pagination.Limit).
			Offset(pagination.Offset)
	}

	query.Order("unit_price ASC")
	totalCount, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, Product, nil
}

// Create entity

func (repo *ProductRepo) Create(ctx context.Context, entity *entity.Product, tx *bun.Tx) error {

	var query *bun.InsertQuery

	if tx != nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}
	_, err := query.Model(entity).
		ExcludeColumn("created_at", "updated_at", "deleted_at", "updated_by", "status_id").
		Returning("*").Exec(ctx)
	return err
}

func (repo *ProductRepo) Delete(ctx context.Context, id int) (int64, error) {

	res, err := repo.db.NewUpdate().
		Model((*entity.Product)(nil)).
		Set("deleted_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()

	return affected, nil
}
func (repo *ProductRepo) GetTx(ctx context.Context) (*bun.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	return &tx, err
}
