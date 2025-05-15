package models

type KeyValue struct {
	ID    int64  `db:"id"`
	Value string `db:"value"`
	KeyID int64  `db:"key_id"`
}
