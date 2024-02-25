package server

import (
	"google.golang.org/grpc"
	"log"
	configs "teacher/internal/config"
	"teacher/internal/service/usecase"
	logger "teacher/pkg/log"
	"teacher/pkg/postgres"
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
	// listen, err := net.Listen("tcp", ":5006")
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

	// pb.RegisterUserServiceServer(server, handler.NewTeacherHandler(uc.ITeacherUseCase(), s.log))

	names := server.GetServiceInfo()
	log.Println(names)

	err = server.Serve(listen)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
