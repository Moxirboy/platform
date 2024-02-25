package repo

import (
	"context"
	"teacher/internal/models"
)

type ITeacherRepository interface {
	AddTest(ctx context.Context, request models.AddTestRequest) error
}
