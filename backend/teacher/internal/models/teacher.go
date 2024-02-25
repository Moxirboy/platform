package models

type AddTestRequest struct {
	QuestionID   int    `json:"question_id"`
	QuestionText string `json:"question_text"`
	TeacherID    int    `json:"teacher_id"`
	AnswerID     int    `json:"answer_id"`
	ChoiceID     int    `json:"choice_id"`
}

type CreateClass struct {
}
