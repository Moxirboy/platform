package models

type AddTestRequest struct {
	QuestionID   string `json:"question_id"`
	QuestionText string `json:"question_text"`
	TeacherID    string `json:"teacher_id"`
	AnswerID     string `json:"answer_id"`
	ChoiceID     string `json:"choice_id"`
}

type Test struct {
	QuestionText string
	Choices      []*Choice
	Answer 	     string
}

type Choice struct {
	ID          string 
	ChoiceText  string 
	IsAnswer bool
}