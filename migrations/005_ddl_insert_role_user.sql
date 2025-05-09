INSERT INTO dionea_role (name)
VALUES ('ROLE_ADMIN');

DELETE FROM dionea_user WHERE username = 'root';

INSERT INTO dionea_user (username, password, enabled, role_id)
VALUES (
           'root',
           '$2a$12$zD8AMBCfUJrWI0AiY.pwuenlyDscXa/AlC58NlRmUdvkdtrPYWaFS',
           TRUE,
           1
       );


INSERT INTO dionea_role (name) VALUES ('ROLE_USER');

CREATE TABLE dionea_contact (
                                id SERIAL NOT NULL PRIMARY KEY,
                                tg_user_id NUMERIC NOT NULL UNIQUE,
                                username TEXT NOT NULL UNIQUE,
                                first_name TEXT,
                                last_name TEXT
);
INSERT INTO dionea_contact (tg_user_id, username) VALUES (-1, 'unknown');
ALTER TABLE dionea_spam ADD COLUMN contact_id INT NOT NULL DEFAULT 1;
ALTER TABLE dionea_spam ADD CONSTRAINT fk_contact_id FOREIGN KEY (contact_id) REFERENCES dionea_contact(id);

CREATE TABLE dionea_chat (
                             id SERIAL NOT NULL PRIMARY KEY,
                             chat_id NUMERIC NOT NULL UNIQUE,
                             username TEXT,
                             title TEXT
);
INSERT INTO dionea_chat (chat_id) VALUES (-1);
ALTER TABLE dionea_spam DROP COLUMN chat_id;
ALTER TABLE dionea_spam DROP COLUMN chat_name;
ALTER TABLE dionea_spam ADD COLUMN chat_id INT NOT NULL DEFAULT 1;
ALTER TABLE dionea_spam ADD CONSTRAINT fk_chat_id FOREIGN KEY (chat_id) REFERENCES dionea_chat(id);

ALTER TABLE dionea_contact ADD COLUMN ham INT NOT NULL DEFAULT 0;
ALTER TABLE dionea_contact ADD COLUMN spam INT NOT NULL DEFAULT 0;

ALTER TABLE dionea_contact ADD COLUMN restrict BOOLEAN NOT NULL DEFAULT false;

