package postgres


const (
	CreateUser=`
	insert into "user" (username,password,role,created_at,updated_at,verified) values ($1,$2,$3,$4,$5,$6) returning id
	`

	GetExist =	`
	SELECT 
    CASE 
        WHEN EXISTS (SELECT * FROM your_table WHERE your_condition) THEN 1
        ELSE 0
    END AS result

	`
	GetUser = `
	select id, password,role from "user" where username=$1
	`

)