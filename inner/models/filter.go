package models

type Filter struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
