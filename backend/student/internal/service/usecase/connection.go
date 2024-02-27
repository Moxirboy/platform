package usecase

import (
	"database/sql"
	logger "student/pkg/log"
	"student/internal/service/repo/postgres"

)

type Usecase interface {
	// StudentUsecase() IStudentUsecase
	QuestionUsecase() IQuestionUsecase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	// _studentUseCase = "student_use_case"
	_questionUseCase = "question_use_case"
)

func New(
	db *sql.DB,
	log logger.Logger,
) Usecase {
	var connections = make(map[string]interface{})
	// connections[_studentUseCase] = NewStudentUseCase(
	// 	postgres.NewStudentRepo(db, log),
	// 	log,
	// )
	connections[_questionUseCase] = NewQuestionUseCase(
		postgres.NewQuestionRepo(db, log),
		postgres.NewExamRepo(db, log),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}
// func (c *UseCase) StudentUsecase() IStudentUsecase {
// 	return c.connections[_studentUseCase].(IStudentUsecase)
// }
func (c *UseCase) QuestionUsecase() IQuestionUsecase {
	return c.connections[_questionUseCase].(IQuestionUsecase)
}


