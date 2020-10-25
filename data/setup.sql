DROP TABLE tasks;
DROP TABLE sessions;
DROP TABLE users;

CREATE TABLE tasks (
  id          serial PRIMARY KEY,
  name        varchar(255) NOT NULL,
  description text,
  created_at  timestamp NOT NULL,
  updated_at  timestamp NOT NULL
);

INSERT INTO tasks (name, description, created_at, updated_at) VALUES ('タスク1', 'ここにタスク1の説明が入ります', now(), now());
INSERT INTO tasks (name, description, created_at, updated_at) VALUES ('タスク2', 'ここにタスク2の説明が入ります', now(), now());
INSERT INTO tasks (name, description, created_at, updated_at) VALUES ('タスク3', 'ここにタスク3の説明が入ります', now(), now());

CREATE TABLE users (
  id               serial PRIMARY KEY,
  name             varchar(255),
  email            varchar(255) NOT NULL UNIQUE,
  crypted_password varchar(255) NOT NULL,
  salt             varchar(255) NOT NULL,
  created_at       timestamp NOT NULL,
  updated_at       timestamp NOT NULL
);

CREATE TABLE sessions (
  id         serial PRIMARY KEY,
  uuid       varchar(255) NOT NULL UNIQUE,
  user_id    integer REFERENCES users(id),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
)
