create table if not exists tickets (
	id bigserial primary key,
	status integer not null,
	categories integer [],
	issue text not null,
	sender_name varchar(100) not null,
	sender_email citext not null,
	sender_address varchar(255) not null,
	inserted timestamp(0) with time zone default now(),
	updated timestamp(0) with time zone default now()
)
