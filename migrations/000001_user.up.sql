CREATE TYPE IF NOT EXISTS "user_role_enum" AS ENUM (
  'admin',
  'student',
  'teacher'
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY,
  "FirstName" varchar(20),
  "LastName" varchar(20),
  "Password" int NOT NULL,
  "user_role" user_role_enum NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "verified" BOOLEAN NOT NULL DEFAULT false
);