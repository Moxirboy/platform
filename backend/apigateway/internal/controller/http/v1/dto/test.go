package dto

type Question struct {
	ID                string      `json:"id"`
	Text              string   `json:"text"`
	Choices           []*Choices 
	CorrectChoiceIndex string      `json:"correctChoiceIndex"`
}

type Choices struct {
	Id 			  string      `json:"id"`
	ChoiceText    string   `json:"choiceText"`
	IsCorrect bool `json:"is_correct"`
}

type Answer struct {
	QuestionID  int `json:"questionId"`
	AnswerIndex int `json:"answerIndex"`
	Answer      int `json:"answer"`
}

type CreateTestRequest struct {
	QuestionText string `json:"questionText"`
	Choices      []Choice
	Answer 	     string	`json:"answer"`
}
type Choice struct {
	ID          string 
	ChoiceText  string `json:"choiceText"`
	IsAnswer bool `json:"isAnswer"`
}
