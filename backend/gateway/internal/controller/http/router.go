package http

import (
	configs "gateway/internal/config"
	"gateway/internal/controller/http/v1/handler"
	logger "gateway/pkg/log"
	pb "gateway/proto"

	"github.com/gin-gonic/gin"
)



func SetUp(
	g *gin.Engine,
	cfg configs.Config,
	pb pb.CreateTeacherClient,
	log logger.Logger,
) {
	SetUpHandlerV1(
		g.Group("/v1"),
		cfg,
		pb,
		log,
	)
	
}

func SetUpHandlerV1(
	Group *gin.RouterGroup,
	cfg configs.Config,
	pb pb.CreateTeacherClient,
	log logger.Logger,
) {

	handler.NewTeacherHandler(
		Group,
		pb,
		log,
	)
}