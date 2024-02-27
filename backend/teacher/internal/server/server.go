package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	configs "teacher/internal/config"
	"teacher/internal/controller/grpc/v3/handler"
	"teacher/internal/service/usecase"
	logger "teacher/pkg/log"
	"teacher/pkg/postgres"
	pb "teacher/proto/teacher"
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
	listen, err := net.Listen("tcp", ":5008")
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

	pb.RegisterTeacherServiceServer(server, handler.NewTeacherHandler(uc.ITeacherUseCase(), uc.IExamUseCase(),s.log))

	names := server.GetServiceInfo()
	log.Println(names)

	err = server.Serve(listen)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
