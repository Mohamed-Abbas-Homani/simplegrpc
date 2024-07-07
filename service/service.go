package service

import (
	"context"
	"tcpserver/interfaces"
	"tcpserver/models"
	"tcpserver/proto"
)

type Server struct {
	proto.UnimplementedMyAppServer
	UserRepo interfaces.UserRepository
}

func convertToProtoUser(modelUser *models.User) *proto.User {
	return &proto.User{
		Id:       uint64(modelUser.ID),
		Username: modelUser.Username,
		Email:    modelUser.Email,
	}
}

func convertToRole(modelRole *models.Role) *proto.Role {
	return &proto.Role{
		Name: modelRole.Name,
	}
}
func (s *Server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := s.UserRepo.GetUserById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.GetUserResponse{
		Id:       uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	users, err := s.UserRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*proto.User, len(users))
	for i, user := range users {
		protoUsers[i] = convertToProtoUser(user)
	}
	return &proto.GetUsersResponse{
		Users: protoUsers,
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	var user models.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.OrganizationID = uint(req.OrganizationId)

	createdUser, err := s.UserRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResponse{
		Id: uint64(createdUser.ID),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	var user models.User
	user.ID = uint(req.Id)
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.OrganizationID = uint(req.OrganizationId)
	updatedUser, err := s.UserRepo.Update(&user)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateUserResponse{
		Id: uint64(updatedUser.ID),
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if err := s.UserRepo.Delete(uint(req.Id)); err != nil {
		return nil, err
	}
	return &proto.DeleteUserResponse{
		Id: uint64(req.Id),
	}, nil
}

func (s *Server) GetFullUser(ctx context.Context, req *proto.GetFullUserRequest) (*proto.GetFullUserResponse, error) {
	fullUser, err := s.UserRepo.GetFullUserById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	protoRoles := make([]*proto.Role, len(fullUser.Roles))
	for i, role := range fullUser.Roles {
		protoRoles[i] = convertToRole(&role)
	}

	return &proto.GetFullUserResponse{
		Id:       uint64(fullUser.ID),
		Username: fullUser.Username,
		Email:    fullUser.Email,
		UserDetail: &proto.UserDetail{
			FullName:    fullUser.UserDetail.FullName,
			PhoneNumber: fullUser.UserDetail.PhoneNumber,
		},
		Roles: protoRoles,
		Organization: &proto.Organization{
			Name: fullUser.Organization.Name,
		},
	}, nil
}
