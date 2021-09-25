CREATE TABLE "macros" (
  "char_id" integer NOT NULL,
  "id" serial   not null,
  "icon" integer,
  "name" varchar(40) default ''::character varying not null,
  "description" varchar(80) default ''::character varying not null,
  "acronym" varchar(4) default ''::character varying not null
);
