package usecase

import (
	"context"
	"teacher/internal/models"
)

type ITeacherUseCase interface {
	AddTest(context.Context, models.AddTestRequest) (string, error)
	StartTest(context.Context, models.CreateClass) (string, error)
}
