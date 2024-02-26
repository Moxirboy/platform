package handler

import "gateway/internal/controller/http/v1/dto"
import "gateway/proto/authpb"

func FromRequestToPb(
	req dto.LoginRequest,
) *authpb.LoginRequest {
	return &authpb.LoginRequest{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Password: req.Password,
	}
}
func FromRequestToPbSignUp(
	req dto.SignUpRequest,
) *authpb.SignUpRequest {
	return &authpb.SignUpRequest{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  req.Password,
		Class:     req.Class,
	}
}
func FromPbToResponse(
	pb *authpb.SignUpResponse,
) dto.SignUpResponse {
	return dto.SignUpResponse{
		Id:      pb.Id,
		Role:    pb.Role,
		ClassId: pb.ClassId,
	}

}
