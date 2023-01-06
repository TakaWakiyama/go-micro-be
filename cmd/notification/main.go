package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

/* proto input
message HelloWorldRequest {
	string name = 1;
	string message = 2;
	int64 timestamp = 3;
	optinal string subscription = 10001;
	optinal string topic = 10002;
}


message HogeEventRequest {
	string hoge = 1;
	string message = 2;
	int64 timestamp = 3;
	int32 count = 4;
	optinal string subscription = 10001;
	optinal string topic = 10002;
}

service HelloWorldService {
	rpc HelloWorld(HelloWorldRequest) returns (EmptyResponse) {}
	rpc HogeEvent(HogeEventRequest) returns (EmptyResponse) {}
}
*/

type HellowWorldEvent struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type HogeEvent struct {
	Hoge      string `json:"hoge"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Count     int32  `json:"count"`
}

type helloWorldService struct {
}

func (s *helloWorldService) HelloWorld(ctx context.Context, req *HellowWorldEvent) error {
	fmt.Printf("req: %+v\n", req)
	fmt.Printf("ctx: %v\n", ctx)
	return nil
}

func (s *helloWorldService) HogeEvent(ctx context.Context, req *HogeEvent) error {
	fmt.Printf("req: %+v\n", req)
	fmt.Printf("ctx: %v\n", ctx)
	return nil
}

type HelloWorldService interface {
	HelloWorld(ctx context.Context, req *HellowWorldEvent) error
	HogeEvent(ctx context.Context, req *HogeEvent) error
}

func newHelloWorldService() HelloWorldService {
	return &helloWorldService{}
}

func main() {
	s := newHelloWorldService()
	run(s)
}

func run(service HelloWorldService) {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	if projectID == "" {
		projectID = "forcusing"
	}

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	// AutoGenerate
	if err := listenHelloWorldEvent(ctx, client, service); err != nil {
		// ER
		panic(err)
	}
	if err := listenHogeEvent(ctx, client, service); err != nil {
		panic(err)
	}
}

func listenHelloWorldEvent(ctx context.Context, client *pubsub.Client, service HelloWorldService) error {
	subscriptionName := "helloworld-subscription" // TODO: AutoGenerate
	topicName := "helloworld-topic"               // TODO: AutoGenerate

	// TODO: メッセージの処理時間の延長を実装する必要がある
	// https://christina04.hatenablog.com/entry/cloud-pubsub
	callback := func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		var event HellowWorldEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Println(err)
			// error 処理
		}
		if err := service.HelloWorld(ctx, &event); err != nil {
			// 再送信させる
			msg.Nack()
		}
	}
	err := pullMsgs(ctx, client, subscriptionName, topicName, callback)
	if err != nil {
		return err
	}

	return nil
}

func listenHogeEvent(ctx context.Context, client *pubsub.Client, service HelloWorldService) error {
	subscriptionName := "hoge-subscription" // TODO: AutoGenerate
	topicName := "hoge-topic"               // TODO: AutoGenerate

	// TODO: メッセージの処理時間の延長を実装する必要がある
	// https://christina04.hatenablog.com/entry/cloud-pubsub
	callback := func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		var event HogeEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Println(err)
			// error 処理
		}
		if err := service.HogeEvent(ctx, &event); err != nil {
			// 再送信させる
			msg.Nack()
		}
	}
	err := pullMsgs(ctx, client, subscriptionName, topicName, callback)
	if err != nil {
		return err
	}

	return nil
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

// 1. Topicを作成する(optional)
// 2. Subscriptionを作成する
// 3. listenする
// 4. 関数に渡す

// Publisher側の実装
func newPubsubClient() *pubsub.Client {
	ctx := context.Background()
	proj := "forcusing"

	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	return client
}

type HelloWorldClent interface {
	PublishHelloWorldEvent(ctx context.Context, event *HellowWorldEvent) (string, error)
	PublishHogeEvent(ctx context.Context, event *HogeEvent) (string, error)
}

type helloWorldClient struct {
	client *pubsub.Client
}

func NewHelloWorldClient() *helloWorldClient {
	return &helloWorldClient{
		client: newPubsubClient(),
	}
}

func (c *helloWorldClient) PublishHelloWorldEvent(ctx context.Context, event *HellowWorldEvent) (string, error) {
	topic := "helloworld-topic"
	return c.publish(topic, event)
}

func (c *helloWorldClient) PublishHogeEvent(ctx context.Context, event *HogeEvent) (string, error) {
	topic := "hoge-topic"
	return c.publish(topic, event)
}

func (c *helloWorldClient) publish(topic string, event any) (string, error) {
	ctx := context.Background()
	t := c.client.Topic(topic)

	// json encode event
	ev, err := json.Marshal(event)
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
