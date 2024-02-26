package repo

import "context"
import "admin/internal/models"

type AuthRepository interface {
	GetUser(
		ctx context.Context,
		firstname string,
		lastname string,
			) (
		string,
		string,
		string,
		error,
	)
	GetExist(
		ctx context.Context,
		firstname string,
		lastname string,
	) (
		bool,
		error,
	)
	CreateUser(
		ctx context.Context,
		auth *models.StudentAuth,
	) (
		id string,
		role string,
		err error,
	)
}
