create extension if not exists citext;

create table if not exists users (
  id bigserial primary key,
  name citext not null,
  email citext unique not null,
  hash bytea not null,
  is_active boolean not null default true,
  inserted timestamp(0)
  with
    time zone default now ()
)
