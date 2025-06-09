drop trigger if exists trg_update_logs_updated on logs;

alter table logs
drop column updated;

alter table logs
drop column updated_by;
