CREATE TABLE IF NOT EXISTS sections (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title      text NOT NULL,
    color      text,
    created_at timestamp DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp DEFAULT clock_timestamp() NOT NULL,
    board_id   uuid NOT NULL REFERENCES on boards(id) ON DELETE CASCADE,
)