package models

type Questions struct {
	ID           string
	ExamID       string
	QuestionText string
}
type TestResult struct {
	Result string
}
type Answers struct {
	ID         string
	QuestionID string
	ChoiceID   string
	IsCorrect  bool
}
type CreateAnswers struct{
	Answers []Answers
}

type QuestionsAll struct {
	Id string 
	QuestionText string
	Choices []Choices
}

type Choices struct {
	Id string
	ChoiceText string
	Is_correct bool
}