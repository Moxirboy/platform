package postgres

const (
	createTeacher=`
	insert into user (username, password, role,created_at, updated_at, verified ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	
)