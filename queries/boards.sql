-- name: GetBoardById :one
SELECT id, title, owner_id, created_at
FROM boards
WHERE id=sqlc.arg(id);