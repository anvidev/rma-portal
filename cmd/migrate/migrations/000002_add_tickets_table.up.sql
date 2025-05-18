create type radio as enum ('yes', 'no', 'unknown');

create table if not exists tickets (
    id text primary key,
    seq_id bigserial not null unique,
    status integer not null,
    categories integer [],
    issue text not null,
    model citext,
    serial_number citext,
    quote radio not null,
    warranty radio not null,
    sender_name citext not null,
    sender_email citext not null,
    sender_phone citext not null,
    sender_street citext not null,
    sender_city citext not null,
    sender_zip citext not null,
    sender_country citext not null,
    billing_name citext not null,
    billing_email citext not null,
    billing_phone citext not null,
    billing_street citext not null,
    billing_city citext not null,
    billing_zip citext not null,
    billing_country citext not null,
    inserted timestamp(0) with time zone default now(),
    updated timestamp(0) with time zone default now()
);

create or replace function generate_ticket_id()
returns trigger as $$
begin
    new.id := lpad(new.seq_id::text, 5, '0') || '-' || 
              upper(substring(md5(random()::text), 1, 5));
    return new;
end;
$$ language plpgsql;

create trigger trg_generate_ticket_id
before insert on tickets
for each row execute function generate_ticket_id();
