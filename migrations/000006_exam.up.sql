CREATE TABLE exam (
  id SERIAL PRIMARY KEY,
  teacher_id VARCHAR(60),
  class_id VARCHAR(60) ,
  period TIMESTAMP
  -- Add other columns as needed
);