package handler

import (
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"

	"github.com/gin-gonic/gin"
)




type dash struct {
	g      *gin.RouterGroup
	client *client.GRPCClient
	log    logger.Logger
}

func NewDashHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
) {
	dash := dash{
		g,
		client,
		log,
	}
	handler := g.Group("/dash")
	handler.GET("/", dash.GetDash)
	
}

func (c *dash) GetDash(ctx *gin.Context) {
	ctx.HTML(200,"dashboard.html",nil)
}