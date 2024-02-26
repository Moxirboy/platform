package models

type Admin struct {
	ID       string
	Username string
	Password string
}

type User struct {
	ID        string
	Firstname  string
	Lastname  string
	Password  string
	Role      string
	CreatedAt string
	UpdateAt  string
	DeletedAt string
	Verified  bool
}

type UserList struct {
	TotalCount int
	TotalPages int
	Page       int
	Size       int
	HasMore    bool
	Users      []*User
}
