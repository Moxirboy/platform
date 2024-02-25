CREATE TABLE users (
                       id uuid PRIMARY KEY,
                       FullName varchar(30),
                       Password int NOT NULL,
                       user_role ENUM ('admin', 'student', 'teacher') NOT NULL,
                       created_at timestamp DEFAULT (now()),
                       updated_at timestamp DEFAULT (now()),
                       verified BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE classes (
                         id uuid PRIMARY KEY,
                         name varchar(20),
                         password varchar(20),
                         student_id uuid,
                         teacher_id uuid,
                         FOREIGN KEY (student_id) REFERENCES users (id),
                         FOREIGN KEY (teacher_id) REFERENCES users (id)
);

CREATE TABLE questions (
                           id uuid PRIMARY KEY,
                           questions_text text NOT NULL,
                           teacher_id uuid,
                           FOREIGN KEY (teacher_id) REFERENCES users (id)
);

CREATE TABLE choices (
                         id uuid PRIMARY KEY,
                         question_id uuid,
                         choice_text text NOT NULL,
                         FOREIGN KEY (question_id) REFERENCES questions (id)
);

CREATE TABLE answers (
                         id uuid PRIMARY KEY,
                         question_id uuid,
                         choice uuid,
                         FOREIGN KEY (question_id) REFERENCES questions (id),
                         FOREIGN KEY (choice) REFERENCES choices (id)
);

CREATE TABLE exam (
                      id uuid PRIMARY KEY,
                      teacher_id uuid,
                      class_id uuid,
                      period timestamp,
                      FOREIGN KEY (teacher_id) REFERENCES users (id),
                      FOREIGN KEY (class_id) REFERENCES classes (id)
);

CREATE TABLE exam_question (
                               id uuid PRIMARY KEY,
                               exam_id uuid,
                               question_text text,
                               FOREIGN KEY (exam_id) REFERENCES exam (id)
);

CREATE TABLE question_choices (
                                  id uuid PRIMARY KEY,
                                  question_id uuid,
                                  choice_text text,
                                  FOREIGN KEY (question_id) REFERENCES exam_question (id)
);

CREATE TABLE question_answer (
                                 id uuid PRIMARY KEY,
                                 choice_id uuid,
                                 FOREIGN KEY (choice_id) REFERENCES question_choices (id)
);

CREATE TABLE student_answers (
                                 id uuid PRIMARY KEY,
                                 student_id uuid,
                                 question_id uuid,
                                 choice text,
                                 is_correct bool,
                                 FOREIGN KEY (student_id) REFERENCES users (id),
                                 FOREIGN KEY (question_id) REFERENCES exam_question (id)
);
