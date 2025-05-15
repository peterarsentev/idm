package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) FindAll(ctx context.Context) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.SelectContext(ctx, &chats, "SELECT * FROM dionea_chat")
	return chats, err
}

func (r *ChatRepository) FindByID(ctx context.Context, ID int64) (models.Chat, error) {
	var chat models.Chat
	err := r.db.GetContext(ctx, &chat, "SELECT * FROM dionea_chat WHERE id=$1", ID)
	return chat, err
}

func (r *ChatRepository) Add(ctx context.Context, chat models.Chat) (int64, error) {
	query := "INSERT INTO dionea_chat (chat_id, username, title) VALUES (:chat_id, :username, :title) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, chat)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}
