package postgres

const (
	GetClassExist = `
	SELECT 
    CASE 
        WHEN EXISTS (SELECT * FROM "class" WHERE name=$1 ) THEN 1
        ELSE 0
    END AS result
	`
	GetClassPassword = `
	SELECT id, password FROM "class" WHERE LOWER(name) = LOWER($1);
	`
	LinkClassUser=`
	insert into class_user (class_id,student_id) values($1,$2)
	`
)
