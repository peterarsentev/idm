package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
)

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
