-- name: GetReportSections :many
SELECT
    s.id,
    s.title,
    s.color,
    COUNT(m.id)::int AS messages_count,
    COUNT(m.pinned_at)::int AS pinned_count
FROM sections s
LEFT JOIN messages m ON s.id = m.section_id
WHERE s.board_id = sqlc.arg(board_id)
GROUP BY s.id, s.title, s.color, s.created_at
ORDER BY s.created_at;

-- name: GetReportReactions :many
SELECT
    r.char,
    r.name,
    COUNT(mr.id)::int AS usage_count
FROM sections s
JOIN messages m ON s.id = m.section_id
JOIN messages_reactions mr ON m.id = mr.message_id
JOIN reactions r ON mr.reaction_id = r.id
WHERE s.board_id = sqlc.arg(board_id)
GROUP BY r.id, r.char, r.name
ORDER BY usage_count DESC;

-- name: GetReportTopMessages :many
SELECT
    m.id AS message_id,
    s.title AS section_title,
    m.text,
    COUNT(mr.id)::int AS reactions_count
FROM sections s
JOIN messages m ON s.id = m.section_id
LEFT JOIN messages_reactions mr ON m.id = mr.message_id
WHERE s.board_id = sqlc.arg(board_id)
GROUP BY m.id, s.title, m.text
HAVING COUNT(mr.id) > 0
ORDER BY reactions_count DESC
LIMIT 5;
