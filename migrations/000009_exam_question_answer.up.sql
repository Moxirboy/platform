CREATE TABLE question_answer (
  id SERIAL PRIMARY KEY,
  choice_id VARCHAR(60) ,
  question_id VARCHAR(60) ,
  -- Add other columns as needed
);