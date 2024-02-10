package usecase

import (
	"admin/internal/service/repo"
	logger "admin/pkg/log"
)

type UserUsecase struct{
	repo repo.IUserRepository
	log logger.Logger
}

func NewUserUsecase (
	repo repo.IUserRepository,
	log logger.Logger,
)IUserUsecase{
	return &UserUsecase{
		repo: repo,
		log:log,
	}
}