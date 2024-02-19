package main

// import (
// 	configs "gateway/internal/config"
// 	"gateway/internal/server"
// 	logger "gateway/pkg/log"
// )
import (
	
	configs "gateway/internal/config"

	"gateway/internal/server"
	logger "gateway/pkg/log"
	

	
)


func main(){
	var (
		config = configs.Load()
	)
	log := logger.NewLogger(config.Logger.Level, config.Logger.Encoding)
	log.InitLogger()
 
	log.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)
	s:=server.NewServer(config,log)
	
	log.Fatal(s.Run())

	// conn,err:=grpc.Dial("0.0.0.0:5006",grpc.WithInsecure())
	// if err!=nil{
		
	// 	log.Println("here")

	// } 



	// log.Println(conn)
	// defer conn.Close()
	// c:=pb.NewCreateTeacherClient(conn)
	// ctx:=context.Background()
	// res,err:=c.CreateUser(ctx,&pb.User{
	// 	ID: "1",
	// 	Username: "moxirboy",
	// 	Password: "moxirboy",
	// 	Role: "teacher",
	// 	CreatedAt: "2004-04-20",
	// 	UpdateAt: "2004-04-20",
	// 	DeletedAt: "2004-04-20",
	// 	Verified: true,
	// })

	// log.Println(err)
	 

	// fmt.Println(res)
}