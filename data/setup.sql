DROP TABLE tasks;
DROP TABLE users;

CREATE TABLE tasks (
  id          serial PRIMARY KEY,
  name        varchar(255) NOT NULL,
  description text,
  created_at  timestamp NOT NULL,
  updated_at  timestamp NOT NULL
);

CREATE TABLE users (
  id               serial PRIMARY KEY,
  name             varchar(255),
  email            varchar(255) NOT NULL,
  crypted_password varchar(255) NOT NULL,
  salt             varchar(255) NOT NULL,
  created_at       timestamp NOT NULL,
  updated_at       timestamp NOT NULL
);
