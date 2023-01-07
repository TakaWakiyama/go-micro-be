// Code generated  by protoc-gen-go-event. DO NOT EDIT.
// versions:
// source: sample.proto

package eventgen

import (
	"fmt"
	"context"
	"encoding/json"
	"cloud.google.com/go/pubsub"
)

type HelloWorldService interface {
	HelloWorld(ctx context.Context, req *HelloWorldRequest) error
	HogeEvent(ctx context.Context, req *HogeEventRequest) error
}

func Run(service HelloWorldService, client *pubsub.Client) {
	ctx := context.Background()
	if err := listenHelloWorld(ctx, service, client); err != nil {
		panic(err)
	}
	if err := listenHogeEvent(ctx, service, client); err != nil {
		panic(err)
	}
}
func listenHelloWorld(ctx context.Context, service HelloWorldService, client *pubsub.Client) error {
	subscriptionName := "helloworldsubscription"
	topicName := "helloworldtopic"
	// TODO: メッセージの処理時間の延長を実装する必要がある
	callback := func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()

		var event HelloWorldRequest
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Println(err)
		}
		if err := service.HelloWorld(ctx, &event); err != nil {
			msg.Nack()
		}
	}
	err := pullMsgs(ctx, client, subscriptionName, topicName, callback)
	if err != nil {
		return err
	}
	return nil
}
func listenHogeEvent(ctx context.Context, service HelloWorldService, client *pubsub.Client) error {
	subscriptionName := "hoge"
	topicName := "hoge"
	// TODO: メッセージの処理時間の延長を実装する必要がある
	callback := func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()

		var event HogeEventRequest
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Println(err)
		}
		if err := service.HogeEvent(ctx, &event); err != nil {
			msg.Nack()
		}
	}
	err := pullMsgs(ctx, client, subscriptionName, topicName, callback)
	if err != nil {
		return err
	}
	return nil
}

func pullMsgs(ctx context.Context, client *pubsub.Client, subScriptionName, topicName string, callback func(context.Context, *pubsub.Message)) error {
	sub := client.Subscription(subScriptionName)
	// topic := client.Topic(topicName)
	fmt.Printf("topicName: %v\n", topicName)
	err := sub.Receive(ctx, callback)
	if err != nil {
		return err
	}
	return nil
}
