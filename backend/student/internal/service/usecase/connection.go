package usecase


type IStudentUsecase interface {
	StudentUsecase() IStudentUsecase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_studentUseCase = "student_use_case"
)

func New(
	db *sql.DB,
	log logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_studentUseCase] = NewStudentUseCase(
		postgres.NewStudentRepo(db, log),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) StudentUsecase() IStudentUsecase {
	return c.connections[_studentUseCase].(IStudentUsecase)
}


