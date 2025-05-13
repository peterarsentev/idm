package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type SpamRepository struct {
	db *sqlx.DB
}

func NewSpamRepository(db *sqlx.DB) *SpamRepository {
	return &SpamRepository{db: db}
}

func (r *SpamRepository) FindAll(ctx context.Context) ([]models.Spam, error) {
	var spams []models.Spam
	err := r.db.SelectContext(ctx, &spams, "SELECT * FROM dionea_spam")
	return spams, err
}

func (r *SpamRepository) FindByID(ctx context.Context, ID int64) (models.Spam, error) {
	var spam models.Spam
	err := r.db.GetContext(ctx, &spam, "SELECT * FROM dionea_spam WHERE id=$1", ID)
	return spam, err
}

func (r *SpamRepository) Add(ctx context.Context, s models.Spam) (int64, error) {
	query := "INSERT INTO dionea_spam (text, time, chat_id, contact_id) VALUES (:text, :time, :chat_id, :contact_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, s)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
