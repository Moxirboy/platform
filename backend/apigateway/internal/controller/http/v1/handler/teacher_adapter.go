package handler

import "gateway/internal/controller/http/v1/dto"
import tp "gateway/proto/teacher"

func FromClassRequestToProto(
	id string,
	req dto.ClassRequest,
) *tp.Class  {
	return &tp.Class{
		Name: req.Name,
		Password: req.Password,
		TeacherID: id,
	}
}

func FromRequestTestsToModel(
	id string,
	req []dto.CreateTestRequest,
) *tp.Tests  {
	tests:=[]*tp.Test{}
	for _,test:=range req{
		tests=append(tests,FromRequestTestToModel(test))
	}
	return &tp.Tests{
		TeacherID: id,
		Tests: tests,
	}
}

func FromRequestTestToModel(
	req dto.CreateTestRequest,
) *tp.Test  {
	choices:=[]*tp.Choice{}
	for _,choice:=range req.Choices{
		choices=append(choices,FromRequestChoiceToModel(choice))
	}
	return &tp.Test{
		QuestionText: req.QuestionText,
		Choices: choices,
		Answer: req.Answer,
	}
}


func FromRequestChoiceToModel(
	req dto.Choice,
) *tp.Choice  {
	return &tp.Choice{
		ChoiceText: req.ChoiceText,
		IsAnswer: req.IsAnswer,
	}
}

func FromExamCreateToProto(
	id string,
	name string,
) *tp.Exam  {
	return &tp.Exam{
		Name: name,
		TeacherID: id,
	}
}