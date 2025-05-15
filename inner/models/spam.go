package models

type Spam struct {
	ID        int64  `db:"id"`
	Text      string `db:"text"`
	Time      string `db:"time"`
	ChatID    int64  `db:"chat_id"`
	ContactID int64  `db:"contact_id"`
}
