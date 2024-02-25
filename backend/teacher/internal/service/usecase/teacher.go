package usecase

import (
	logger "admin/pkg/log"
	"golang.org/x/net/context"
	"teacher/internal/models"
	"teacher/internal/service/repo"
)

type TeacherUseCase struct {
	repo repo.ITeacherRepository
	log  logger.Logger
}

func NewTeacherUseCase(repo repo.ITeacherRepository, log logger.Logger) ITeacherUseCase {
	return &TeacherUseCase{
		repo: repo,
		log:  log,
	}
}
func (t *TeacherUseCase) AddTest(ctx context.Context, request models.AddTestRequest) (string, error) {
	questionID, err := t.repo.AddTest(ctx, request)
	if err != nil {
		t.log.Error("error is while add test", err.Error())
		return "", err
	}
	return questionID, nil
}

func (t *TeacherUseCase) StartTest(ctx context.Context, class models.CreateClass) (string, error) {
	classID, err := t.repo.StartTest(ctx, class)
	if err != nil {
		t.log.Error("error is while starting test", err.Error())
		return "", err
	}
	return classID, nil
}
