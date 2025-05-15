package models

type Vote struct {
	ID        int64 `db:"id"`
	ChatID    int64 `db:"chat_id"`
	MessageID int64 `db:"message_id"`
	UserID    int64 `db:"user_id"`
	Vote      int64 `db:"vote"`
}
