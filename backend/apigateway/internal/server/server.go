package server

import (
	configs "gateway/internal/config"

	ht "gateway/internal/controller/http"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"
	"log"

	"github.com/gin-gonic/gin"
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
	log.Println(s.cfg.Admin_RPC.Hosts,s.cfg.Admin_RPC.Port)
	grpc_client,err:=client.New(*s.cfg)
	if err!=nil{
		log.Println("here")
		return err
	}
	
	ht.SetUp(r,*s.cfg,grpc_client,s.log)

	return r.Run(":5005")
	
}