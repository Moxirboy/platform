package usecase
import (
	"context"
)
func (u *questionUsecase) GetExamExist(ctx context.Context, userId string) (bool,string ,error) {
	exist, err := u.Examrepo.GetExist(ctx, userId)
	u.log.Info(exist)
	if err != nil {	
		return false,"",err
	}
	if !exist{
		return false,"",nil
	}
	u.log.Info(exist)
	examId, err := u.Examrepo.GetExamId(ctx, userId)
	if err != nil {
		return false,"",err
	}
	return true,examId,nil
}