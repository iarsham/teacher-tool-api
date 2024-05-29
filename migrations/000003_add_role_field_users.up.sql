ALTER TABLE users
    ADD COLUMN role INT DEFAULT 0
        NOT NULL CHECK (role IN (0, 1, 2));