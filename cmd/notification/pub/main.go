package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	proj := "forcusing"

	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	const topic = "my-topic"
	// List all the topics from the project.
	fmt.Println("Listing all topics from the project:")
	topics, err := list(client)
	if err != nil {
		log.Fatalf("Failed to list topics: %v", err)
	}
	for _, t := range topics {
		fmt.Println(t)
	}
	// Publish a text message on the created topic.
	if err := publish(client, topic, "hello world!"); err != nil {
		log.Fatalf("Failed to publish: %v", err)
	}
}

func list(client *pubsub.Client) ([]*pubsub.Topic, error) {
	ctx := context.Background()
	var topics []*pubsub.Topic
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}

func create(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	t, err := client.CreateTopic(ctx, topic)
	if err != nil {
		return err
	}
	fmt.Printf("Topic created: %v\n", t)
	return nil
}

func publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	t := client.Topic(topic)

	e := newCustomEvent()
	// json encode event
	ev, err := json.Marshal(e)
	if err != nil {
		return err
	}
	fmt.Printf("msg: %v\n", msg)

	result := t.Publish(ctx, &pubsub.Message{
		Data: ev,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

type customEvent struct {
	CreatedAt int64  `json:"created_at"`
	Message   string `json:"message"`
	Number    int32  `json:"number"`
}

func newCustomEvent() *customEvent {
	return &customEvent{
		CreatedAt: time.Now().Unix(),
		Message:   "Hello, World!",
		Number:    42,
	}
}
