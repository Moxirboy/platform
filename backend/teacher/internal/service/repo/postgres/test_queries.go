package postgres

const (
	CreateTest=`
	insert into "questions" ("questions_text", "teacher_id") values ($1, $2) returning "id"
	`
	CreateChoice= `
	insert into "choices" ("choice_text", "question_id") values ($1, $2) returning "id"
	`
	CreateAnswer=`
	insert into "answers" ("question_id", "choice_id") values ($1, $2)
	`
)