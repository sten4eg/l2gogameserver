CREATE TABLE "macros" (
  "char_id" integer NOT NULL,
  "id" integer NOT NULL DEFAULT nextval('macros_id_seq'::regclass),
  "icon" integer,
  "name" varchar(40) default ''::character varying not null,
  "desc" varchar(80) default ''::character varying not null,
  "acronym" varchar(4) default ''::character varying not null
);
ALTER TABLE "macros" ADD CONSTRAINT "macros_pkey" PRIMARY KEY ("id");
