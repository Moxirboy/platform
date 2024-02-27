package models


type CreateClass struct {
	ID        string `json:"id"`
	TeacherID string `json:"teacher_id"`
	ClassID   string `json:"class_id"`
	Period    string `json:"period"`
}
type Choices struct {
	Id string
	Choice string
}

type Class struct {
	ID string
	TeacherID string
	Name string
	Password string
	Period string
}