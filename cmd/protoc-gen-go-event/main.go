package main

import (
	"flag"

	"github.com/TakaWakiyama/forcusing-backend/cmd/protoc-gen-go-event/option"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// https://github.com/infobloxopen/protoc-gen-gorm/blob/main/plugin/plugin.go
// https://cloud.google.com/pubsub/docs/samples/pubsub-publish-proto-messages
// https://github.com/golang/protobuf/issues/1260
// protoc --go_out=. --go_opt=paths=source_relative  --go-event_out=. --go-event_opt=paths=source_relative sample.proto

var font *string

func main() {
	var flags flag.FlagSet
	font = flags.String("font", "doom", "font list available in github.com/common-nighthawk/go-figure")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
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
	g.P(`import "context"`)

	if len(file.Services) == 0 {
		return g
	}
	svc := file.Services[0]

	svcName := svc.GoName
	g.P("type ", svcName, " interface {")
	for _, m := range svc.Methods {
		// g.P("// ", m.Comments.Leading, " ", m.Comments.Trailing)
		g.P(m.GoName, "(ctx context.Context, req *", m.Input.GoIdent, ") ", m.Output.GoIdent)
		options := m.Desc.Options().(*descriptorpb.MethodOptions)
		ext := proto.GetExtension(options, option.E_PubSubOption)
		opt, ok := ext.(*option.PubSubOption)
		if !ok {
			continue
		}
		g.P("//", opt.Topic, opt.Subscription)
	}
	g.P("}")

	return g
}
