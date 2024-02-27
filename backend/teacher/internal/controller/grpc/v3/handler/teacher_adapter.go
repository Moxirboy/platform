package handler

import (
	"teacher/internal/models"
	tb "teacher/proto/teacher"
)


func FromRequestClassToModel(
	req *tb.Class,
) *models.Class {
	return &models.Class{
		TeacherID: req.TeacherID,
		Name:      req.Name,
		Password:  req.Password,
	}
}
func FromModelToResponseClass(id string) *tb.Response {
	return &tb.Response{
		Id: id,
	}
}

func FromRequestExamToModel(
	
	req *tb.Exam,
) *models.Exam {
	return &models.Exam{
		TeacherID: req.TeacherID,
		Name:   req.Name,

	}
}

func FromRequestTestsToModel(
	req *tb.Tests,
) []*models.Test {
	tests:=[]*models.Test{}
	for _,v:=range req.Tests{
		tests=append(tests,FromRequestTestToModel(v))
	}
	return tests
}


func FromRequestTestToModel(
	req *tb.Test,
) *models.Test {
	choice:=[]*models.Choice{}
	for _,v:=range req.Choices{
		choice=append(choice,FromRequestChoicesToModel(v))
	}
	return &models.Test{
		QuestionText: req.QuestionText,
		Choices:      choice,
		Answer:       req.Answer,
	}
}

func FromRequestChoicesToModel(
	req *tb.Choice,
) *models.Choice {
	return &models.Choice{
		ChoiceText: req.ChoiceText,
		IsAnswer:   req.IsAnswer,
	}
}