package handler

import (
	"gateway/internal/controller/http/v1/dto"
	client "gateway/pkg/grpc"
	"gateway/pkg/jwt"
	logger "gateway/pkg/log"
	"gateway/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type auth struct {
	g      *gin.RouterGroup
	client *client.GRPCClient
	log    logger.Logger
}

func NewAuthHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
) {
	handler := auth{
		g,
		client,
		log,
	}
	auths := g.Group("/auth")
	auths.POST("/loginpost", handler.Login)
	auths.POST("/signuppost", handler.SignUp)
	auths.GET("/login", handler.LoginPage)
	auths.GET("/signup", handler.SignPage)
}
func (c *auth) LoginPage(ctx *gin.Context){
	ctx.HTML(200,"login.html",nil)
}

func (c *auth) SignPage(ctx *gin.Context){
	ctx.HTML(200,"sign.html",nil)
}

func (c *auth) Login(ctx *gin.Context) {
	req := dto.LoginRequest{}
	s:=sessions.Default(ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.log.Info("login request", req.Firstname, req.Lastname, req.Password)
	invalidParams := utils.Validate(req)
	if invalidParams != nil {
		c.log.Error("invalid request", invalidParams)
		ctx.JSON(400, dto.ErrorResponse{Message: "invalid request"})
		return
	}
	pb := FromRequestToPb(req)

	pbRes, err := c.client.AuthService().Login(ctx.Request.Context(), pb)
	if err != nil {
		c.log.Error("error calling login", err)
		ctx.JSON(500, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	if !pbRes.Exist {
		ctx.JSON(403, dto.ErrorResponse{Message: "No Such User"})
		return
	}

	token,err :=jwt.CreateToken(pbRes.Id, pbRes.Role)
	
	if err != nil {
		c.log.Error("error creating token", err)
		ctx.JSON(500, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	s.Set("id",pbRes.Id)
	s.Set("role",pbRes.Role)
	s.Save()
	ctx.JSON(200, dto.Response{
		Token: token,
		Role: pbRes.Role,
	})
}

func (c *auth) SignUp(ctx *gin.Context) {
	req := dto.SignUpRequest{}
	s:=sessions.Default(ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: "invalid request"})
		return
	}
	invalidParams := utils.Validate(req)
	if invalidParams != nil {
		c.log.Error("invalid request", invalidParams)
		ctx.JSON(400, dto.ErrorResponse{Message: "invalid request"})
		return
	}
	pb := FromRequestToPbSignUp(req)
	pbRes, err := c.client.AuthService().SignUp(ctx.Request.Context(), pb)
	
	if err != nil {
		c.log.Error("error calling signup", err)
		ctx.JSON(500, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	token,err :=jwt.CreateToken(pbRes.Id, pbRes.Role)
	
	if err != nil {
		c.log.Error("error creating token", err)
		ctx.JSON(500, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	s.Set("id",pbRes.Id)
	s.Set("role",pbRes.Role)
	s.Save()
	ctx.JSON(200, dto.Response{Token: token})
}
