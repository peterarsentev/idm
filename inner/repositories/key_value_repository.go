package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type KeyValueRepository struct {
	db *sqlx.DB
}

func NewKeyValueRepository(db *sqlx.DB) *KeyValueRepository {
	return &KeyValueRepository{db: db}
}

func (r *KeyValueRepository) FindAll(ctx context.Context) ([]models.KeyValue, error) {
	var keyValues []models.KeyValue
	err := r.db.SelectContext(ctx, &keyValues, "SELECT * FROM dionea_key_value")
	return keyValues, err
}

func (r *KeyValueRepository) FindByID(ctx context.Context, ID int64) (models.KeyValue, error) {
	var keyValue models.KeyValue
	err := r.db.GetContext(ctx, &keyValue, "SELECT * FROM dionea_key_value WHERE id=$1", ID)
	return keyValue, err
}

func (r *KeyValueRepository) Add(ctx context.Context, kv models.KeyValue) (int64, error) {
	query := "INSERT INTO dionea_key_value (value, key_id) VALUES (:value, :key_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, kv)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
