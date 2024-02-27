package usecase

import (
	logger "admin/pkg/log"
	domain "admin/internal/models"
	"admin/internal/service/repo"
	errors "admin/pkg/errors"
	"context"
)

type authUsecase struct {
	auth  repo.AuthRepository
	class repo.ClassRepository
	log   logger.Logger
}

func NewAuthUsecase(
	auth repo.AuthRepository,
	class repo.ClassRepository,
	log logger.Logger,
) AuthUsecase {
	return &authUsecase{
		auth:  auth,
		class: class,
		log:   log,
	}
}

func (u *authUsecase) GetUser(
	ctx context.Context,
	firstname string,
	lastname string,
	password string,
) (
	bool,
	string,
	string,
	error,
) {
	exist, err := u.auth.GetExist(ctx, firstname,lastname)
	if err != nil {
		return false, "", "", err
	}

	if !exist {
		return false, "", "", errors.ErrUserNotFound
	}
	id, pass, role, err := u.auth.GetUser(ctx, firstname,lastname)
	if err != nil {
		return false, "", "", err
	}
	if pass != password {
		return false, "", "", errors.ErrInvalidPassword
	}
	return true, id, role, nil
}

func (u *authUsecase) CreateUser(
	ctx context.Context,
	student *domain.StudentAuth,
) (
	string,
	string,
	error,
) {
	var classid string
	exist, err := u.class.GetClassExist(ctx, student.Class)
	if err != nil {
		return "", "", err
	}
	if !exist {
		return "", "", errors.ErrClassNotFound
	}
	var pass string
	pass, classid, err = u.class.GetClassPassword(ctx, student.Class)
	u.log.Info(pass)
	if err != nil {
		return "", "", err
	}
	if pass != student.Password {
		return "", "", errors.ErrInvalidCredentials
	}
	student.Role="student"
	u.log.Info(student)
	id,  err := u.auth.CreateUser(ctx, student)
	if err != nil {
		return "", "", err
	}

	err = u.class.LinkClassUser(ctx, classid, id)
	if err != nil {
		return "", "", err
	}
	return id, student.Role, nil
}
