CREATE TABLE IF NOT EXISTS reactions (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    char TEXT NOT NULL UNIQUE
);

INSERT INTO reactions (name, char) VALUES
    ('thumbsUp', '👍'),
    ('thumbsDown', '👎'),
    ('fire', '🔥');

CREATE TABLE IF NOT EXISTS messages_reactions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id   UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reaction_id INT NOT NULL REFERENCES reactions(id) ON DELETE CASCADE,
    message_id  UUID NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    created_at  TIMESTAMP DEFAULT clock_timestamp() NOT NULL,

    CONSTRAINT UC_message_reaction_author UNIQUE (message_id, reaction_id, author_id)
);