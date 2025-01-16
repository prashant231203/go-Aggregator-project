-- +goose Up
CREATE TABLE feeds(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id INTEGER NOT NULL
);

--+goose Down
DROP TABLE feeds;