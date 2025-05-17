create table if not exists logs (
  id bigserial primary key,
  ticket_id text references tickets (id) on delete cascade,
  status integer not null,
  initiator varchar(100) not null,
  external_comment text,
  internal_comment text,
  inserted timestamp(0)
  with
    time zone default now ()
)
