package postgres

const (
	CreateTest=`
	insert into "questions" ("questions_text", "teacher_id") values ($1, $2) returning "id"
	`
	CreateChoice= `
	insert into "choices" ("choice_text", "question_id","is_correct") values ($1, $2) returning "id"
	`

)