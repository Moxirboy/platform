package repo

import (
	"context"
	"teacher/internal/models"
)

type TestRepository interface {
	CreateTest(
		ctx context.Context,
		Id string,
		model []*models.Test,
	) error 
}