package repository

import (
	"context"

	"github.com/techno/entity"

	"github.com/uptrace/bun"
)

type SupplierRepo struct {
	db *bun.DB
}

func NewSupplierRepo(db *bun.DB) *SupplierRepo {
	return &SupplierRepo{
		db: db,
	}
}

func (repo *SupplierRepo) Update(ctx context.Context, Supplier *entity.Supplier) (int64, error) {
	res, err := repo.db.NewUpdate().Model(Supplier).
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

func (repo *SupplierRepo) GetByID(ctx context.Context, SupplierId int) (entity.Supplier, error) {
	entityDB := entity.Supplier{}
	err := repo.db.NewSelect().Model(&entityDB).
		Where("id = ?", SupplierId).
		Scan(ctx)
	return entityDB, err
}

// Get All Entity

func (repo *SupplierRepo) List(ctx context.Context, pagination entity.Pagination, filter entity.SupplierFilter) (int, []entity.Supplier, error) {
	Supplier := []entity.Supplier{}

	query := repo.db.NewSelect().Model(&Supplier)
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
	return totalCount, Supplier, nil
}

// Create entity

func (repo *SupplierRepo) Create(ctx context.Context, entity *entity.Supplier) error {
	_, err := repo.db.NewInsert().Model(entity).
		ExcludeColumn("created_at", "updated_at", "deleted_at", "updated_by", "status_id").
		Returning("*").Exec(ctx)
	return err
}

func (repo *SupplierRepo) Delete(ctx context.Context, id int) (int64, error) {

	res, err := repo.db.NewUpdate().
		Model((*entity.Supplier)(nil)).
		Set("deleted_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()

	return affected, nil
}
