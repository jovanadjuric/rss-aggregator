-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    published_at TIMESTAMP,
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_posts_feed FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
ALTER TABLE posts
    DROP CONSTRAINT IF EXISTS fk_posts_feed;
DROP TABLE IF EXISTS posts;

