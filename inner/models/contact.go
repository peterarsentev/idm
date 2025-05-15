package models

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
