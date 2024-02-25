CREATE TABLE IF NOT EXISTS "answers" (
  "id" uuid PRIMARY KEY,
  "question_id" uuid,
  "choice_id" uuid,
  "is_correct" boolean
);


ALTER TABLE "answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "answers" ADD FOREIGN KEY ("choice") REFERENCES "choices" ("id");