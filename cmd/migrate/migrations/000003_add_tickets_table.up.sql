create table if not exists tickets (
	id bigserial primary key,
	status integer not null,
	categories integer [],
	issue text not null,
	sender_name citext not null,
	sender_email citext not null,
	sender_address citext not null,
	sender_phone citext not null,
	billing_name citext not null,
	billing_email citext not null,
	billing_address citext not null,
	billing_phone citext not null,
	inserted timestamp(0) with time zone default now(),
	updated timestamp(0) with time zone default now()
)
