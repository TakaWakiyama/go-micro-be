package main

import (
	"log"
	"net"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/pb"
	"github.com/TakaWakiyama/forcusing-backend/cmd/user/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	service := service.NewUserService()
	server := grpc.NewServer()
	pb.RegisterUsersServer(server, service)

	log.Println("Listening on port 8080...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
