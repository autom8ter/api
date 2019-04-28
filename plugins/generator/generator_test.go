package generator_test

import (
	"github.com/autom8ter/api/plugins/generator"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"testing"
)

func TestNew(t *testing.T) {
	g := generator.Default()
	g.Generate(g.NewFileHandler(".test.go", func(descriptorProto *descriptor.FileDescriptorProto) {

	}))
}
