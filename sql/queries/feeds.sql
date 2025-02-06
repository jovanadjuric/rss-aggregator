-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
  )
  RETURNING *;

-- name: GetFeeds :many
SELECT f.id as f_id, f.created_at as f_created_at, f.updated_at as f_updated_at, f.name as f_name, f.url as f_url, u.name as u_name FROM feeds f LEFT JOIN users u ON u.id = f.user_id ORDER BY f.created_at DESC;

-- name: GetFeed :one
SELECT f.id as f_id, f.created_at as f_created_at, f.updated_at as f_updated_at, f.name as f_name, f.url as f_url, u.name as u_name FROM feeds f LEFT JOIN users u ON u.id = f.user_id WHERE f.url = $1 LIMIT 1;

-- name: DeleteFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at = $2, last_fetched_at = $2 WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
