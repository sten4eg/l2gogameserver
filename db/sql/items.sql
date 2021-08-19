create table items
(
    owner_id int not null,
    object_id int not null,
    item int not null,
    count int default 1 not null,
    enchant_level int default 0 not null,
    loc varchar(10) default null,
    loc_data int,
    time_of_use int,
    custom_type1 int default 0,
    custom_type2 int default 0,
    mana_left decimal default -1,
    time int default 0,
    agathion_energy int default 0
);

