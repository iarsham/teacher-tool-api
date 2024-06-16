CREATE TABLE IF NOT EXISTS questions
(
    id         SERIAL PRIMARY KEY,
    lesson     INT                      NOT NULL,
    title      VARCHAR(255)             NOT NULL,
    grade      INT                      NOT NULL CHECK (level IN (0, 1, 2, 3, 4, 5)),
    level      INT                      NOT NULL CHECK (level IN (0, 1, 2)),
    views      INT                      NOT NULL DEFAULT 0,
    used       INT                      NOT NULL DEFAULT 0,
    file       VARCHAR(255)             NOT NULL,
    userid     INT REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
)