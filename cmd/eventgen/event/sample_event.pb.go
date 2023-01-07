// Code generated  by protoc-gen-go-event. DO NOT EDIT.
// versions:
// source: sample.proto

package eventgen

import "context"

type HelloWorldService interface {
	HelloWorld(ctx context.Context, req *HelloWorldRequest) EmptyResponse
	// helloworldtopichelloworldsubscription
	HogeEvent(ctx context.Context, req *HogeEventRequest) EmptyResponse
	// hogehoge
}
