CREATE TABLE IF NOT EXISTS jwt (
    user_uuid uuid PRIMARY KEY,
    refresh_token_hash varchar(128) NOT NULL
)