package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
)

func main() {
	// pflag.String("subscriber", "example-subscription", "name of subscriber")
	// pflag.Parse()
	// viper.BindPFlags(pflag.CommandLine)
	// ctx := context.Background()
	// proj := "forcusing"
	// client, err := pubsub.NewClient(ctx, proj)
	// if err != nil {
	// log.Fatalf("Could not create pubsub Client: %v", err)
	// }
	// sub := "sample" // viper.GetString("subscriber") // retrieve values from viper instead of pflag
	// t := createTopicIfNotExists(client, "my-topic")
	// // Create a new subscription.
	// log.Println("start")
	// if err := pullMsgs(client, sub, t, false); err != nil {
	// log.Fatal(err)
	// }
	run()

}

func run() {

}

// func createTopicIfNotExists(c *pubsub.Client, topic string) *pubsub.Topic {
// ctx := context.Background()
// t := c.Topic(topic)
// ok, err := t.Exists(ctx)
// if err != nil {
// log.Fatal(err)
// }
// if ok {
// return t
// }
// t, err = c.CreateTopic(ctx, topic)
// if err != nil {
// log.Fatalf("Failed to create the topic: %v", err)
// }
// return t
// }

// func create(client *pubsub.Client, name string, topic *pubsub.Topic) error {
// ctx := context.Background()
// sub, err := client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
// Topic:       topic,
// AckDeadline: 20 * time.Second,
// })

// if err != nil {
// return err
// }
// fmt.Printf("Created subscription: %v\n", sub)
// // [END pubsub_create_pull_subscription]
// return nil
// }

func pullMsgs(client *pubsub.Client, name string, topic *pubsub.Topic) error {
	ctx := context.Background()

	var mu sync.Mutex
	received := 0
	sub := client.Subscription(name)
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		var event customEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Got message: %+v\n", event)
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
		// msg.Nack()
	})
	if err != nil {
		return err
	}
	return nil
}

type customEvent struct {
	CreatedAt int64  `json:"created_at"`
	Message   string `json:"message"`
	Number    int32  `json:"number"`
}

// 1. Topicを作成する
// 2. Subscriptionを作成する
// 3. listenする
// 4. 関数に渡す
