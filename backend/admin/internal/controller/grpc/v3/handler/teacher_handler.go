package handler

import (
	"admin/internal/service/usecase"
	logger "admin/pkg/log"
	pb "admin/proto"
	"context"
)


type teacher struct {
	pb.CreateTeacherServer
	uc usecase.ITeacherUsecase
	logger logger.Logger
}

func NewTeacherHandler(
	usecase usecase.ITeacherUsecase,
	log logger.Logger,
) teacher {
	return teacher{
		uc: usecase,
		logger: log,
	}

}
func (h *teacher) CreateUser(ctx context.Context, req *pb.User) *pb.ErrorResponse{
	user:=fromReqToModel(req)
	err:=h.uc.Create(ctx,&user)
	if err!=nil{
		h.logger.Errorf(err.Error())
		return FromModelToResponse(err)
	}
	return nil
}
