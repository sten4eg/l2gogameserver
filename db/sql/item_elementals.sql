create table item_elementals
(
    object_id     integer default 0,
    element_type  integer default '-1'::integer not null,
    element_value integer default '-1'::integer
);