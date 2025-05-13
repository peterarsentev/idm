-- +goose Up
CREATE TABLE dionea_vote (
                             id SERIAL NOT NULL,
                             chat_id NUMERIC,
                             message_id NUMERIC,
                             user_id NUMERIC,
                             vote INT,
                             PRIMARY KEY (id)
);

ALTER TABLE dionea_vote
    ADD CONSTRAINT dionea_vote_chat_message_user UNIQUE (chat_id, message_id, user_id);

-- +goose Down
ALTER TABLE dionea_vote
    DROP CONSTRAINT IF EXISTS dionea_vote_chat_message_user;

DROP TABLE IF EXISTS dionea_vote;
