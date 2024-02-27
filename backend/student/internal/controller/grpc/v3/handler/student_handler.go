package handler

import (
	"student/pkg/log"
	pb "student/proto/exam"
	"student/internal/service/usecase"
	"context"
)

type student struct {
	pb.UnimplementedTestServiceServer
	uc     usecase.IQuestionUsecase
	logger logger.Logger
}

func NewStudentHandler(
	usecase usecase.IQuestionUsecase,
	log logger.Logger,
) *student {
	return &student{
		uc:     usecase,
		logger: log,
	}
}

func (h *student) TakeExamQuestions(ctx context.Context, in *pb.TestTake) (*pb.TestResponse, error) {
	test,err:=h.uc.GetAllQuestions(ctx,in.ExamId)
	if err!=nil{
		return nil,err
	}
	return ConvertToTestResponse(test),nil
}
func(h *student) GetExamExist(ctx context.Context, in *pb.IdRequest) (*pb.IdResponse, error){
	exist ,id,err:=h.uc.GetExamExist(ctx,in.Id)
	if err!=nil{
		return nil,err
	}


	return &pb.IdResponse{Exist:exist,Id:id},nil
}
