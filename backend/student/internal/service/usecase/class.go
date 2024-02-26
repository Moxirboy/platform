package usecase


import (
	"admin/internal/models"
	"admin/internal/service/repo"
	logger "admin/pkg/log"
	"context"
)

type StudentUsecase struct {
	repo repo.IStudentRepository
	log  logger.Logger
}

func NewStudentUseCase(repo repo.IStudentRepository, log logger.Logger) IStudentUsecase {
	return &StudentUsecase{
		repo: repo,
		log:  log,
	}
}

func (s *StudentUsecase) TakeExam(ctx context.Context, answers models.CreateAnswers) (models.TestResult, error) {
	testresult, err := s.repo.TakeExam(ctx, answers)
	if err != nil {
		s.log.Errorf(err.Error())
		return err
	}
	return nil

}

func (s *StudentUsecase) GetTestResults(ctx context.Context, examID string) ([]models.TestResult, error) {
	testresult, err := s.repo.GetTestResults(examID)
	if err != nil {
		s.log.Errorf(err.Error())
		return nil, err
	}
	return testresult, nil

}
