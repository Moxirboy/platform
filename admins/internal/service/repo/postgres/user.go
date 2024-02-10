package postgres

import (
	"admin/internal/service/repo"
	logger "admin/pkg/log"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
	log logger.Logger
}

func NewUserRepository(
	db *sql.DB,
	log logger.Logger,
) repo.IUserRepository {
	return &userRepository{
		db: db,
		log: log,
	}
}
