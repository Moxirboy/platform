package models

type AddTestRequest struct {
	QuestionID   string `json:"question_id"`
	QuestionText string `json:"question_text"`
	TeacherID    string `json:"teacher_id"`
	AnswerID     string `json:"answer_id"`
	ChoiceID     string `json:"choice_id"`
}

type CreateClass struct {
	ID        string `json:"id"`
	TeacherID string `json:"teacher_id"`
	ClassID   string `json:"class_id"`
	Period    string `json:"period"`
}
