package models

type Auth struct {
	Id         string
	Firstname  string
	Lastname   string
	Password   string
	Role       string
	Created_at string
	Updated_at string
	Verified   string
}
type StudentAuth struct {
	ID         string
	Firstname  string
	Lastname   string
	Password   string
	Class      string
	Role       string
	Created_at string
	Updated_at string
}
