package usecase

import (
	"context"
	"teacher/internal/models"
)

type ITeacherUseCase interface {
	AddTest(context.Context, models.AddTestRequest) error
}
