CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL,
    title TEXT NOT NULL,
    done BOOLEAN DEFAULT false,
    CONSTRAINT "Tasks_pkey" PRIMARY KEY (id)
);