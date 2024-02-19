package repo
import "context"
import "auth/internal/domain"

type AuthRepository interface{
	GetUser(
		ctx context.Context,
		username string,
	) (
		string ,
		string,
		string,
		error,
		)
	GetExist(
			ctx context.Context,
			username string,
			) (
			bool,
			error,
			)
	CreateUser(
			ctx context.Context,
			auth *domain.Auth,
			)(
			id string,
			error error,
			)
}