package usecase

import (
	"context"
	"student/internal/models"
	"student/internal/service/repo"
	"student/pkg/log"
)



type questionUsecase struct {
	repo repo.QuestionRepo
	Examrepo repo.ExamRepository
	log  logger.Logger
}

func NewQuestionUseCase(repo repo.QuestionRepo, Examrepo repo.ExamRepository,log logger.Logger) IQuestionUsecase {
	return &questionUsecase{
		repo: repo,
		Examrepo: Examrepo,
		log:  log,
	}
}

func (u *questionUsecase) GetAllQuestions(ctx context.Context, examId string) ([]models.QuestionsAll, error) {
	questions, err := u.repo.GetQuestions(ctx, examId)
	if err != nil {
		u.log.Errorf(err.Error())
		return nil, err
	}
	for i, question := range questions {
		choices, err := u.repo.GetChoices(ctx, question.Id)
		if err != nil {
			u.log.Errorf(err.Error())
			return nil, err
		}
		questions[i].Choices = choices
		
	}

	return questions, nil
}