drop table tasks;

create table tasks (
  id          serial primary key,
  name        varchar(255) not null,
  description text,
  created_at  timestamp not null,
  updated_at  timestamp not null
);
