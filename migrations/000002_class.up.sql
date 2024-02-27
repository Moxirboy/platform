CREATE TABLE IF NOT EXISTS "class" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(20),
  "password" varchar(100),
  "teacher_id" VARCHAR(60)
);
 ALTER TABLE "class" ADD FOREIGN KEY ("teacher_id") REFERENCES "user" ("id");

CREATE TABLE IF NOT EXISTS "class_user" (
  "id" SERIAL PRIMARY KEY,
  "class_id" VARCHAR(60),
  "student_id"  VARCHAR(60)
);
 ALTER TABLE "class_user" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");
 ALTER TABLE "class_user" ADD FOREIGN KEY ("student_id") REFERENCES "user" ("id");
