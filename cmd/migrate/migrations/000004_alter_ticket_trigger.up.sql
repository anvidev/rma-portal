drop trigger if exists trg_generate_ticket_id on tickets;

create or replace function generate_ticket_id()
returns trigger as $$
begin
    new.id := lpad(new.seq_id::text, 5, '0') || '-' || 
              upper(substring(md5(random()::text), 1, 10));
    return new;
end;
$$ language plpgsql;

create trigger trg_generate_ticket_id
before insert on tickets
for each row execute function generate_ticket_id();
