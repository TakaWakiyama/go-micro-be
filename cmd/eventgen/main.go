package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/TakaWakiyama/forcusing-backend/cmd/eventgen/eventgen"
	"github.com/google/uuid"
)

type service struct{}

func (s service) HelloWorld(ctx context.Context, req *eventgen.HelloWorldRequest) error {
	fmt.Printf("HelloWorld Event req: %+v\n", req)
	return nil
}

func (s service) HogeEvent(ctx context.Context, req *eventgen.HogeEventRequest) error {
	fmt.Printf("H oge Event req : %+v\n", req)
	return nil
}

func main() {
	ctx := context.Background()
	proj := "forcusing"

	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	t, _ := eventgen.GetOrCreateTopicIfNotExists(client, "helloworldtopic")
	client.CreateSubscription(ctx, "helloworldsubscription", pubsub.SubscriptionConfig{
		Topic:       t,
		AckDeadline: 30 * time.Second,
	})

	fun := os.Getenv("PFUNC")
	if fun == "" {
		fmt.Println("Service Start")
		s := service{}
		eventgen.Run(s, client)
	} else {
		c := eventgen.NewHelloWorldServiceClient(client)
		msg := uuid.New().String()
		eid, err := c.PublishHelloWorld(ctx, &eventgen.HelloWorldRequest{
			Name:      "Taka",
			Message:   msg,
			Timestamp: time.Now().Unix(),
		})
		fmt.Printf("eid: %v\n", eid)
		fmt.Printf("err: %v\n", err)
	}
}

// export GOOGLE_APPLICATION_CREDENTIALS="/Users/wakiyamasora/Documents/product/forcusing/backend/credential.json"
