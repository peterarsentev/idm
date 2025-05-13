-- +goose Up
CREATE TABLE dionea_role (
                             id SERIAL NOT NULL,
                             name TEXT NOT NULL,
                             PRIMARY KEY (id)
);

CREATE TABLE dionea_user (
                             id SERIAL NOT NULL,
                             username TEXT NOT NULL,
                             password TEXT NOT NULL,
                             enabled BOOLEAN,
                             role_id INT NOT NULL,
                             PRIMARY KEY (id),
                             UNIQUE (username),
                             UNIQUE (role_id),
                             CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES dionea_role(id)
);

-- +goose Down
ALTER TABLE dionea_user DROP CONSTRAINT IF EXISTS fk_role_id;
DROP TABLE IF EXISTS dionea_user;
DROP TABLE IF EXISTS dionea_role;
