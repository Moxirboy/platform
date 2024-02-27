package repo

import (
	"student/internal/models"
	"context"
)

type IStudentRepository interface {
	TakeExamQuestions(ctx context.Context, examID string) ([]models.Questions, error)
	TakeExam(ctx context.Context, student models.Questions) (models.TestResult, error)
	GetTestResults(ctx context.Context, examID string) ([]models.TestResult, error)
}

