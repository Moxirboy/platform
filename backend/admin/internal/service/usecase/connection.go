package usecase

import (
	logger "admin/pkg/log"
	"database/sql"
	"admin/internal/service/repo/postgres"
)


type IUseCase interface {
	TeacherUsecase() ITeacherUsecase
}
type UseCase struct {
	connections map[string]interface{} 
}

const(
	_teacherUseCase= "teacher_use_case"
)
func New(
	db *sql.DB,
	log logger.Logger,
)IUseCase{
	var connections = make(map[string]interface{})
	connections[_teacherUseCase]= NewTeacherUseCase(
		postgres.NewTeacherRepo(db,log),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) TeacherUsecase() ITeacherUsecase{
	return c.connections[_teacherUseCase].(ITeacherUsecase)
}
