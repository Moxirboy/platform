package handler

import (
	domain "admin/internal/models"
	"admin/proto/authpb"
)

func FromModelToPb(
	exist bool,
	id string,
	role string,
) *authpb.LoginResponse {
	return &authpb.LoginResponse{
		Exist: exist,
		Id:    id,
		Role:  role,
	}
}
func FromPbToModel(
	pb *authpb.SignUpRequest,
) domain.StudentAuth {
	return domain.StudentAuth{
		Firstname: pb.Firstname,
		Lastname:  pb.Lastname,
		Password:  pb.Password,
		Class:     pb.Class,
	}
}
func FromModelToPbSignUp(
	id string,
	role string,
	class string,
) *authpb.SignUpResponse {
	return &authpb.SignUpResponse{
		Id:      id,
		Role:    role,
		ClassId: class,
	}
}
