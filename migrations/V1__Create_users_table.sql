CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY,
    username varchar(32),
    password varchar(256) NOT NULL,
    email varchar(32) UNIQUE NOT NULL
)