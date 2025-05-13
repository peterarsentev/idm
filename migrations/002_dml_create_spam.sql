-- +goose Up
CREATE TABLE dionea_spam (
                             id SERIAL NOT NULL,
                             text TEXT,
                             time TIMESTAMP WITHOUT TIME ZONE,
                             chat_id NUMERIC,
                             chat_name TEXT,
                             PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE IF EXISTS dionea_spam;
