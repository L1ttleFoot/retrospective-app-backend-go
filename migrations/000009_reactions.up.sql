CREATE TABLE IF NOT EXISTS reactions (
    reaction TEXT UNIQUE PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS messages_reactions (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id  UUID REFERENCES users(id) ON DELETE CASCADE,
    reaction   TEXT REFERENCES reactions(reaction) ON DELETE CASCADE,
    message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL,

    CONSTRAINT UC_message_reaction_author UNIQUE (message_id, reaction, author_id)
);