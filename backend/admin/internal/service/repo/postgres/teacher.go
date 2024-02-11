package postgres

import (
	"admin/internal/models"
	logger "admin/pkg/log"
	"context"
	"database/sql"
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

func (r *teacherRepo) Create(ctx context.Context, teacher *models.User) error {
	err := r.db.QueryRowContext(ctx, createTeacher, &teacher.Username, &teacher.Password, &teacher.Role, &teacher.CreatedAt, &teacher.UpdateAt, &teacher.Role).Scan(&teacher.ID)
	if err != nil {
		return err
	}
	return nil
}
