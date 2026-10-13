CREATE TABLE IF NOT EXISTS messages (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text       TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    updated_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    section_id UUID NOT NULL REFERENCES sections(id) ON DELETE CASCADE,
    author_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);