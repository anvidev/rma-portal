alter table tickets
add column sender_company citext not null default '';

alter table tickets
add column billing_company citext not null default '';
