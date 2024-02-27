package server

import (
	"log"
	"net"
	configs "student/internal/config"
	"student/internal/controller/grpc/v3/handler"
	"student/internal/service/usecase"
	pb "student/proto/exam"
	logger "student/pkg/log"
	"student/pkg/postgres"
	"google.golang.org/grpc"
)

type Server struct {
	cfg *configs.Config
	log logger.Logger
}

func NewServer(cfg *configs.Config, log logger.Logger) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

func (s Server) Run() error {
	listen, err := net.Listen("tcp", ":5007")
	log.Println(listen)
	if err != nil {
		s.log.Fatalf("failed to listen: %v", err)
	}
	pg, err := postgres.DB(&s.cfg.Postgres)
	if err != nil {
		s.log.Fatal(err)
	}

	uc := usecase.New(pg, s.log)

	server := grpc.NewServer()
	pb.RegisterTestServiceServer(server,handler.NewStudentHandler(uc.QuestionUsecase(),s.log))

	names := server.GetServiceInfo()
	log.Println(names)

	err = server.Serve(listen)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
