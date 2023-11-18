package repository

import (
	"context"

	"github.com/techno/entity"

	"github.com/uptrace/bun"
)

type BrandRepo struct {
	db *bun.DB
}

func NewBrandRepo(db *bun.DB) *BrandRepo {
	return &BrandRepo{
		db: db,
	}
}

func (repo *BrandRepo) Update(ctx context.Context, Brand *entity.Brand) (int64, error) {
	res, err := repo.db.NewUpdate().Model(Brand).
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

func (repo *BrandRepo) GetByID(ctx context.Context, BrandId int) (entity.Brand, error) {
	entityDB := entity.Brand{}
	err := repo.db.NewSelect().Model(&entityDB).
		Where("id = ?", BrandId).
		Scan(ctx)
	return entityDB, err
}

// Get All Entity

func (repo *BrandRepo) List(ctx context.Context, pagination entity.Pagination, filter entity.BrandFilter) (int, []entity.Brand, error) {
	Brand := []entity.Brand{}

	query := repo.db.NewSelect().Model(&Brand)

	if filter.StatusId != nil {
		query.Where("status_id = ?", filter.StatusId)
	}
	if pagination.Limit != 0 {
		query.Limit(pagination.Limit).
			Offset(pagination.Offset)
	}

	totalCount, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, Brand, nil
}

// Create entity

func (repo *BrandRepo) Create(ctx context.Context, entity *entity.Brand) error {
	_, err := repo.db.NewInsert().Model(entity).
		ExcludeColumn("created_at", "updated_at", "deleted_at", "updated_by", "status_id").
		Returning("*").Exec(ctx)
	return err
}

func (repo *BrandRepo) Delete(ctx context.Context, id int) (int64, error) {

	res, err := repo.db.NewUpdate().
		Model((*entity.Brand)(nil)).
		Set("deleted_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()

	return affected, nil
}
