package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/pb"
)

type userService struct{}

func NewUserService() pb.UsersServer {
	return userService{}
}

// GetUsers
func (s userService) GetUsers(ctx context.Context, req *pb.EmptyReq) (*pb.GetUsersResponse, error) {
	u := &pb.User{
		Name: "aaaaaaaa",
		Age:  111,
	}

	users := []*pb.User{u}

	return &pb.GetUsersResponse{
		Users: users,
	}, nil
}

func (s userService) Sample(ctx *pb.EmptyReq, server pb.Users_SampleServer) error {
	var res pb.SampleResponse
	i := 0
	for {
		fmt.Printf("Cound: %d \n", i)
		res = pb.SampleResponse{
			Id: int32(i),
		}
		server.Send(&res)
		time.Sleep(time.Second)
		i++
		if i > 100 {
			break
		}
	}

	return nil
}
