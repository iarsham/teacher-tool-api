CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    phone_number varchar(20)              NOT NULL UNIQUE,
    password     varchar(100)             NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);