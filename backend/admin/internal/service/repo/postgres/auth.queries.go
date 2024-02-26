package postgres

const (
	CreateUser = `
	insert into "users" (firstname,lastname,password,role,created_at,updated_at,verified) values ($1,$2,$3,$4,$5,$6,$7) returning id,role
	`

	GetExist = `
	SELECT EXISTS (
		SELECT 1
		FROM "users"
		WHERE firstname = $1 AND lastname = $2
	);
	`
	
	GetUser = `
	select id, password,role from "users" where firstname=$1 AND lastname=$2
	`
)
