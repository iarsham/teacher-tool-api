CREATE TABLE IF NOT EXISTS template
(
    id         SERIAL PRIMARY KEY,
    userid     INT REFERENCES users (id),
    file       VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);