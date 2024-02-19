package usecase
import "context"
import "auth/internal/domain"
type AuthUsecase interface{
	GetUser(
		ctx context.Context,
		username string,
		password string,
	) (
		bool,
		string,
		string,
		error,
		)
	CreateUser(
		ctx context.Context,
		auth *domain.Auth,
	)(
		id string,
		err error,
	)
	
}