package main

import (
	"flag"

	"github.com/common-nighthawk/go-figure"
	"google.golang.org/protobuf/compiler/protogen"
)

// protoc --go_out=. --go_opt=paths=source_relative --eventgen_out=. --eventgen_opt=paths=source_relative event/sample.proto
// protoc --go_out=. --go_opt=paths=source_relative \ --go-ascii_out=. --go-ascii_opt=paths=source_relative example/example.proto

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
	filename := file.GeneratedFilenamePrefix + "_ascii.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated  by protoc-gen-go-ascii. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, msg := range file.Messages {
		fig := figure.NewFigure(msg.GoIdent.GoName, *font, false)

		g.P("func (x *", msg.GoIdent, ") Ascii() string {")
		g.P("return `", fig.String(), "`")
		g.P("}")
	}

	return g
}
