package usecase

import (
	logger "teacher/pkg/log"
	"context"
	"teacher/internal/models"
	"teacher/internal/service/repo"
)

type TeacherUseCase struct {
	repo repo.ITeacherRepository
	testRepo repo.TestRepository
	log  logger.Logger
}

func NewTeacherUseCase(repo repo.ITeacherRepository, log logger.Logger) ITeacherUseCase {
	return &TeacherUseCase{
		repo: repo,
		log:  log,
	}
}
func (t *TeacherUseCase) CreateClass(ctx context.Context, class models.Class) (string, error) {
	classID, err := t.repo.CreateClass(ctx, class)
	if err != nil {
		t.log.Error("error is while creating class", err.Error())
		return "", err
	}
	return classID, nil
}


func (t *TeacherUseCase) AddTest(ctx context.Context, class models.AddTestRequest) (string, error) {
	questionID, err := t.repo.AddTest(ctx, class)
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

