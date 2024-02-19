package http

import (
	configs "gateway/internal/config"
	"gateway/internal/controller/http/v1/handler"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"

	"github.com/gin-gonic/gin"
)



func SetUp(
	g *gin.Engine,
	cfg configs.Config,
	client *client.GRPCClient,
	log logger.Logger,
) {
	SetUpHandlerV1(
		g.Group("/v1"),
		cfg,
		client,
		log,
	)
	
}

func SetUpHandlerV1(
	Group *gin.RouterGroup,
	cfg configs.Config,
	client *client.GRPCClient,
	log logger.Logger,
) {
	handler.NewTeacherHandler(
		Group,
		client,
		log,
	)
}