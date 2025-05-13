package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

type FilterRepository struct {
	db *sqlx.DB
}

func NewFilterRepository(db *sqlx.DB) *FilterRepository {
	return &FilterRepository{db: db}
}

func (r *FilterRepository) FindAll(ctx context.Context) (filters []models.Filter, err error) {
	err = r.db.SelectContext(ctx, &filters, "SELECT * FROM dionea_filter")
	return filters, err
}

func (r *FilterRepository) FindByID(ctx context.Context, ID int64) (filter models.Filter, err error) {
	err = r.db.GetContext(ctx, &filter, "SELECT * FROM dionea_filter WHERE id=$1", ID)
	return filter, err
}

func (r *FilterRepository) Add(ctx context.Context, f models.Filter) (int64, error) {
	query := "INSERT INTO dionea_filter (name) VALUES (:name) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, f)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

type KeyRepository struct {
	db *sqlx.DB
}

func NewKeyRepository(db *sqlx.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

func (r *KeyRepository) FindAll(ctx context.Context) ([]models.Key, error) {
	var keys []models.Key
	err := r.db.SelectContext(ctx, &keys, "SELECT * FROM dionea_key")
	return keys, err
}

func (r *KeyRepository) FindByID(ctx context.Context, ID int64) (models.Key, error) {
	var key models.Key
	err := r.db.GetContext(ctx, &key, "SELECT * FROM dionea_key WHERE id=$1", ID)
	return key, err
}

type KeyValueRepository struct {
	db *sqlx.DB
}

func NewKeyValueRepository(db *sqlx.DB) *KeyValueRepository {
	return &KeyValueRepository{db: db}
}

func (r *KeyValueRepository) FindAll(ctx context.Context) ([]models.KeyValue, error) {
	var keyValues []models.KeyValue
	err := r.db.SelectContext(ctx, &keyValues, "SELECT * FROM dionea_key_value")
	return keyValues, err
}

func (r *KeyValueRepository) FindByID(ctx context.Context, ID int64) (models.KeyValue, error) {
	var keyValue models.KeyValue
	err := r.db.GetContext(ctx, &keyValue, "SELECT * FROM dionea_key_value WHERE id=$1", ID)
	return keyValue, err
}

type SpamRepository struct {
	db *sqlx.DB
}

func NewSpamRepository(db *sqlx.DB) *SpamRepository {
	return &SpamRepository{db: db}
}

func (r *SpamRepository) FindAll(ctx context.Context) ([]models.Spam, error) {
	var spams []models.Spam
	err := r.db.SelectContext(ctx, &spams, "SELECT * FROM dionea_spam")
	return spams, err
}

func (r *SpamRepository) FindByID(ctx context.Context, ID int64) (models.Spam, error) {
	var spam models.Spam
	err := r.db.GetContext(ctx, &spam, "SELECT * FROM dionea_spam WHERE id=$1", ID)
	return spam, err
}

type VoteRepository struct {
	db *sqlx.DB
}

func NewVoteRepository(db *sqlx.DB) *VoteRepository {
	return &VoteRepository{db: db}
}

func (r *VoteRepository) FindAll(ctx context.Context) ([]models.Vote, error) {
	var votes []models.Vote
	err := r.db.SelectContext(ctx, &votes, "SELECT * FROM dionea_vote")
	return votes, err
}

func (r *VoteRepository) FindByID(ctx context.Context, ID int64) (models.Vote, error) {
	var vote models.Vote
	err := r.db.GetContext(ctx, &vote, "SELECT * FROM dionea_vote WHERE id=$1", ID)
	return vote, err
}

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.SelectContext(ctx, &roles, "SELECT * FROM dionea_role")
	return roles, err
}

func (r *RoleRepository) FindByID(ctx context.Context, ID int64) (models.Role, error) {
	var role models.Role
	err := r.db.GetContext(ctx, &role, "SELECT * FROM dionea_role WHERE id=$1", ID)
	return role, err
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.SelectContext(ctx, &users, "SELECT * FROM dionea_user")
	return users, err
}

func (r *UserRepository) FindByID(ctx context.Context, ID int64) (models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM dionea_user WHERE id=$1", ID)
	return user, err
}

type ContactRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) FindAll(ctx context.Context) ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.SelectContext(ctx, &contacts, "SELECT * FROM dionea_contact")
	return contacts, err
}

func (r *ContactRepository) FindByID(ctx context.Context, ID int64) (models.Contact, error) {
	var contact models.Contact
	err := r.db.GetContext(ctx, &contact, "SELECT * FROM dionea_contact WHERE id=$1", ID)
	return contact, err
}

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

func (r *KeyRepository) Add(ctx context.Context, k models.Key) (int64, error) {
	query := "INSERT INTO dionea_key (name, filter_id) VALUES (:name, :filter_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, k)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *KeyValueRepository) Add(ctx context.Context, kv models.KeyValue) (int64, error) {
	query := "INSERT INTO dionea_key_value (value, key_id) VALUES (:value, :key_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, kv)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *SpamRepository) Add(ctx context.Context, s models.Spam) (int64, error) {
	query := "INSERT INTO dionea_spam (text, time, chat_id, contact_id) VALUES (:text, :time, :chat_id, :contact_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, s)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *VoteRepository) Add(ctx context.Context, v models.Vote) (int64, error) {
	query := "INSERT INTO dionea_vote (chat_id, message_id, user_id, vote) VALUES (:chat_id, :message_id, :user_id, :vote) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, v)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *RoleRepository) Add(ctx context.Context, role models.Role) (int64, error) {
	query := "INSERT INTO dionea_role (name) VALUES (:name) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, role)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *UserRepository) Add(ctx context.Context, u models.User) (int64, error) {
	query := "INSERT INTO dionea_user (username, password, enabled, role_id) VALUES (:username, :password, :enabled, :role_id) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, u)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
}

func (r *ContactRepository) Add(ctx context.Context, c models.Contact) (int64, error) {
	query := "INSERT INTO dionea_contact (tg_user_id, username, first_name, last_name, ham, spam, restrict) VALUES (:tg_user_id, :username, :first_name, :last_name, :ham, :spam, :restrict) RETURNING id"
	var ID int64
	rows, err := r.db.NamedQueryContext(ctx, query, c)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		return ID, rows.Scan(&ID)
	}
	return ID, err
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
