// Code generated  by protoc-gen-go-event. DO NOT EDIT.
// versions:
// source: sample.proto

package eventgen

import (
	"fmt"
	"context"
	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
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
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
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

type HelloWorldServiceClient interface {
	PublishHelloWorld(ctx context.Context, req *HelloWorldRequest) (string, error)
	PublishHogeEvent(ctx context.Context, req *HogeEventRequest) (string, error)
}

type innerHelloWorldServiceClient struct {
	client *pubsub.Client
}

func NewHelloWorldServiceClient(client *pubsub.Client) *innerHelloWorldServiceClient {
	return &innerHelloWorldServiceClient{
		client: client,
	}
}

func (c *innerHelloWorldServiceClient) publish(topic string, event protoreflect.ProtoMessage) (string, error) {
	ctx := context.Background()
	// TODO: メモリに持っておく
	t := c.client.Topic(topic)

	ev, err := proto.Marshal(event)
	if err != nil {
		return "", err
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: ev,
	})
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *innerHelloWorldServiceClient) PublishHelloWorld(ctx context.Context, req *HelloWorldRequest) (string, error) {
	return c.publish("helloworldtopic", req)
}
func (c *innerHelloWorldServiceClient) PublishHogeEvent(ctx context.Context, req *HogeEventRequest) (string, error) {
	return c.publish("hoge", req)
}

// GetOrCreateTopicIfNotExists: topicが存在しない場合は作成する
func GetOrCreateTopicIfNotExists(c *pubsub.Client, topic string) (*pubsub.Topic, error) {
	ctx := context.Background()
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if ok {
		return t, nil
	}
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		return nil, err
	}
	return t, nil
}
