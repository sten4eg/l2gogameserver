alter table macros_commands
    add char_id integer not null;

alter table macros_commands
    add constraint macros_commands_characters_null_fk
        foreign key (char_id) references characters (object_id);