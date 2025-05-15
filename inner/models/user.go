package models

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Enabled  bool   `db:"enabled"`
	RoleID   int64  `db:"role_id"`
}
