package main

import (
	"admin/internal/configs"
	"admin/internal/server"
	logger "admin/pkg/log"
)

func main(){

	var (
		config=configs.Load()
	)
	log:=logger.NewLogger(config.Logger.Level,config.Logger.Encoding)
	log.InitLogger()

	log.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)
	s:=server.NewServer(config,log)
	log.Fatal(s.Run())
}
// type serve struct{
// 	pb.UnimplementedCreateTeacherServer
// }
// func (h *serve) CreateUser(ctx context.Context, req *pb.User) (*pb.ErrorResponse,error){
// return &pb.ErrorResponse{
// 	ErrorMessage: "hello",
// },nil}

// func main(){

// 	lis,err:=net.Listen("tcp","0.0.0.0:50051")
// 	server:=grpc.NewServer()
// 	pb.RegisterCreateTeacherServer(server,&serve{})
// 	names:=server.GetServiceInfo()
// 	log.Println(names)
// 	// pb.RegisterCreateTeacherServer(server,&Serve{})
// 	err=server.Serve(lis)
// 	if err!=nil{
// 		log.Println(err)
// 	}
// }