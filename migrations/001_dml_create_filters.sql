-- +goose Up
CREATE TABLE dionea_filter (
                               id SERIAL NOT NULL,
                               name VARCHAR(2000),
                               PRIMARY KEY (id)
);

CREATE TABLE dionea_key (
                            id SERIAL NOT NULL,
                            name VARCHAR(2000),
                            filter_id INT NOT NULL,
                            PRIMARY KEY (id),
                            CONSTRAINT fk_filter_id FOREIGN KEY (filter_id) REFERENCES dionea_filter(id)
);

CREATE TABLE dionea_key_value (
                                  id SERIAL NOT NULL,
                                  value VARCHAR(2000),
                                  key_id INT NOT NULL,
                                  PRIMARY KEY (id),
                                  CONSTRAINT fk_key_id FOREIGN KEY (key_id) REFERENCES dionea_key(id)
);

-- +goose Down
DROP TABLE IF EXISTS dionea_key_value;
DROP TABLE IF EXISTS dionea_key;
DROP TABLE IF EXISTS dionea_filter;
