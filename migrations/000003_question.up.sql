
CREATE TABLE IF NOT EXISTS "questions" (
  "id" SERIAL PRIMARY KEY,
  "questions_text" text NOT NULL,
  "teacher_id" VARCHAR(60)
);

ALTER TABLE "questions" ADD FOREIGN KEY ("teacher_id") REFERENCES "user" ("id");
