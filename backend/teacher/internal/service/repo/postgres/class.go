package postgres

import (
	logger "teacher/pkg/log"
	"context"
	"database/sql"
	"github.com/google/uuid"
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
func (r *teacherRepo) CreateClass(ctx context.Context, class models.Class) (string,error){
	r.log.Info(class)
	err:=r.db.QueryRowContext(
		ctx,
		CreateClass,
		&class.TeacherID,
		&class.Name,
		&class.Password,
		).Scan(&class.ID)
		if err!=nil{
			r.log.Error("error is while creating class", err.Error())
			return "",err
		}
		return class.ID,nil
}

func (r *teacherRepo) GetClassIdByName(ctx context.Context, name string) (string, error) {
	var id string
	err:=r.db.QueryRowContext(
		ctx,
		GetClassIdByName,
		&name,
		).Scan(&id)
		if err!=nil{
			r.log.Error("error is while getting class name", err.Error())
			return "",err
		}
		return id,nil
}





func (t *teacherRepo) AddTest(ctx context.Context, request models.AddTestRequest) (string, error) {
	addTestQuery := `BEGIN;
			                INSERT INTO question (id, text, teacher_id) 
                            VALUES($1, $2, $3);
							INSERT INTO answer(id, question_id, choice_id)
							VALUES($4, $5, $6);
						COMMIT;`
	if _, err := t.db.ExecContext(ctx, addTestQuery,
		&request.QuestionID,
		&request.QuestionText,
		&request.TeacherID,
		&request.AnswerID,
		&request.QuestionID,
		&request.ChoiceID,
	); err != nil {
		t.log.Error("error is while adding test", err)
		return "", err
	}
	return request.QuestionID, nil
}



func (t *teacherRepo) StartTest(ctx context.Context, class models.CreateClass) (string, error) {
	id := uuid.New()
	createClassQuery := `insert into exam(id, teacher_id, class_id, period) 
                                       values($1, $2, $3, $4);`
	if _, err := t.db.ExecContext(ctx, createClassQuery,
		&id,
		&class.TeacherID,
		&class.ClassID,
		&class.Period,
	); err != nil {
		t.log.Error("error is while inserting exam", err.Error())
		return "", err
	}
	return id.String(), nil
}
