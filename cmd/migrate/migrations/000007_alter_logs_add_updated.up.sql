alter table logs
add column updated timestamp(0)
with
  time zone default null;

alter table logs
add column updated_by varchar(255) default null;

create or replace function set_updated()
returns trigger as $$
begin
	new.updated := now();
	return new;
end;
$$ language plpgsql;

create trigger trg_update_logs_updated before
update on logs for each row execute function set_updated();
