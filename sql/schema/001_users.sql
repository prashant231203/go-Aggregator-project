-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- If using pgcrypto for UUID
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
