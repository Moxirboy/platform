package handler

import (
	"admin/internal/service/usecase"
	"admin/proto/authpb"
	"context"
	logger "admin/pkg/log"
)

type authHandler struct {
	authpb.UnimplementedAuthServiceServer
	uc  usecase.AuthUsecase
	log logger.Logger
}

func NewAuthHandler(
	uc usecase.AuthUsecase,
	log logger.Logger,
) *authHandler {
	return &authHandler{
		uc:  uc,
		log: log,
	}
}

func (h *authHandler) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	exist, id, role, error := h.uc.GetUser(ctx, in.Firstname,in.Lastname, in.Password)
	if !exist {
		return FromModelToPb(exist, "", ""), error
	}
	return FromModelToPb(exist, id, role), nil
}

func (h *authHandler) SignUp(ctx context.Context, in *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	req := FromPbToModel(in)
	h.log.Info(req)
	id, role, err := h.uc.CreateUser(ctx, &req)
	if err != nil {
		return nil, err
	}
	
	res := FromModelToPbSignUp(id, role, req.Class)

	return res, nil
}
