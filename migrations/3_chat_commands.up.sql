CREATE TABLE IF NOT EXISTS twitch_commands
(
    id VARCHAR
(
    255
) PRIMARY KEY NOT NULL,
    user_id VARCHAR
(
    255
) NOT NULL,
    command VARCHAR
(
    30
) NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY
(
    user_id
) REFERENCES users
(
    id
) ON DELETE CASCADE,
    INDEX user_command_idx
(
    user_id,
    command
)
    );