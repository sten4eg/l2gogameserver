create table character_skills
(
    char_id int default 0 not null,
    skill_id int default 0 not null,
    skill_level int default 1 not null,
    class_id int default 0 not null
);

create index character_skills_char_id_skill_id_class_index_index
    on character_skills (char_id, skill_id, class_id)