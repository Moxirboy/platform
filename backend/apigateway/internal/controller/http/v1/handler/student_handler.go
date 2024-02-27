package handler

import (
	"gateway/internal/controller/http/v1/dto"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type student struct {
	g      *gin.RouterGroup
	client *client.GRPCClient
	log    logger.Logger
}

func NewStudentHandler(
	g *gin.RouterGroup,
	client *client.GRPCClient,
	log logger.Logger,
){
	handler := &student{
		g:      g,
		client: client,
		log:    log,
	}
	students:=g.Group("/student")
	students.GET("/exam",handler.GetExistExam)
	students.GET("/questions",handler.GetQuestions)

}
func (c *student) GetExistExam(ctx *gin.Context) {

	s:=sessions.Default(ctx)
	UserId:=s.Get("id").(string)
	c.log.Info(UserId)
	req:=FromIdToExamProto(UserId)
	c.log.Info(req)

	res,err:=c.client.StudentService().GetExamExist(ctx.Request.Context(), req)
	if err != nil {
		c.log.Error(err)

		ctx.JSON(404, err)
		return
	}
	s.Set("examId",res.Id)
	s.Save()
	ctx.JSON(http.StatusCreated, gin.H{
		"testAvailable": res.Exist,
	})

}
func (c *student) GetQuestions(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	UserId:=s.Get("id").(string)
	ExamId:=s.Get("examId").(string)
	res,err:=c.client.StudentService().TakeExamQuestions(ctx.Request.Context(), FromTestTakeToProto(UserId,ExamId))
	if err != nil {
		c.log.Error(err)
		ctx.JSON(404, err)
		return
	}
	ctx.JSON(http.StatusCreated, FromProtoToResponse(res))
}

func(c *student) SubmitExam(ctx *gin.Context){
	// s:=sessions.Default(ctx)
	// UserId:=s.Get("id").(string)
	// ExamId:=s.Get("examId").(string)
	req := []*dto.Answer{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var score int
	for _, ans := range req {
		if ans.AnswerIndex == ans.Answer {
			score++
		}
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"score": score,
	})
}