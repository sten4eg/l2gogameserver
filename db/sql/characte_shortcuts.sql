create table character_shortcuts
(
    char_id int not null,
    slot int not null,
    page int not null,
    type int,
    shortcut_id int,
    level int,
    class_index int default 0 not null,
    constraint character_shortcuts_pk
        primary key (char_id, slot, page, class_index)
);
create index character_shortcuts_pk_2
    on character_shortcuts using btree (shortcut_id);

