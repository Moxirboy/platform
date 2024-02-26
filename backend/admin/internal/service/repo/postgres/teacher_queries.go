package postgres

const (
	createTeacher = `
	INSERT INTO "users" (firstname,lastname, password, role, created_at, updated_at, verified)
	VALUES ($1, $2, $3, $4, $5, $6,$7)
	RETURNING id;
		`
	GetAll = `
	select * from "users" OFFSET $1
	LIMIT $2
	`
	GetById = `
	select * from "users" where id=$1
	`
	DeleteTeacher = `
	Update "users" set deletedat=$1 where firstname=$2, lastname=$3
	`
	countID = `
	select count(id) from "users" where 1=1
	`
)
