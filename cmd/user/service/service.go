package service

import (
	"context"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/pb"
)

type UserService interface {
	GetUsers(context.Context, *pb.EmptyReq) (*pb.GetUsersResponse, error)
}

type userService struct{}

func NewUserService() UserService {
	return userService{}
}

func (s userService) GetUsers(ctx context.Context, req *pb.EmptyReq) (*pb.GetUsersResponse, error) {
	u := &pb.User{
		Name: "a",
		Age:  122,
	}

	res := &pb.GetUsersResponse{
		Users: []*pb.User{u, u, u},
	}

	return res, nil
}
