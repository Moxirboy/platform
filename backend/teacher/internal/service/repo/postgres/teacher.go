package postgres

import (
	logger "admin/pkg/log"
	"context"
	"database/sql"
	"teacher/internal/models"
)

type teacherRepo struct {
	db  *sql.DB
	log logger.Logger
}

func NewTeacherRepo(db *sql.DB, log logger.Logger) *teacherRepo {
	return &teacherRepo{
		db:  db,
		log: log,
	}
}

func (t *teacherRepo) AddTest(ctx context.Context, request models.AddTestRequest) error {
	addTestQuery := `BEGIN;
			                INSERT INTO question (id, text, teacher_id) 
                            VALUES($1, $2, $3);
							INSERT INTO answer(id, question_id, choice_id)
							VALUES($4, $5, $6);
						COMMIT;`
	if _, err := t.db.ExecContext(ctx, addTestQuery,
		request.QuestionID,
		request.QuestionText,
		request.TeacherID,
		request.AnswerID,
		request.QuestionID,
		request.ChoiceID,
	); err != nil {
		t.log.Error("error is while adding test", err)
		return err
	}
	return nil
}

func (t *teacherRepo) StartTest(ctx context.Context) error {
	return nil
}
func (t *teacherRepo) StudentStatistics(ctx context.Context) error {
	return nil
}
