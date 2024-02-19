package handler

import "gateway/internal/controller/http/v1/dto"
import "gateway/protos/authpb"

func FromRequestToPb(
	req dto.LoginRequest,
) *authpb.LoginRequest {
return &authpb.LoginRequest{
	Username: req.Username,
	Password: req.Password,
}
}