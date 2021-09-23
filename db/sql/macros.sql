DROP TABLE IF EXISTS "public"."macros";
CREATE TABLE "public"."macros" (
  "char_id" int4 NOT NULL,
  "id" int4 NOT NULL DEFAULT nextval('macros_id_seq'::regclass),
  "icon" int4,
  "name" varchar(40) COLLATE "pg_catalog"."default",
  "desc" varchar(80) COLLATE "pg_catalog"."default",
  "acronym" varchar(4) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."macros" ADD CONSTRAINT "macros_pkey" PRIMARY KEY ("id");
