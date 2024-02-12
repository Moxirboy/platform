package handler

import (
	"gateway/internal/controller/http/v1/dto"
	logger "gateway/pkg/log"
	pb "gateway/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type teacher struct {
	g *gin.RouterGroup
	client pb.CreateTeacherClient
	log logger.Logger
}

func NewTeacherHandler(
	g *gin.RouterGroup,
	client pb.CreateTeacherClient,
	log logger.Logger,
){
	teacher:=teacher{
		g,
		client,
		log,
	}
	handler:=g.Group("/admin")
	handler.POST("/create",teacher.CreateTeacher)
}

func (c *teacher) CreateTeacher(ctx *gin.Context) {
	log.Println("post")
	req:=dto.Teacher{}
	if err:=ctx.ShouldBind(&req);err!=nil{
		c.log.Error(err)
	 ctx.JSON(404,err)
	 return
	}
	teacher:=FromReqToTeacher(req)
	m,err:=c.client.CreateUser(ctx, &teacher)
	if err!=nil{
		c.log.Error(err)

		ctx.JSON(404,err)
		return
	   }
	if m.ErrorMessage!=""{
		c.log.DPanic(m.ErrorMessage)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id":      teacher.ID,
		"teacher": teacher.Username,
		// Add other relevant fields
	})
	
}