CREATE TABLE IF NOT EXISTS "questions" (
  "id" uuid PRIMARY KEY,
  "questions_text" text NOT NULL,
  "teacher_id" uuid
);

ALTER TABLE "questions" ADD FOREIGN KEY ("teacher_id") REFERENCES "users" ("id");
