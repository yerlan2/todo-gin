CREATE TABLE IF NOT EXISTS users (
    id            SERIAL       NOT NULL UNIQUE,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists (
    id      SERIAL       NOT NULL UNIQUE,
    name    VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    UNIQUE (name, user_id)
);

CREATE TABLE IF NOT EXISTS todo_items (
    id          SERIAL       NOT NULL UNIQUE,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(511) ,
    done        BOOLEAN      NOT NULL DEFAULT false,
    list_id     INT          NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);
