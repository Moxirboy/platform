package usecase

import (
	"admin/internal/service/repo/postgres"
	logger "admin/pkg/log"
	"database/sql"
)


type IUseCase interface {
	IUserUsecase() IUserUsecase
}

type UseCase struct{
	connections map[string]interface{}
}

const (
	_userUseCase = "user_use_case"
)

func New(
	db *sql.DB,
	log logger.Logger,
) IUseCase{
	var connections = make(map[string]interface{})
	connections[_userUseCase] = NewUserUsecase(
		postgres.NewUserRepository(db, log),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}

func (u *UseCase) IUserUsecase() IUserUsecase {
	return u.connections[_userUseCase].(IUserUsecase)
}