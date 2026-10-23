-- name: CreateMessage :one
INSERT INTO messages (text, author_id, section_id)
VALUES (sqlc.arg(text), sqlc.arg(author_id), sqlc.arg(section_id))
RETURNING id, text, author_id, created_at, updated_at, pinned_at, section_id;

-- name: SetReaction :exec
WITH
target_reaction AS (
    SELECT id FROM reactions WHERE name = sqlc.arg(reaction_name)
),
deleted AS (
    DELETE FROM messages_reactions
    WHERE reaction_id = (SELECT id FROM target_reaction)
    AND message_id = sqlc.arg(message_id)
    AND author_id = sqlc.arg(author_id)
    RETURNING 1
)
INSERT INTO messages_reactions (reaction_id, message_id, author_id)
SELECT (SELECT id FROM target_reaction), sqlc.arg(message_id), sqlc.arg(author_id)
WHERE EXISTS (SELECT 1 FROM target_reaction)
AND NOT EXISTS (SELECT 1 FROM deleted);

-- name: GetMessagesBySectionID :many
SELECT
    m.id,
    m.text,
    m.author_id,
    m.created_at,
    m.updated_at,
    m.pinned_at,
    m.section_id,
    COALESCE(
        (
            SELECT json_agg(
                json_build_object(
                    'name', sub.name,
                    'char', sub.char,
                    'count', sub.cnt,
                    'isSelected', sub.is_selected
                ) ORDER BY sub.cnt DESC, sub.name ASC
            )
            FROM (
                SELECT
                    r.name,
                    r.char,
                    COUNT(*) AS cnt,
                    COALESCE(BOOL_OR(mr.author_id = sqlc.narg(user_id)::uuid), false) AS is_selected
                FROM messages_reactions mr
                JOIN reactions r ON mr.reaction_id = r.id
                WHERE mr.message_id = m.id
                GROUP BY r.name, r.char
            ) sub
        ), '[]'::json
    )::json AS reactions
FROM messages m
WHERE m.section_id = sqlc.arg(section_id)
ORDER BY
    (m.pinned_at IS NULL) ASC,
    m.pinned_at ASC,
    m.updated_at ASC;
