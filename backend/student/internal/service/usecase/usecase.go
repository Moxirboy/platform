package usecase

import (
	"admin/internal/models"
	"context"
)

type IStudentRepository interface {
	TakeExam(ctx context.Context, student models.Questions) (models.TestResult, error)
	GetTestResults(ctx context.Context, examID string) ([]models.TestResult, error)
}
