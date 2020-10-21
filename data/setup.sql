drop table tasks;
drop table users;

create table tasks (
  id          serial primary key,
  name        varchar(255) not null,
  description text,
  created_at  timestamp not null,
  updated_at  timestamp not null
);

create table users (
  id serial primary key,
  name varchar(255),
  email varchar(255) not null,
  crypted_password varchar(255) not null,
  salt varchar(255) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);
