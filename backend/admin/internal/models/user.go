package models

type Admin struct {
	ID string
	Username string	
	Password string
}

type User struct {
	ID string
	Username string
	Password string
	Role string
	CreatedAt string
	UpdateAt string
	DeletedAt string
	Verified bool
}

