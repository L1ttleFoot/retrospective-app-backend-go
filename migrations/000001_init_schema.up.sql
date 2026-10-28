CREATE TABLE IF NOT EXISTS roles (
    id   SERIAL PRIMARY KEY,
    role TEXT NOT NULL UNIQUE
);

INSERT INTO roles (role) VALUES
    ('ADMIN'),
    ('USER')
    ON CONFLICT (role) DO NOTHING;

CREATE TABLE IF NOT EXISTS users (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username   TEXT NOT NULL UNIQUE,
    password   TEXT NOT NULL,
    is_guest   BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT clock_timestamp() NOT NULL
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,

    PRIMARY KEY (user_id, role_id)
);