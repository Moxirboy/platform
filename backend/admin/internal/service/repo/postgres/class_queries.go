package postgres

const (
	GetClassExist = `
	SELECT 
    CASE 
        WHEN EXISTS (SELECT * FROM "class" WHERE class_name=$1 ) THEN 1
        ELSE 0
    END AS result
	`
	GetClassPassword = `
	SELECT password FROM "class" where class_name=$1 returning id
	`
)
