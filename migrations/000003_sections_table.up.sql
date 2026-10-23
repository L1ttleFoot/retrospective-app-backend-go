CREATE TABLE IF NOT EXISTS sections (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    color      TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    updated_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    board_id   UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE
);