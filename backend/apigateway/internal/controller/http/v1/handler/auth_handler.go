package handler

import (
	"gateway/internal/controller/http/v1/dto"
	logger "gateway/pkg/log"
	"gateway/pkg/utils"
	"gateway/protos/authpb"
	client "gateway/pkg/grpc"

	"github.com/gin-gonic/gin"
)

type auth struct {
	g  *gin.RouterGroup
	client *client.GRPCClient
	log logger.Logger
}

func NewAuthHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
) {
	auth := auth{
		g,
		client,
		log,
	}
}

func (c *auth) Login(ctx *gin.Context){
	req:=dto.LoginRequest{}
	if err:=ctx.ShouldBindJSON(&req); err!=nil{
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: "invalid request"})
		return
	}
	invalidParams:=utils.Validate(req)
	if invalidParams!=nil{
		c.log.Error("invalid request", invalidParams)
		ctx.JSON(400, dto.ErrorResponse{Message: "invalid request"})
		return
	}
	pb:=FromRequestToPb(req)
	pbRes,err:=c.client.AuthService().Login(ctx,pb)
	if err!=nil{
		c.log.Error("error calling login", err)
		ctx.JSON(500, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	if !pbRes.Exist{
		ctx.JSON(403, dto.ErrorResponse{Message: "No Such User"})
		return
	}


}