-- This is a sample migration.

create table people(
  id serial primary key,
  title varchar not null,
  content varchar not null
);

---- create above / drop below ----

drop table people;
