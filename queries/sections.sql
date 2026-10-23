-- name: CreateSection :batchone
INSERT INTO sections (title, color, board_id)
VALUES (sqlc.arg(title), sqlc.arg(color), sqlc.arg(board_id))
RETURNING id, created_at;
