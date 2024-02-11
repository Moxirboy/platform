package postgres

import (
	"admin/internal/configs"
	"database/sql"
	"fmt"
	"sync"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/v4/stdlib" 
)

var (
	instance *sql.DB
	once     sync.Once
)

// DB return database connection
func DB(cfg *configs.Postgres) (*sql.DB, error) {
	var err error
	once.Do(func() {
		psqlString := fmt.Sprintf(
			`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Database,
		)
		instance, err = sql.Open("pgx", psqlString)
	})

	if err != nil {
		return nil, errors.Wrap(err, "pgx.Connect")
	}

	return instance, nil
}