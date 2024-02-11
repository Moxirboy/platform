package usecase

import (
	"context"
	"admin/internal/models"
	"admin/internal/service/repo"
	logger "admin/pkg/log"
)


type TeacherUsecase struct {
	repo repo.ITeacherRepository
	log logger.Logger
}

func NewTeacherUseCase(repo repo.ITeacherRepository, log logger.Logger) ITeacherUsecase {
	return &TeacherUsecase{
		repo:repo,
		log:log,
	}
}

func (u *TeacherUsecase) Create(ctx context.Context, teacher *models.User) error {
	err:=u.repo.Create(ctx,teacher)
	if err!=nil{
		u.log.Errorf(err.Error())
	}
	return nil
}