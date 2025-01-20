-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE, -- Added data type for `url`
    feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
