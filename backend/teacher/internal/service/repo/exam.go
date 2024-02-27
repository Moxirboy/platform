package repo

import (
	"context"
	"teacher/internal/models"
)
type ExamRepository interface {
	CreateExam(ctx context.Context, exam  models.Exam) (string, error)
}