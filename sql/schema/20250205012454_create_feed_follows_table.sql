-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_feed_follows_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    feed_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_feed_follows_feed FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
ALTER TABLE feed_follows
    DROP CONSTRAINT IF EXISTS fk_feed_follows_user,
    DROP CONSTRAINT IF EXISTS fk_feed_follows_feed;

DROP TABLE IF EXISTS feed_follows;

