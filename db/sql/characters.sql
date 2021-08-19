-- auto-generated definition
create table characters
(
    login       bytea                  not null,
    id          serial                 not null
        constraint table_name_pk
            primary key,
    level       smallint default 1     not null,
    max_hp      integer  default 100   not null,
    cur_hp      integer  default 100   not null,
    max_mp      integer  default 100   not null,
    cur_mp      integer  default 100   not null,
    face        smallint               not null,
    hair_style  smallint               not null,
    hair_color  smallint               not null,
    sex         smallint               not null,
    x           integer                not null,
    y           integer,
    z           integer,
    exp         integer  default 0     not null,
    sp          integer  default 0     not null,
    karma       integer  default 0     not null,
    pvp_kills   integer  default 0     not null,
    pk_kills    integer  default 0     not null,
    clan_id     integer  default 0     not null,
    race        smallint               not null,
    class_id    integer                not null,
    base_class  integer  default 0     not null,
    title       bytea,
    online_time integer  default 0     not null,
    nobless     integer  default 0     not null,
    vitality    integer  default 20000 not null,
    char_name   bytea                  not null
);

alter table characters
    owner to postgres;

create unique index table_name_id_uindex
    on characters (id);

