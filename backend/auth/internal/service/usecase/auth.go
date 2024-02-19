package usecase

import (
	logger "admin/pkg/log"
	"auth/internal/domain"
	"auth/internal/service/repo"
	errors "auth/pkg/errors"
	"context"
)

type authUsecase struct{
	repo repo.AuthRepository
	log logger.Logger
}

func NewAuthUsecase (
	repo repo.AuthRepository,
	log logger.Logger,
) AuthUsecase {
	return &authUsecase{
		repo: repo,
		log: log,
	}
}

func (u *authUsecase) GetUser(
	ctx context.Context,
	username string,
	password string,
) (
	bool,
	string,
	string,
	error,
	){
		exist,err :=u.repo.GetExist(ctx,username)
		if err!=nil{
			return false,"","",err
		}

		if !exist{
			return false,"","",errors.ErrUserNotFound
		}
		id,pass,role,err:=u.repo.GetUser(ctx,username)
		if err!=nil{
			return false,"","",err
		}
		if pass!=password{
			return false,"","",errors.ErrInvalidPassword
		}
		return true,id,role,nil
	}

func (u *authUsecase) CreateUser(
	ctx context.Context,
	auth *domain.Auth,
)(
	id string,
	err error,
){
	id,err=u.repo.CreateUser(ctx,auth)
	if err!=nil{
		return "" ,err
	}
	return id,nil
}