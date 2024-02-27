package postgres

import (
	"teacher/internal/models"
	"context"
	"database/sql"
	logger "teacher/pkg/log"
	"teacher/internal/service/repo"
)
type testRepo struct {
	db *sql.DB
	log logger.Logger
}

func NewTestRepo(db *sql.DB, log logger.Logger) repo.TestRepository {
	return &testRepo{
		db: db,
		log: log,
	}
}

func (r *testRepo) CreateTest(
	ctx context.Context,
	Id string,
	model []*models.Test,
) error {
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
			},
			)
	if err != nil {
		r.log.Error("error is while creating test", err.Error())
		return err
	}
	var Questionid string
	var Choiceid string
	for _, test := range model {
		
		 err := tx.QueryRowContext(
			ctx,
			CreateTest,
			&test.QuestionText,
			&Id,
		).Scan(&Questionid)
		if err != nil {
			tx.Rollback()
			r.log.Error("error is while creating test", err.Error())
			return err
		}
		for _, choice := range test.Choices {
			err:=tx.QueryRowContext(
				ctx,
				CreateChoice,
				&choice.ChoiceText,
				&choice.IsAnswer,
			).Scan(&Choiceid)
			if err != nil {
				tx.Rollback()
				r.log.Error("error is while creating test", err.Error())
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		r.log.Error("error is while creating test", err.Error())
		return err
	}
	return nil
}