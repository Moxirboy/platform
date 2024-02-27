package handler

import (
	"gateway/internal/controller/http/v1/dto"
	sp "gateway/proto/exam"
)
func FromIdToExamProto(
	id string,
) *sp.IdRequest  {
	return &sp.IdRequest{Id:id}
}

func FromTestTakeToProto(
	userId string,
	examId string,
) *sp.TestTake  {
	return &sp.TestTake{ExamId:examId}
}

func FromProtoToResponse(
	res *sp.TestResponse,
) []*dto.Question {
	questions:=[]*dto.Question{}
	for _,question:=range res.Test{
		questions=append(questions,FromProtoResponseToTest(question))
	}
	return questions
}



func FromProtoResponseToTest(
	res *sp.Test,
) *dto.Question {
	choices:=[]*dto.Choices{}
	for _,choice:=range res.Choices{
		choices=append(choices,FromChoiceProtoToDto(choice))
	}
	return &dto.Question{
		ID:                res.Id,
		Text:              res.Question,
		Choices:           choices,
		CorrectChoiceIndex: res.Answer,
}}

func FromChoiceProtoToDto(
	res *sp.Choices,
) *dto.Choices {
	return &dto.Choices{
		Id:  res.Id,
		ChoiceText: res.Choices,
		IsCorrect: res.IsCorrect,
}}