package postgres

import (
	"context"
	"database/sql"
	"student/internal/models"
	"student/internal/service/repo"
	logger "student/pkg/log"
)

type questionRepo struct{
	db  *sql.DB
	log logger.Logger
}

func NewQuestionRepo(db *sql.DB, log logger.Logger ) repo.QuestionRepo{
	return &questionRepo{
		db:db,
		log:log,
	}
}

func (r *questionRepo) GetQuestions(ctx context.Context, examId string) ([]models.QuestionsAll,error) {

	questions:=[]models.QuestionsAll{}

	rows,err:=r.db.QueryContext(
		ctx,
		GetQuestions,
	)
	if err!=nil{
		r.log.Errorf(err.Error())
		return nil,err
	}
	defer rows.Close()

	for rows.Next(){
		var question models.QuestionsAll
		err:=rows.Scan(
			&question.Id,
			&question.QuestionText,
		)
		if err!=nil{
			r.log.Errorf(err.Error())
			return nil,err
		}
		questions=append(questions,question)
	}
	return questions,nil
}

func (r *questionRepo) GetChoices(ctx context.Context, questionId string) ([]models.Choices,error) {

	Choices:=[]models.Choices{}

	rows,err:=r.db.QueryContext(
		ctx,
		GetChoices,
		&questionId,
	)
	defer rows.Close()
	if err!=nil{
		r.log.Errorf(err.Error())
		return []models.Choices{},err
	}
	for rows.Next(){
		choice:=models.Choices{}
		err:=rows.Scan(
			&choice.Id,
			&choice.ChoiceText,
		)
		if err!=nil{
			r.log.Errorf(err.Error())
			return []models.Choices{},err
		}
		Choices=append(Choices,choice)
	}
	return Choices,nil
}

func (r *questionRepo) GetAnswer(ctx context.Context, questionId string) (string,error) {

	var answer string

	err:=r.db.QueryRowContext(
		ctx,
		GetAnswer,
		&questionId,
	).Scan(
		&answer,
	)
	if err!=nil{
		r.log.Errorf(err.Error())
		return "",err
	}
	return answer,nil
}