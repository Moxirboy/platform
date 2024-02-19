package usecase

import (
	logger "admin/pkg/log"
	"auth/internal/service/repo/postgres"
	"database/sql"
)

type Usecase interface{
	AuthUseCase() AuthUsecase
}

type usecase struct {
	connections map[string]interface{}
}

const (
	_authUsecase="auth_use_case"
)
func New(
	// cfg 
	db *sql.DB,
	log logger.Logger,
) *usecase {
	var connections = make(map[string]interface{})
	connections[_authUsecase]= NewAuthUsecase(
		postgres.NewAuthRepository(
			db,
			log,
		),
		log,
	)
	return &usecase{
		connections: connections,
	}
}

func (u *usecase) AuthUseCase() AuthUsecase {
	return u.connections[_authUsecase].(AuthUsecase)
}