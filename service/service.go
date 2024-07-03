package service

import (
	"context"
	"fmt"
	"strconv"
	"tcpserver/database"
	"tcpserver/models"
	"tcpserver/proto"
)

type Server struct {
	proto.UnimplementedMyAppServer
}

func (s *Server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var user models.User
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("could not find user with ID %d: %v", id, err)
	}
	return &proto.GetUserResponse{
		Id:       fmt.Sprint(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	var user models.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}
	return &proto.CreateUserResponse{
		Id: fmt.Sprint(user.ID),
	}, nil
}
