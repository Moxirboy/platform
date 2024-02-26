
CREATE TABLE "answers" (
    "id" SERIAL PRIMARY KEY,
    "question_id" VARCHAR(60) NOT NULL,
    "choice_id" VARCHAR(60) NOT NULL,
);

ALTER TABLE "answers" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");
ALTER TABLE "answers" ADD FOREIGN KEY ("choice_id") REFERENCES "choices" ("id");

