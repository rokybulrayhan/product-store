package repository

import (
	"context"

	"github.com/go-contact-service/entity"

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

	if filter.StatusId != nil {
		query.Where("status_id = ?", filter.StatusId)
	}

	if pagination.Limit != 0 {
		query.Limit(pagination.Limit).
			Offset(pagination.Offset)
	}

	//query.Where("contact_id = ?", filter.ContactID)

	totalCount, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, Product, nil
}

// Create entity

func (repo *ProductRepo) Create(ctx context.Context, entity *entity.Product) error {
	_, err := repo.db.NewInsert().Model(entity).
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
