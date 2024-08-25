-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedById :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetFeedsByUser :many
SELECT * FROM feeds WHERE user_id = $1 LIMIT $2;

-- name: UpdateFeed :one
UPDATE feeds SET updated_at = $1, name = $2, url = $3 WHERE id = $4 AND user_id = $5
RETURNING *;

-- name: DeleteFeed :one
DELETE FROM feeds WHERE id = $1 AND user_id = $2
RETURNING *;