CREATE TABLE exam_question (
  id SERIAL PRIMARY KEY,
  exam_id VARCHAR(60) ,
  question_text TEXT
  -- Add other columns as needed
);