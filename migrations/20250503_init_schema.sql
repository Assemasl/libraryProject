-- +goose Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       is_author BOOLEAN DEFAULT false
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       description TEXT,
                       author_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE reactions (
                           id SERIAL PRIMARY KEY,
                           user_id INTEGER NOT NULL REFERENCES users(id),
                           book_id INTEGER NOT NULL REFERENCES books(id),
                           type VARCHAR(10) CHECK (type IN ('like', 'dislike')),
                           UNIQUE(user_id, book_id)
);

-- +goose Down
DROP TABLE IF EXISTS reactions;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;
