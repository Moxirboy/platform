package handler

import (
	"teacher/internal/models"
	pb "teacher/proto"
)

func FromRequestToModel(request *pb.AddTestRequest) models.AddTestRequest {
	return models.AddTestRequest{
		QuestionID:   request.QuestionId,
		QuestionText: request.QuestionText,
		TeacherID:    request.TeacherID,
		AnswerID:     request.AnswerID,
		ChoiceID:     request.ChoiceID,
	}
}

func FromModelToResponse(id string) *pb.Response {
	return &pb.Response{
		Id: id,
	}
}

func FromReqToModel(request *pb.CreateClass) models.CreateClass {
	return models.CreateClass{
		ID:        request.ID,
		TeacherID: request.TeacherID,
		ClassID:   request.ClassID,
		Period:    request.Period,
	}
}
