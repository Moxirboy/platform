CREATE OR REPLACE FUNCTION insert_random_exam_questions()
RETURNS TRIGGER AS $$
DECLARE
    question_count INTEGER;
    random_offset INTEGER;
BEGIN
    -- Get the total count of questions available
    SELECT COUNT(*) INTO question_count FROM questions;
    
    -- Ensure there are at least 50 questions available
    IF question_count < 50 THEN
        RAISE EXCEPTION 'Not enough questions available in the questions table';
    END IF;
    
    -- Loop to insert 50 random questions
    FOR i IN 1..50 LOOP
        -- Generate a random offset to select a random question
        random_offset := floor(random() * question_count);
        
        -- Insert a random question into exam_question
        INSERT INTO exam_question (id, exam_id, question_text)
        SELECT uuid_generate_v4(), NEW.id, q.questions_text
        FROM questions q
        OFFSET random_offset
        LIMIT 1;
    END LOOP;

    -- Loop to insert choices for each inserted question
    FOR q_row IN SELECT * FROM exam_question WHERE exam_id = NEW.id LOOP
        INSERT INTO question_choices (id, question_id, choice_text)
        SELECT uuid_generate_v4(), q_row.id, c.choice_text
        FROM choices c
        WHERE c.question_id = q_row.id;
        
        -- Copy answers associated with the question to question_answers
        INSERT INTO question_answers (id, choice_id)
        SELECT uuid_generate_v4(), a.choice_id
        FROM answers a
        WHERE a.question_id = q_row.id;
    END LOOP;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER insert_random_exam_questions_trigger
AFTER INSERT ON exam
FOR EACH ROW
EXECUTE FUNCTION insert_random_exam_questions();
