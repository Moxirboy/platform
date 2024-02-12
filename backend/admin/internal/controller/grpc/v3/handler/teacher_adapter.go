package handler

import (
	"admin/internal/models"
	pb "admin/proto"
)

func FromReqToModel(
	req *pb.User,
) models.User{
	return models.User{
		ID: req.GetID(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Role: req.GetRole(),
		CreatedAt: req.GetCreatedAt(),
		UpdateAt: req.GetUpdateAt(),
		DeletedAt: req.GetDeletedAt(),
		Verified: req.GetVerified(),
	}
}
func FromModelToResponse (
	err error,
) *pb.ErrorResponse{
	return &pb.ErrorResponse{
		ErrorMessage: err.Error(),
	}
}