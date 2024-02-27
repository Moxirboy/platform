package handler

import (
	"gateway/internal/controller/http/v1/dto"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type teachers struct {
	g      *gin.RouterGroup
	client *client.GRPCClient
	log    logger.Logger
}

func NewTeachersHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
){
	handler := &teachers{
		g:      g,
		client: client,
		log:    log,
	}
	teachers:=g.Group("/teacher")
	teachers.POST("/class",handler.CreateClass)
	teachers.POST("/createExam",handler.CreateExam)
}

func (c *teachers) CreateClass(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	req := dto.ClassRequest{}
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.log.Info(req)
	id:=s.Get("id").(string)
	log.Println(id)
	res,err:=c.client.TeacherService().ClassCreate(ctx.Request.Context(), FromClassRequestToProto(id,req))
	if err != nil {
		c.log.Error(err)
		ctx.JSON(404, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "class created",
		"classId": res.Id,
	})
}

func (c *teachers) CreateTest(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	req := []dto.CreateTestRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	id:=s.Get("id").(string)
	_,err:=c.client.TeacherService().CreateTest(ctx.Request.Context(), FromRequestTestsToModel(id,req))
	if err != nil {
		c.log.Error(err)
		ctx.JSON(404, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "test created",
	})
}

func (c *teachers) CreateExam(ctx *gin.Context) {
	
	s:=sessions.Default(ctx)
	req := dto.Class{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("error binding request", err)
		ctx.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	id:=s.Get("id").(string)
	res,err:=c.client.TeacherService().CreateExam(ctx.Request.Context(), FromExamCreateToProto(id,req.Name))
	if err != nil {
		c.log.Error(err)
		ctx.JSON(404, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "exam created",
		"examId": res.Id,
	})
}