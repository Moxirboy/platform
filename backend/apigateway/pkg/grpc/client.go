package client

import (
	"fmt"
	configs "gateway/internal/config"
	pb "gateway/proto"
	"gateway/proto/authpb"
	sp "gateway/proto/exam"
	tp "gateway/proto/teacher"
	"log"

	"google.golang.org/grpc"
)

type IGRPCClient interface {
	UserService() pb.UserServiceClient
	AuthService() authpb.AuthServiceClient
	StudentService() sp.TestServiceClient
	TeacherService() tp.TeacherServiceClient
}

type GRPCClient struct {
	cfg         configs.Config
	connections map[string]interface{}
}

const (
	_adminRpc = "admin_rpc"
	_authRpc  = "auth_rpc"
	_studentRpc = "student_rpc"
	_teacherRpc = "teacher_rpc"
)

func New(cfg configs.Config) (*GRPCClient, error) {
	var connections = make(map[string]interface{})
	{
		connAdmin, err := grpc.Dial(fmt.Sprintf(
			"%s:%s", cfg.Admin_RPC.Hosts, cfg.Admin_RPC.Port,
		), grpc.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
				cfg.Admin_RPC.Hosts, cfg.Admin_RPC.Port, err,
			)
		}
		connAuth, err := grpc.Dial(fmt.Sprintf(
			"%s:%s", cfg.Admin_RPC.Hosts, cfg.Admin_RPC.Port,
		), grpc.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
				cfg.Admin_RPC.Hosts, cfg.Admin_RPC.Port, err,
			)
		}
		log.Println(cfg.Student_RPC.Hosts,cfg.Student_RPC.Port)

		// connStudent, err := grpc.Dial(fmt.Sprintf(
		// 	"%s:%s", cfg.Student_RPC.Hosts, cfg.Student_RPC.Port,
		// ), grpc.WithInsecure())
		connStudent, err := grpc.Dial("student:5007",grpc.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",

				cfg.Student_RPC.Hosts, cfg.Student_RPC.Port, err,
			)
		}
		log.Println(cfg.Teacher_RPC.Hosts,cfg.Teacher_RPC.Port)
		connTeacher, err := grpc.Dial(fmt.Sprintf(
			"%s:%s", cfg.Teacher_RPC.Hosts, cfg.Teacher_RPC.Port,
		), grpc.WithInsecure())
			
		if err != nil {
			return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",

				cfg.Teacher_RPC.Hosts, cfg.Teacher_RPC.Port, err,
			)
		}




		connections[_adminRpc] = pb.NewUserServiceClient(connAdmin)
		connections[_authRpc] = authpb.NewAuthServiceClient(connAuth)
		connections[_studentRpc] = sp.NewTestServiceClient(connStudent)
		connections[_teacherRpc] = tp.NewTeacherServiceClient(connTeacher)
	}
	return &GRPCClient{
		cfg:         cfg,
		connections: connections,
	}, nil
}

func (g *GRPCClient) UserService() pb.UserServiceClient {
	return g.connections[_adminRpc].(pb.UserServiceClient)
}
func (g *GRPCClient) AuthService() authpb.AuthServiceClient {
	return g.connections[_authRpc].(authpb.AuthServiceClient)
}
func (g *GRPCClient) StudentService() sp.TestServiceClient {
	return g.connections[_studentRpc].(sp.TestServiceClient)
}

func (g *GRPCClient) TeacherService() tp.TeacherServiceClient {
	return g.connections[_teacherRpc].(tp.TeacherServiceClient)
}
