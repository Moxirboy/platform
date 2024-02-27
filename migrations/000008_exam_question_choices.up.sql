CREATE TABLE question_choices (
  id SERIAL PRIMARY KEY,
  question_id VARCHAR(60) ,
  choice_text TEXT,
  is_correct BOOLEAN
  -- Add other columns as needed
);