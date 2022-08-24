create table gameservers
(
    server_id integer default 0 not null,
    hex_id    varchar(50)       not null,
    host      varchar(50)       not null
);

alter table gameservers
    owner to postgres;

create unique index hex_id_unique
    on gameservers (hex_id);

