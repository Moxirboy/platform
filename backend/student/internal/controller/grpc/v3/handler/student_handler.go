package handler

import (
	"log"
)

type student struct {
	pb.UnimplementedUserServiceServer
	uc     usecase.IStudentUsecase
	logger logger.Logger
}

func NewStudentHandler(
	usecase usecase.IStudentUsecase,
	log logger.Logger,
) *student {
	return &student{
		uc:     usecase,
		logger: log,
	}
}
