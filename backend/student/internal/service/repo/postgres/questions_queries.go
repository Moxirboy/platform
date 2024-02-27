package postgres

const (
	GetQuestions = `
	SELECT id, question_text 
	FROM question
	ORDER BY RAND()
	LIMIT 50;

	`
	GetChoices = `
	select id,choice_text from choices where question_id = $1
	`
	GetAnswer = `
	select choice_id from answer where question_id = $1 
	`
)