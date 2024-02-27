package usecase

import (
	"database/sql"
	logger "teacher/pkg/log"
	"teacher/internal/service/repo/postgres"
)

type IUseCase interface {
	ITeacherUseCase() ITeacherUseCase
	IExamUseCase() IExamUseCase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_teacherUseCase = "teacher_use_case"
	_examUseCase    = "exam_use_case"
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
	connections[_examUseCase] = NewExamUseCase(
		postgres.NewExamRepo(db, log),
		postgres.NewTeacherRepo(db, log),
		log,
	)

	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) ITeacherUseCase() ITeacherUseCase {
	return c.connections[_teacherUseCase].(ITeacherUseCase)
}

func (c *UseCase) IExamUseCase() IExamUseCase {
	return c.connections[_examUseCase].(IExamUseCase)
}