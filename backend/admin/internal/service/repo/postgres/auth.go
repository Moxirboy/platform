package postgres

import (
	logger "admin/pkg/log"
	domain "admin/internal/models"
	"admin/internal/service/repo"
	"context"
	"database/sql"
)

type authRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewAuthRepository(
	db *sql.DB,
	log logger.Logger,
) repo.AuthRepository {
	return &authRepository{
		db:  db,
		log: log,
	}
}

func (r *authRepository) GetExist(
	ctx context.Context,
	firstname string,
	lastname string,
) (
	bool,
	error,
) {
	var exist bool
	err := r.db.QueryRowContext(
		ctx,
		GetExist,
		firstname,
		lastname,
	).Scan(&exist)
	if err != nil {
		r.log.Error("repo.auth.getexist error : ", err)
		return false, err
	}
	
	return exist, nil
}

func (r *authRepository) GetUser(
	ctx context.Context,
	firstname string,
	lastname string,
	) (
	string,
	string,
	string,
	error,
) {
	var id, password, role string
	err := r.db.QueryRowContext(
		ctx,
		GetUser,
		firstname,
		lastname,
	).Scan(
		&id,
		&password,
		&role,
	)
	if err != nil {
		r.log.Error("repo.auth.GetUser error : ", err.Error())
		return "", "", "", err
	}
	return id, password, role, nil
}

func (r *authRepository) CreateUser(
	ctx context.Context,
	auth *domain.StudentAuth,
) (
	id string,
	role string,
	err error,
) {
	err = r.db.QueryRowContext(
		ctx,
		CreateUser,
		auth.Firstname,
		auth.Lastname,
		auth.Password,
		auth.Role,
		auth.Created_at,
		auth.Updated_at,
	).Scan(
		&id,
		&role,
	)
	if err != nil {
		return "", "", nil
	}
	return id, role, nil
}
