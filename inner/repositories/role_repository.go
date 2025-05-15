package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.SelectContext(ctx, &roles, "SELECT * FROM dionea_role")
	return roles, err
}

func (r *RoleRepository) FindByID(ctx context.Context, ID int64) (models.Role, error) {
	var role models.Role
	err := r.db.GetContext(ctx, &role, "SELECT * FROM dionea_role WHERE id=$1", ID)
	return role, err
}

func (r *RoleRepository) Add(ctx context.Context, role models.Role) (int64, error) {
	query := "INSERT INTO dionea_role (name) VALUES (:name) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, role)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
