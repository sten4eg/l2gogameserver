create table accounts
(
	login varchar(45) not null,
	password varchar(65) not null,
	created_at timestamp default current_timestamp not null,
	last_active timestamp default null,
	access_level smallint default 0 not null,
	last_ip varchar(15) default null,
	last_server smallint default 1 not null
);

create unique index accounts_login_uindex
	on accounts (login);

alter table accounts
	add constraint accounts_pk
		primary key (login);

