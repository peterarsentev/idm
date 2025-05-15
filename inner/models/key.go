package models

type Key struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	FilterID int64  `db:"filter_id"`
}
