CREATE TABLE IF NOT EXISTS "classes" (
  "id" uuid PRIMARY KEY,
  "name" varchar(20),
  "studentCount" int,
  "student_id" uuid,
  "teacher_id" uuid
);
 ALTER TABLE "classes" ADD FOREIGN KEY ("teacher_id") REFERENCES "users" ("id");