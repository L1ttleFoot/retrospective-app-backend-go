CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at timestamp DEFAULT clock_timestamp() NOT NULL
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    role_id SERIAL REFERENCES roles(id) ON DELETE CASCADE,

    PRIMARY KEY (user_id, role_id)
)