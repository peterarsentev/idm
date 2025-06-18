package models

type Chat struct {
	ID       int64   `db:"id"`
	ChatID   int64   `db:"chat_id"`
	Username *string `db:"username"`
	Title    *string `db:"title"`
}
