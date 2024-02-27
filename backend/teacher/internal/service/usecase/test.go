package usecase

import (
	"context"
	"teacher/internal/models"
)

func (u *TeacherUseCase) CreateTest(
	ctx context.Context,
	Id string,
	model []*models.Test,
) error {
	err:=u.testRepo.CreateTest(ctx, Id, model)
	if err != nil {
		return err
	}
	return nil
}