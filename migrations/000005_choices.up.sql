CREATE TABLE "choices" (
    "id" SERIAL PRIMARY KEY,
    "question_id" VARCHAR(60) NOT NULL,
    "text" TEXT NOT NULL
);
ALTER TABLE "choices" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");