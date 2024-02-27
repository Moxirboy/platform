package usecase


import (
	"student/internal/models"
	"context"
)

type IStudentUsecase interface {
	TakeExam(ctx context.Context, student models.Questions) (models.TestResult, error)
	GetTestResults(ctx context.Context, examID string) ([]models.TestResult, error)
}

type IQuestionUsecase interface {
	GetAllQuestions(ctx context.Context, examId string) ([]models.QuestionsAll, error)
	GetExamExist(ctx context.Context, userId string) (bool,string ,error)
}