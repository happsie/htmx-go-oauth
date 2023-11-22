CREATE TABLE IF NOT EXISTS user_tokens
(
    user_id VARCHAR
(
    255
) PRIMARY KEY NOT NULL,
    access_token VARCHAR
(
    100
) NOT NULL,
    refresh_token VARCHAR
(
    100
) NOT NULL,
    FOREIGN KEY
(
    user_id
) REFERENCES users
(
    id
) ON DELETE CASCADE
    );