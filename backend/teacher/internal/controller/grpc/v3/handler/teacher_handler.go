package handler

import (
	"context"
	
	"teacher/internal/service/usecase"
	logger "teacher/pkg/log"
	tp "teacher/proto/teacher"
)

type Teacher struct {
	tp.UnimplementedTeacherServiceServer
	uc     usecase.ITeacherUseCase
	u 	usecase.IExamUseCase
	logger logger.Logger
}


func NewTeacherHandler(useCase usecase.ITeacherUseCase,u 	usecase.IExamUseCase, log logger.Logger) *Teacher {
	return &Teacher{
		uc:     useCase,
		u: u,
		logger: log,
	}

}

func (t *Teacher)CreateTest(ctx context.Context, in *tp.Tests) (*tp.Res, error){
	id:=in.TeacherID
	req:=FromRequestTestsToModel(in)
	err:=t.uc.CreateTest(ctx,id,req)
	if err!=nil{
		t.logger.Errorf("error is while creating test in handler err: %v", err.Error())
		return nil,err
	}
	return nil,nil
}

// create class with name and password


func (t *Teacher) ClassCreate(ctx context.Context,in *tp.Class) (*tp.Response, error){
	t.logger.Info(in)
	req := FromRequestClassToModel(in)
	classID, err := t.uc.CreateClass(ctx, *req)
	if err != nil {
		t.logger.Errorf("error is while creating class in handler err: %v", err.Error())
		return FromModelToResponseClass(""), err
	}
	return FromModelToResponseClass(classID), nil
}

func(t *Teacher) CreateExam(ctx context.Context,in *tp.Exam) (*tp.Response, error){
	req:=FromRequestExamToModel(in)

	examID,err:=t.u.CreateExam(ctx,req)
	if err!=nil{
		t.logger.Errorf("error is while creating exam in handler err: %v", err.Error())
		return FromModelToResponseClass(""),err
	}
	return FromModelToResponseClass(examID),nil
}





