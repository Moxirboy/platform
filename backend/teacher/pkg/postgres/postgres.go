package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"sync"
	"teacher/internal/config"
)

var (
	instance *sql.DB
	once     sync.Once
)

// DB return database connection
func DB(cfg *configs.Postgres) (*sql.DB, error) {
	var err error

	psqlString := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
	instance, err = sql.Open("pgx", psqlString)
	// db, err := sql.Open("sqlite3", "storage/platform.db")
	if err != nil {
		return nil, errors.Wrap(err, "pgx.Connect")
	}

	return instance, nil
}
