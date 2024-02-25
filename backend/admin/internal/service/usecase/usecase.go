package usecase

import (
	"admin/internal/models"
	"admin/pkg/utils"
	"context"
)

type ITeacherUseCase interface {
	Create(ctx context.Context, teacher *models.User) error
	GetAll(ctx context.Context, PaginationQuery utils.PaginationQuery) (*models.UserList, error)
}
