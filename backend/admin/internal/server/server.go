package server

import (
	"admin/internal/configs"
	"admin/internal/controller/grpc/v3/handler"
	"admin/internal/service/usecase"
	logger "admin/pkg/log"
	"admin/pkg/postgres"
	"net"
	pb "admin/proto"
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
	lis,err:=net.Listen("tcp",s.cfg.RPC.Port)
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
	serve:=handler.NewTeacherHandler(uc.TeacherUsecase(),s.log)
	server:=grpc.NewServer()
	pb.RegisterCreateTeacherServer(server,serve.CreateTeacherServer)
	return server.Serve(lis)
}