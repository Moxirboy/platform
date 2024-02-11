package usecase
import (
	"context"
	"admin/internal/models"
)

type ITeacherUsecase interface {
	Create(ctx context.Context, teacher *models.User) error
}