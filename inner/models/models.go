package models

type Filter struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Key struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	FilterID int64  `db:"filter_id"`
}

type KeyValue struct {
	ID    int64  `db:"id"`
	Value string `db:"value"`
	KeyID int64  `db:"key_id"`
}

type Spam struct {
	ID        int64  `db:"id"`
	Text      string `db:"text"`
	Time      string `db:"time"`
	ChatID    int64  `db:"chat_id"`
	ContactID int64  `db:"contact_id"`
}

type Vote struct {
	ID        int64 `db:"id"`
	ChatID    int64 `db:"chat_id"`
	MessageID int64 `db:"message_id"`
	UserID    int64 `db:"user_id"`
	Vote      int64 `db:"vote"`
}

type Role struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Enabled  bool   `db:"enabled"`
	RoleID   int64  `db:"role_id"`
}

type Contact struct {
	ID        int     `db:"id"`
	TgUserID  int64   `db:"tg_user_id"`
	Username  string  `db:"username"`
	FirstName *string `db:"first_name"` // Changed to *string to handle NULL
	LastName  *string `db:"last_name"`  // Changed to *string to handle NULL
	Ham       int     `db:"ham"`
	Spam      int     `db:"spam"`
	Restrict  bool    `db:"restrict"`
}

type Chat struct {
	ID       int64   `db:"id"`
	ChatID   int64   `db:"chat_id"`
	Username *string `db:"username"`
	Title    *string `db:"title"`
}
