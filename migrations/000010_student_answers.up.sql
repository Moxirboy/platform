CREATE TABLE student_answers (
  id SERIAL PRIMARY KEY,
  student_id VARCHAR(60) REFERENCES users(id),
  question_id VARCHAR(60) REFERENCES exam_question(id),
  choice TEXT,
  is_correct BOOLEAN
  -- Add other columns as needed
);