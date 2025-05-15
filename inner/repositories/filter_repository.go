package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type FilterRepository struct {
	db *sqlx.DB
}

func NewFilterRepository(db *sqlx.DB) *FilterRepository {
	return &FilterRepository{db: db}
}

func (r *FilterRepository) FindAll(ctx context.Context) (filters []models.Filter, err error) {
	err = r.db.SelectContext(ctx, &filters, "SELECT * FROM dionea_filter")
	return filters, err
}

func (r *FilterRepository) FindByID(ctx context.Context, ID int64) (filter *models.Filter, err error) {
	filter = &models.Filter{}
	err = r.db.GetContext(ctx, filter, "SELECT * FROM dionea_filter WHERE id=$1", ID)
	return filter, err
}

func (r *FilterRepository) Add(ctx context.Context, f models.Filter) (int64, error) {
	query := "INSERT INTO dionea_filter (name) VALUES (:name) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, f)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
