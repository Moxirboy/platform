package usecase

import (
	"admin/internal/models"
	"admin/internal/service/repo"
	logger "admin/pkg/log"
	"admin/pkg/utils"
	"context"
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
		return err
	}
	return nil
}
func (u *TeacherUsecase) 	GetAll(ctx context.Context,PaginationQuery utils.PaginationQuery) (*models.UserList,error){
	list,err:=u.repo.GetAll(ctx,PaginationQuery)
	if err!=nil{
		u.log.Errorf(err.Error())
		return nil,err
	}
	return list,nil
}