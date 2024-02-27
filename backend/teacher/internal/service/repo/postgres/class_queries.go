package postgres

const (
	// AddTestQuery is a query to add a test
	CreateClass=`
	insert into class (teacher_id,name,password) values($1,$2,$3) returning id
	`
	GetClassIdByName=`
	select id from class where LOWER(name)=LOWER($1)
	`
)
