package handler

import "auth/proto/authpb"

func FromModelToPb(
	exist bool,
	id string,
	role string,
) *authpb.LoginResponse {
	return &authpb.LoginResponse{
		Exist: exist,
		Id: id,
		Role: role,
	}
}