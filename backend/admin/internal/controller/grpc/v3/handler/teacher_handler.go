package handler

import (
	"admin/internal/service/usecase"
	logger "admin/pkg/log"
	pb "admin/proto"
	"context"
	"log"
)


type teacher struct {
	pb.UnimplementedUserServiceServer
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
	
	err:=h.uc.Create(ctx,&user)
	if err!=nil{
		h.logger.Errorf(err.Error())
		return FromModelToResponse(err),err
	}
	log.Println(user)
	return &pb.ErrorResponse{
		ErrorMessage: user.ID,
	},nil
	
}
func (h *teacher) GetUserList(ctx context.Context, in *pb.PaginationQuery, ) (*pb.UserList, error) {
	query:=FromReqQueryToModel(in)
	list,err:=h.uc.GetAll(ctx,query)
	if err!=nil{
		h.logger.Errorf(err.Error())
		return nil,err
	}

	res:=FromModelToResponseQuery(list)

	return res,nil
}
