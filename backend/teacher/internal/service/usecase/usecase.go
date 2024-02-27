package usecase

import (
	"context"
	"teacher/internal/models"
)

type ITeacherUseCase interface {
	CreateClass(ctx context.Context, class models.Class) (string, error)
	AddTest(context.Context, models.AddTestRequest) (string, error)
	StartTest(context.Context, models.CreateClass) (string, error)
	CreateTest(
		ctx context.Context,
		Id string,
		model []*models.Test,
	) error 

}

type IExamUseCase interface {
	CreateExam(ctx context.Context, exam *models.Exam) (string, error)
}