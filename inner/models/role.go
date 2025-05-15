package models

type Role struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
