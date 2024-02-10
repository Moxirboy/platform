package postgres

import (
	"database/sql"
	"log"
	"platform/admins/internal/service/repo"
)

type userRepository struct {
	db *sql.DB
	log log.Logger
}

func NewUserRepository(
	db *sql.DB,
	log log.Logger,
) repo.IUserRepository {
	return &userRepository{
		db: db,
		log: log,
	}
}
