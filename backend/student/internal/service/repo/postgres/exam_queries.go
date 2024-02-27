package postgres

const (
	GetExist = `
	SELECT EXISTS (
		SELECT 1
		FROM "class_user"
		WHERE student_id = $1 
	);
	`
	GetclassId=`
	SELECT class_id
		FROM "class_user"
		WHERE student_id = $1 
	`
	
	GetExamId = `
	select id from exam where class_id = $1
	`
)