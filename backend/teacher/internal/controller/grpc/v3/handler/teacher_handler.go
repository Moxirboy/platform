package handler

import (
	"teacher/internal/service/usecase"
	logger "teacher/pkg/log"
)

type Teacher struct {
	// pb.UnimplementedUserServiceServer
	uc     usecase.ITeacherUseCase
	logger logger.Logger
}

func NewTeacherHandler(useCase usecase.ITeacherUseCase, log logger.Logger) *Teacher {
	return &Teacher{
		uc:     useCase,
		logger: log,
	}

}

// func (h *Teacher) AddTest(ctx context.Context, req *pb.User) (*pb.ErrorResponse, error) {
//user := FromReqToModel(req)
//
//err := h.uc.Create(ctx, &user)
//if err != nil {
//	h.logger.Errorf(err.Error())
//	return FromModelToResponse(err), err
//}
//log.Println(user)
//return &pb.ErrorResponse{
//	ErrorMessage: user.ID,
//}, nil
//}
