package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type ContactRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) FindAll(ctx context.Context) ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.SelectContext(ctx, &contacts, "SELECT * FROM dionea_contact")
	return contacts, err
}

func (r *ContactRepository) FindByID(ctx context.Context, ID int64) (models.Contact, error) {
	var contact models.Contact
	err := r.db.GetContext(ctx, &contact, "SELECT * FROM dionea_contact WHERE id=$1", ID)
	return contact, err
}

func (r *ContactRepository) Add(ctx context.Context, c models.Contact) (int64, error) {
	query := "INSERT INTO dionea_contact (tg_user_id, username, first_name, last_name, ham, spam, restrict) VALUES (:tg_user_id, :username, :first_name, :last_name, :ham, :spam, :restrict) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, c)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
