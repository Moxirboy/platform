package repo
import  (
	"context"
	"admin/internal/models"
	"admin/pkg/utils"
)

type ITeacherRepository interface {
	Create(ctx context.Context, teacher *models.User) error
	GetAll(ctx context.Context,PaginationQuery utils.PaginationQuery) (*models.UserList,error)
}