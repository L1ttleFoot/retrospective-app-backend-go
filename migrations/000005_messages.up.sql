CREATE TABLE IF NOT EXISTS messages (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    text       text NOT NULL,
    author_id  UUID NOT NULL,
    created_at timestamp DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp DEFAULT clock_timestamp() NOT NULL,
    section_id uuid NOT NULL REFERENCES sections(id) ON DELETE CASCADE
);