create type files_domain as enum ('tickets', 'logs');

create table if not exists files (
  id bigserial primary key,
  file_name varchar(255) not null,
  file_url varchar(255) not null,
  file_domain files_domain not null,
  reference_id varchar(255) not null,
  mime_type varchar(100) not null,
  inserted timestamp(0)
  with
    time zone default now () not null
);

create index if not exists idx_files_reference_id on files (file_domain, reference_id);
