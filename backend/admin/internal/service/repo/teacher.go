package repo
import  (
	"context"
	"admin/internal/models"
)

type ITeacherRepository interface {
	Create(ctx context.Context, teacher *models.User) error
}