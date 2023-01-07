package main

import (
	"errors"

	"github.com/TakaWakiyama/forcusing-backend/cmd/protoc-gen-go-event/option"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// https://github.com/infobloxopen/protoc-gen-gorm/blob/main/plugin/plugin.go
// https://cloud.google.com/pubsub/docs/samples/pubsub-publish-proto-messages
// https://github.com/golang/protobuf/issues/1260
// protoc --go_out=. --go_opt=paths=source_relative  --go-event_out=. --go-event_opt=paths=source_relative sample.proto

const (
	GO_IMPORT_FMT     = "fmt"
	GO_IMPORT_CONTEXT = "context"
	GO_IMPORT_JSON    = "encoding/json"
	GO_IMPORT_PUBSUB  = "cloud.google.com/go/pubsub"
)

var allImportMods = []string{
	GO_IMPORT_FMT,
	GO_IMPORT_CONTEXT,
	GO_IMPORT_JSON,
	GO_IMPORT_PUBSUB,
}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_event.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated  by protoc-gen-go-event. DO NOT EDIT.")
	g.P("// versions:")
	// g.P("//  protoc-gen-go ", file) // TODO: get version from protogen
	// g.P("//  protoc        v3.21.9")                            // TODO: get version from protogen
	g.P("// source: ", file.Proto.GetName())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("import (")
	for _, mod := range allImportMods {
		g.P(`"`, mod, `"`)
	}
	g.P(")")

	if len(file.Services) == 0 {
		return g
	}
	svc := file.Services[0]

	svcName := svc.GoName
	// Create Client Code

	// Create Server Code
	// Create interface
	g.P("type ", svcName, " interface {")
	for _, m := range svc.Methods {
		// g.P("// ", m.Comments.Leading, " ", m.Comments.Trailing)
		g.P(m.GoName, "(ctx context.Context, req *", m.Input.GoIdent, ") error ")
	}
	g.P("}")

	// Create Run Function
	g.P("func Run(service ", svcName, ", client *pubsub.Client) {")
	g.P("ctx := context.Background()")
	for _, m := range svc.Methods {
		g.P("if err := listen", m.GoName, "(ctx, service, client); err != nil {")
		g.P("panic(err)")
		g.P("}")
	}
	g.P("}")
	// Create listen function
	for _, m := range svc.Methods {
		opt, _ := getPubSubOption(m)
		g.P("func listen", m.GoName, "(ctx context.Context, service ", svcName, ", client *pubsub.Client) error {")
		g.P("subscriptionName := ", `"`, opt.Subscription, `"`)
		g.P("topicName := ", `"`, opt.Topic, `"`)
		g.P("//", "TODO: メッセージの処理時間の延長を実装する必要がある")
		// https://christina04.hatenablog.com/entry/cloud-pubsub
		g.P(`callback := func(ctx context.Context, msg *pubsub.Message) {
			msg.Ack()
		`)
		g.P("var event ", m.Input.GoIdent)
		g.P("if err := json.Unmarshal(msg.Data, &event); err != nil {")
		g.P("fmt.Println(err)")
		// error 処理
		g.P("}")
		g.P("if err := service.", m.GoName, "(ctx, &event); err != nil {")
		// 再送信させる
		g.P("msg.Nack()")
		g.P("}")
		g.P("}")
		g.P("err := pullMsgs(ctx, client, subscriptionName, topicName, callback)")
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P("return nil")
		g.P("}")
	}

	// Create pullMsgs function
	g.P(`
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
	`)

	return g
}

func getPubSubOption(m *protogen.Method) (*option.PubSubOption, error) {
	options := m.Desc.Options().(*descriptorpb.MethodOptions)
	ext := proto.GetExtension(options, option.E_PubSubOption)
	opt, ok := ext.(*option.PubSubOption)
	if !ok {
		return nil, errors.New("no pubsub option")
	}
	return opt, nil
}
