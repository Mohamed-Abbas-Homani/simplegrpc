package service

import (
	"context"
	"tcpserver/proto"
)

type Server struct {
	proto.UnimplementedMyAppServer
}

func (s *Server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	return &proto.GetUserResponse{
		Id:       "1",
		Username: "mash",
		Email:    "mash@gmail.com",
	}, nil
}
