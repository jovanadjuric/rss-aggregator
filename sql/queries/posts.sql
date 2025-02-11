-- name: CreatePost :one
INSERT INTO posts (
  id,
  created_at,
  updated_at,
  published_at,
  title,
  url,
  description,
  feed_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetPostsByUserId :many
SELECT p.id as p_id, p.created_at as p_created_at, p.updated_at as p_updated_at, p.published_at as p_published_at, p.title as p_title, p.url as p_url, p.description as p_description, f.id as f_id, f.name as f_name, f.user_id as f_user_id FROM posts p INNER JOIN feeds f on p.feed_id = f.id INNER JOIN feed_follows ff on ff.feed_id = f.id WHERE ff.user_id = $1 ORDER BY p.created_at DESC LIMIT $2;
