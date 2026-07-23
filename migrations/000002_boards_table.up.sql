CREATE TABLE IF NOT EXISTS boards (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title      text NOT NULL,
    created_at timestamp DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp DEFAULT clock_timestamp() NOT NULL,
    owner_id   uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);