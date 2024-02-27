package usecase

import (
	"teacher/internal/service/repo"
	logger "teacher/pkg/log"
	"context"
	"teacher/internal/models"
)

type examUseCase struct {
	repo repo.ExamRepository
	classRepo repo.ITeacherRepository
	log logger.Logger
}


func NewExamUseCase(repo repo.ExamRepository,classRepo repo.ITeacherRepository,log logger.Logger) IExamUseCase {
	return &examUseCase{
		repo: repo,
		classRepo:classRepo,
		log:  log,
	}
}

func (u *examUseCase) CreateExam(ctx context.Context, exam *models.Exam) (string, error){
	u.log.Info(exam)
	id,err:= u.classRepo.GetClassIdByName(ctx, exam.Name)
	if err!=nil{
		u.log.Error("error is while getting class id", err.Error())
		return "", err
	}
	exam.ClassID=id
	examID, err := u.repo.CreateExam(ctx, *exam)
	if err != nil {
		u.log.Error("error is while creating exam", err.Error())
		return "", err
	}
	return examID, nil
}
