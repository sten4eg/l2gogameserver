DROP TABLE IF EXISTS "public"."macros_commands";
CREATE TABLE "public"."macros_commands" (
  "command_id" int4 NOT NULL,
  "index" int4,
  "type" int4,
  "skill_id" int4,
  "shortcut_id" int4,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;
