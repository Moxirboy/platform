package handler

import (
	"auth/internal/service/usecase"
	logger "gateway/pkg/log"
	"context"
	"auth/proto/authpb"
)


type authHandler struct{
	authpb.UnimplementedAuthServiceServer
	uc usecase.AuthUsecase
	log logger.Logger
}

func NewAuthHandler(
	uc usecase.AuthUsecase,
	log logger.Logger,
) *authHandler {
	return &authHandler{
		uc: uc,
		log: log,
	}
}

func (h *authHandler) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error){
	exist,id,role,error:=h.uc.GetUser(ctx,in.Username,in.Password)
	if !exist{
		return FromModelToPb(exist,"",""),error
	}
	return FromModelToPb(exist,id,role),nil
}

func (h *authHandler) SignUp(ctx context.Context,)