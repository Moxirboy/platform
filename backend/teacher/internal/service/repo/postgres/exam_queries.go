package postgres

const (
	CreateExam = `
	insert into exam (teacher_id,class_id) values($1,$2) returning id
	`

)