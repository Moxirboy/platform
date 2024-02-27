package postgres
import (
	"database/sql"
	logger "teacher/pkg/log"
	"context"
	"teacher/internal/models"
)
type examRepo struct {
	db *sql.DB
	log logger.Logger
}

func NewExamRepo(db *sql.DB, log logger.Logger) *examRepo {
	return &examRepo{
		db: db,
		log: log,
	}
}
func (r *examRepo) CreateExam(ctx context.Context, exam  models.Exam) (string, error) {

	r.log.Info("exam.TeacherID")


	err := r.db.QueryRowContext(
		ctx,
		CreateExam,
		&exam.TeacherID,
		&exam.ClassID,
	).Scan(
		&exam.ID,
	)
	if err != nil {
		r.log.Error("error is while creating exam", err.Error())
		return "", err
	}
	return exam.ID, nil
}