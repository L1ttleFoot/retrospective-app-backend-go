CREATE TABLE IF NOT EXISTS templates (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    author_id  UUID REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS templates_sections (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    color      TEXT,
    templateId UUID REFERENCES templates(id) ON DELETE CASCADE
);