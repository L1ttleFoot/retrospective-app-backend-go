CREATE TABLE IF NOT EXISTS boards (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    updated_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    owner_id   UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);