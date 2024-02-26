package usecase

import (
	"admin/internal/models"
	"admin/pkg/utils"
	"context"
	domain "admin/internal/models"

)

type ITeacherUseCase interface {
	Create(ctx context.Context, teacher *models.User) error
	GetAll(ctx context.Context, PaginationQuery utils.PaginationQuery) (*models.UserList, error)
}



type AuthUsecase interface {
	GetUser(
		ctx context.Context,
		firstname string,
		lastname string,
		password string,
	) (
		bool,
		string,
		string,
		error,
	)
	CreateUser(
		ctx context.Context,
		student *domain.StudentAuth,
	) (
		string,
		string,
		error,
	)
}

