package postgres

import (
	logger "admin/pkg/log"
	"admin/internal/service/repo"
	"context"
	"database/sql"
)

type classRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewclassRepository(
	db *sql.DB,
	log logger.Logger,
) repo.ClassRepository {
	return &classRepository{
		db:  db,
		log: log,
	}
}

func (r *classRepository) GetClassExist(ctx context.Context, className string) (bool, error) {
	number := 0
	err := r.db.QueryRowContext(
		ctx,
		GetClassExist,
		className,
	).Scan(
		&number,
	)
	if err != nil {
		return false, err
	}
	if number == 0 {
		return false, nil
	}
	return true, nil
}

func (r *classRepository) GetClassPassword(ctx context.Context, className string) (string, string, error) {
	var password string
	var id string
	err := r.db.QueryRowContext(
		ctx,
		GetClassPassword,
		className,
	).Scan(
		&password,
		&id,
	)
	if err != nil {
		return "", "", err
	}
	return password, id, nil
}
