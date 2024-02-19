package postgres

const (
	createTeacher=`
	INSERT INTO "user" (username, password, role, created_at, updated_at, verified)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;
		`
	GetAll=`
	select * from "user" OFFSET $1
	LIMIT $2
	`
	GetById=`
	select * from "user" where id=$1
	`
	DeleteTeacher=`
	Update "user" set deletedat=$1 where username=$2
	`
	countID = `
	select count(id) from "user" where 1=1
	`
)