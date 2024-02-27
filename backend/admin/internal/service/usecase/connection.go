package usecase

import (
	"admin/internal/service/repo/postgres"
	logger "admin/pkg/log"
	"database/sql"
)

type IUseCase interface {
	TeacherUsecase() ITeacherUseCase
	AuthUseCase() AuthUsecase

}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_teacherUseCase = "teacher_use_case"
	_authUsecase = "auth_use_case"

)

func New(
	db *sql.DB,
	log logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_teacherUseCase] = NewTeacherUseCase(
		postgres.NewTeacherRepo(db, log),
		log,
	)
	connections[_authUsecase] = NewAuthUsecase(
		postgres.NewAuthRepository(
			db,
			log,
		),
		postgres.NewclassRepository(
			db,
			log,
		),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) TeacherUsecase() ITeacherUseCase {
	return c.connections[_teacherUseCase].(ITeacherUseCase)
}

func (u *UseCase) AuthUseCase() AuthUsecase {
	return u.connections[_authUsecase].(AuthUsecase)
}

