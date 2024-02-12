package handler

import (
	"admin/internal/service/usecase"
	logger "admin/pkg/log"
	pb "admin/proto"
	"context"
	"log"
)


type teacher struct {
	pb.UnimplementedCreateTeacherServer
	uc usecase.ITeacherUsecase
	logger logger.Logger
}

func NewTeacherHandler(
	usecase usecase.ITeacherUsecase,
	log logger.Logger,
) *teacher {
	return &teacher{
		uc: usecase,
		logger: log,
		
	}

}

func (h *teacher) CreateUser(ctx context.Context, req *pb.User) (*pb.ErrorResponse,error){
	user:=FromReqToModel(req)
	log.Println(user)
	err:=h.uc.Create(ctx,&user)
	if err!=nil{
		h.logger.Errorf(err.Error())
		return FromModelToResponse(err),err
	}
	return &pb.ErrorResponse{
		ErrorMessage: "so bolasla",
	},nil
}
