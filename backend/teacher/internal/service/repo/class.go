package repo

import (
	"context"
	"teacher/internal/models"
)

type ITeacherRepository interface {
	CreateClass(ctx context.Context, class models.Class) (string,error)
	GetClassIdByName(ctx context.Context, name string) (string, error)
	AddTest(context.Context, models.AddTestRequest) (string, error)
	StartTest(context.Context, models.CreateClass) (string, error)
}
