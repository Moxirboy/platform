package client

import (
	"fmt"
	configs "gateway/internal/config"
	pb "gateway/proto"
	"gateway/proto/authpb"
	"google.golang.org/grpc"
)

type IGRPCClient interface {
	UserService() pb.UserServiceClient
	AuthService() authpb.AuthServiceClient
}

type GRPCClient struct {
	cfg         configs.Config
	connections map[string]interface{}
}

const (
	_adminRpc = "admin_rpc"
	_authRpc  = "auth_rpc"
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

		connections[_adminRpc] = pb.NewUserServiceClient(connAdmin)
		connections[_authRpc] = authpb.NewAuthServiceClient(connAuth)
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
