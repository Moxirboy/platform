package server

import (
	configs "gateway/internal/config"
	"gateway/internal/controller/http"
	logger "gateway/pkg/log"
	pb "gateway/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


type Server struct {
	cfg *configs.Config
	log logger.Logger
}

func NewServer(
	cfg *configs.Config,
	log logger.Logger,
) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

func (s *Server) Run() error {
	r:=gin.New()
	conn,err:=grpc.Dial("localhost:5006",grpc.WithInsecure())
	if err!=nil{
		return err
	}
	defer conn.Close()
	c:=pb.NewCreateTeacherClient(conn)
	http.SetUp(r,*s.cfg,c,s.log)

	return r.Run(":5005")
}