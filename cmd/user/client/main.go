package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/pb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	addr := "localhost:8080"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUsersClient(conn)
	ctx := context.Background()
	watcher, err := client.Sample(ctx, &pb.EmptyReq{})
	if err != nil {
		panic(err)
	}

	for {
		res, err := watcher.Recv()

		if err != nil && err.Error() == "EOF" {
			fmt.Println("Stream END")
			return
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("Response %+v \n", res)
	}
}
