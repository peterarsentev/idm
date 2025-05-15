package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.SelectContext(ctx, &users, "SELECT * FROM dionea_user")
	return users, err
}

func (r *UserRepository) FindByID(ctx context.Context, ID int64) (models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM dionea_user WHERE id=$1", ID)
	return user, err
}

func (r *UserRepository) Add(ctx context.Context, u models.User) (int64, error) {
	query := "INSERT INTO dionea_user (username, password, enabled, role_id) VALUES (:username, :password, :enabled, :role_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, u)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
