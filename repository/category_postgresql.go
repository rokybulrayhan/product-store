package repository

import (
	"context"

	"github.com/go-contact-service/entity"

	"github.com/uptrace/bun"
)

type CategoryRepo struct {
	db *bun.DB
}

func NewCategoryRepo(db *bun.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (repo *CategoryRepo) Update(ctx context.Context, Category *entity.Category) (int64, error) {
	res, err := repo.db.NewUpdate().Model(Category).
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

func (repo *CategoryRepo) GetByID(ctx context.Context, CategoryId int) (entity.Category, error) {
	entityDB := entity.Category{}
	err := repo.db.NewSelect().Model(&entityDB).
		Where("id = ?", CategoryId).
		Scan(ctx)
	return entityDB, err
}

// Get All Entity

func (repo *CategoryRepo) List(ctx context.Context, pagination entity.Pagination, filter entity.CategoryFilter) (int, []entity.Category, error) {
	Category := []entity.Category{}

	query := repo.db.NewSelect().Model(&Category)

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
	return totalCount, Category, nil
}

// Create entity

func (repo *CategoryRepo) Create(ctx context.Context, entity *entity.Category) error {
	_, err := repo.db.NewInsert().Model(entity).
		ExcludeColumn("created_at", "updated_at", "deleted_at", "updated_by", "status_id").
		Returning("*").Exec(ctx)
	return err
}

func (repo *CategoryRepo) Delete(ctx context.Context, id int) (int64, error) {

	res, err := repo.db.NewUpdate().
		Model((*entity.Category)(nil)).
		Set("deleted_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	affected, _ := res.RowsAffected()

	return affected, nil
}
