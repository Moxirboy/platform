package server

import (
	 "auth/internal/configs"
	"auth/internal/controller/handler"
	"log"
	// "context"
	"auth/internal/service/usecase"
	logger "admin/pkg/log"
	"auth/pkg/postgres"
	"auth/proto/authpb"
	"net"

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
func (s Server) Run() error {
	lis,err:=net.Listen("tcp",":5007")
	log.Println(lis)
	if err != nil {
		s.log.Fatalf("failed to listen: %v", err)
	}
	pg,err:=postgres.DB(
		&s.cfg.Postgres,
	)
	if err!=nil{
		s.log.Fatal(
			err,
		)
	} 
	uc:=usecase.New(pg,s.log)


	server:=grpc.NewServer()
	
	authpb.RegisterAuthServiceServer(server,handler.NewAuthHandler(uc.AuthUseCase(),s.log))
	
	names:=server.GetServiceInfo()
	log.Println(names)
	// pb.RegisterCreateTeacherServer(server,&Serve{})
	err=server.Serve(lis)
	if err!=nil{
		log.Println(err)
  		return err
	}

	return nil
}
