package server

import (
	"admin/internal/configs"
	"admin/internal/controller/grpc/v3/handler"
	"log"
	// "context"
	"admin/internal/service/usecase"
	logger "admin/pkg/log"
	"admin/pkg/postgres"
	pb "admin/proto"
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
	lis,err:=net.Listen("tcp",":5006")
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
	pb.RegisterCreateTeacherServer(server,handler.NewTeacherHandler(uc.TeacherUsecase(),s.log))
	
	// pb.RegisterCreateTeacherServer(server,&Serve{})
	err=server.Serve(lis)
	if err!=nil{
		log.Println(err)

return err
	}

	return nil
}
// type Serve struct{
// 	pb.UnimplementedCreateTeacherServer
// }
// func (h *Serve) CreateUser(ctx context.Context, req *pb.User) (*pb.ErrorResponse,error){
// 	user:=handler.FromReqToModel(req)
// 	log.Println(user)
// 	return &pb.ErrorResponse{
// 		ErrorMessage: "so bolasla",
// 	},nil
// }