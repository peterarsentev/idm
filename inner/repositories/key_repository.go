package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type KeyRepository struct {
	db *sqlx.DB
}

func NewKeyRepository(db *sqlx.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

func (r *KeyRepository) FindAll(ctx context.Context) ([]models.Key, error) {
	var keys []models.Key
	err := r.db.SelectContext(ctx, &keys, "SELECT * FROM dionea_key")
	return keys, err
}

func (r *KeyRepository) FindByID(ctx context.Context, ID int64) (models.Key, error) {
	var key models.Key
	err := r.db.GetContext(ctx, &key, "SELECT * FROM dionea_key WHERE id=$1", ID)
	return key, err
}

func (r *KeyRepository) Add(ctx context.Context, k models.Key) (int64, error) {
	query := "INSERT INTO dionea_key (name, filter_id) VALUES (:name, :filter_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, k)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
