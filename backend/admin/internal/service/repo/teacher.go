package repo

import (
	"admin/internal/models"
	"admin/pkg/utils"
	"context"
)

type ITeacherRepository interface {
	Create(ctx context.Context, teacher *models.User) error
	GetAll(ctx context.Context, PaginationQuery utils.PaginationQuery) (*models.UserList, error)
}
