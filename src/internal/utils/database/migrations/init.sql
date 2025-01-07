CREATE TABLE statuses (
    id SERIAL PRIMARY KEY,
    title TEXT
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT,
    points INTEGER
);

CREATE TABLE referrers (
    id SERIAL PRIMARY KEY,
    code TEXT
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    active BOOLEAN NOT NULL,
    registered_at TIMESTAMPTZ NOT NULL,
    user_name TEXT NOT NULL,
    status_id INTEGER NOT NULL,
    last_visit TIMESTAMPTZ,
    balance BIGINT,
    isadmin BOOLEAN DEFAULT FALSE,
    referrer_id INTEGER,
    FOREIGN KEY (status_id) REFERENCES statuses(id) ON DELETE RESTRICT,
    FOREIGN KEY (referrer_id) REFERENCES referrers(id) ON DELETE SET NULL
);

CREATE TABLE tasks_activity (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    task_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);
