package postgres

import (
	"context"
	"database/sql"
	logger "student/pkg/log"
	"student/internal/service/repo"
)
type examRepo struct {
	db *sql.DB
	log logger.Logger
}

func NewExamRepo(db *sql.DB, log logger.Logger) repo.ExamRepository {
	return &examRepo{
		db: db,
		log: log,
	}
}

func (r *examRepo) GetExist(ctx context.Context, userId string) (bool, error) {
	var exist bool
	err := r.db.QueryRowContext(
		ctx,
		GetExist,
		&userId,
		).Scan(&exist)
		if err != nil {
			r.log.Error("repo.auth.getexist error : ", err)
			return false, err
		}
	
	return exist, nil
}

func (r *examRepo) GetExamId(ctx context.Context, userId string) (string, error) {
	var classId string
	err:= r.db.QueryRowContext(
		ctx,
		GetclassId,
		&userId,
		).Scan(&classId)
		if err != nil {
			r.log.Error("repo.auth.getclassid error : ", err)
			return "", err
		}
	var examId string
	err = r.db.QueryRowContext(
		ctx,
		GetExamId,
		&userId,
		).Scan(&classId)
		if err != nil {
			r.log.Error("repo.auth.getexamid error : ", err)
			return "", err
		}
	
	return examId, nil
}