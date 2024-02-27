package postgres

const (
	CreateUser = `
	insert into "user" (firstname,lastname,password,role,verified) values ($1,$2,$3,$4,$5) returning id
	`

	GetExist = `
	SELECT EXISTS (
		SELECT 1
		FROM "user"
		WHERE LOWER(firstname) = LOWER($1) AND LOWER(lastname) = LOWER($2)
	);
	`
	
	GetUser = `
	select id, password,role from "user" where firstname=$1 AND lastname=$2
	`
)
