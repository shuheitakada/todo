DROP TABLE sessions;
DROP TABLE tasks;
DROP TABLE users;

CREATE TABLE users (
  id               serial PRIMARY KEY,
  name             varchar(255),
  email            varchar(255) NOT NULL UNIQUE,
  crypted_password varchar(255) NOT NULL,
  salt             varchar(255) NOT NULL,
  created_at       timestamp NOT NULL,
  updated_at       timestamp NOT NULL
);

CREATE TABLE tasks (
  id          serial PRIMARY KEY,
  name        varchar(255) NOT NULL,
  description text,
  user_id     integer REFERENCES users(id) NOT NULL,
  created_at  timestamp NOT NULL,
  updated_at  timestamp NOT NULL
);

CREATE TABLE sessions (
  id         serial PRIMARY KEY,
  uuid       varchar(255) NOT NULL UNIQUE,
  user_id    integer REFERENCES users(id) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
)
