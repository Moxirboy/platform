package handler

import (
	"gateway/internal/controller/http/v1/dto"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"
	"gateway/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type teacher struct {
	
	g *gin.RouterGroup
	client *client.GRPCClient
	log logger.Logger
}

func NewTeacherHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
){
	teacher:=teacher{
		g,
		client,
		log,
	}
	handler:=g.Group("/admin")
	handler.POST("/create",teacher.CreateTeacher)
	handler.GET("/getall",teacher.GetAll)
}

func (c *teacher) CreateTeacher(ctx *gin.Context) {
	req:=dto.Teacher{}
	if err:=ctx.ShouldBind(&req);err!=nil{
		c.log.Error(err)
	 ctx.JSON(404,err)
	 return
	}
	teacher:=FromReqToTeacher(req)
	m,err:=c.client.UserService().CreateUser(ctx, teacher)
	if err!=nil{
		c.log.Error(err)

		ctx.JSON(404,err)
		return
	   }
	   teacher.ID=m.ErrorMessage
	ctx.JSON(http.StatusCreated, gin.H{
		"id":      teacher.ID,
		"teacher": teacher.Username,
		// Add other relevant fields
	})
	
}

func (c *teacher) GetAll(ctx *gin.Context) {
	log.Println("smt")
	query,err:=utils.GetQueryFromCtx(ctx)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	protoQuery:=FromResQueryToQuery(*query)

	protoRes,err:=c.client.UserService().GetUserList(ctx,protoQuery)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	res:=FromResToProtoRes(protoRes)

	ctx.JSON(200,res)
}