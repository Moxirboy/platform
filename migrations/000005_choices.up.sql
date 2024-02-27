CREATE TABLE IF NOT EXISTS "choices" (
  "id" SERIAL PRIMARY KEY,
  "question_id" VARCHAR(60),
  "choice_text" text NOT NULL
);

ALTER TABLE "choices" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");