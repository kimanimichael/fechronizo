-- name: CreateFeeds :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, user_name)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM feeds;