CREATE TABLE IF NOT EXISTS outbox (
    id           SERIAL PRIMARY KEY,
    event_id     UUID NOT NULL DEFAULT gen_random_uuid(),
    topic        TEXT NOT NULL,
    key          TEXT NOT NULL,
    payload      JSONB NOT NULL,
    created_at   TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
    published_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_outbox_unpublished 
ON outbox (created_at) 
WHERE published_at IS NULL;
