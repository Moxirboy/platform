package repo

import (
	"context"
)
type ExamRepository interface {
	GetExamId(ctx context.Context, userId string) (string, error) 
	GetExist(ctx context.Context, userId string) (bool, error)
}
