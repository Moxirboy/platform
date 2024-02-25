CREATE TABLE IF NOT EXISTS "choices" (
  "id" uuid PRIMARY KEY,
  "question_id" uuid,
  "choice_text" text NOT NULL
);

ALTER TABLE "choices" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");