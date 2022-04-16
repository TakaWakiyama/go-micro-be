package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/pb"
	"github.com/TakaWakiyama/forcusing-backend/cmd/user/service"
	"google.golang.org/grpc"
)

func main() {
	os.Environ()

	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	service := service.NewUserService()
	server := grpc.NewServer()
	pb.RegisterUsersServer(server, service)

	log.Printf("Listening on port %s... \n", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
