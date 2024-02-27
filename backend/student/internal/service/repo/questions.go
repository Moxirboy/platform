package repo
import (
	"context"
	"student/internal/models"
)
type QuestionRepo interface {
	GetQuestions(ctx context.Context, examId string) ([]models.QuestionsAll,error)
	GetChoices(ctx context.Context, questionId string) ([]models.Choices,error)
	GetAnswer(ctx context.Context, questionId string) (string,error) 
}