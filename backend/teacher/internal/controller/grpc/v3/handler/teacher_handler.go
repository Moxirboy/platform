package handler

import (
	"context"
	"log"
	"teacher/internal/service/usecase"
	logger "teacher/pkg/log"
	pb "teacher/proto"
)

type Teacher struct {
	pb.UnimplementedUserServiceServer
	uc     usecase.ITeacherUseCase
	logger logger.Logger
}

func NewTeacherHandler(useCase usecase.ITeacherUseCase, log logger.Logger) *Teacher {
	return &Teacher{
		uc:     useCase,
		logger: log,
	}

}

func (t *Teacher) AddTest(ctx context.Context, request *pb.AddTestRequest) (*pb.Response, error) {
	addTest := FromRequestToModel(request)

	testID, err := t.uc.AddTest(ctx, addTest)
	if err != nil {
		t.logger.Errorf("error is while adding testID in handler err: %v", err.Error())
		return FromModelToResponse(""), err
	}

	log.Println(addTest)

	return &pb.Response{
		Id: testID,
	}, nil
}

func (t *Teacher) StartTest(ctx context.Context, class *pb.CreateClass) (*pb.Response, error) {
	createClass := FromReqToModel(class)

	questionID, err := t.uc.StartTest(ctx, createClass)
	if err != nil {
		t.logger.Errorf("error is while starting test in handler err: %v", err.Error())
		return FromModelToResponse(""), err
	}
	return &pb.Response{
		Id: questionID,
	}, nil
}
