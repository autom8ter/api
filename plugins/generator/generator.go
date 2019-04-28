package generator

import (
	"fmt"
	"github.com/autom8ter/api/plugins/generator/base"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"io"
)

type DescriptorFunc func(descriptorProto *descriptor.FileDescriptorProto)

type Generator struct {
	base *base.Generator
}

func Default() *Generator {
	return &Generator{base: base.Default()}
}

func New(r io.Reader, w io.Writer) *Generator {
	return &Generator{base: base.New(r, w)}
}

func (g *Generator) GoFileName(protoName *string, extension string) string {
	return g.base.ProtoFileBaseName(*protoName) + extension
}

func (g *Generator) Init(pkgname string) {
	g.base.WriteString(fmt.Sprintf("package %s", pkgname))
}
func (g *Generator) Imports(importString string) {
	g.base.WriteString(fmt.Sprintf(`
import (
	%s
	pb "%s"
)`, importString, g.base.Param()["pkg_path"]))
}

func (g *Generator) Variables(varString string) {
	g.base.WriteString(`
var (
	%s
)`, varString)
}

func (g *Generator) Constants(constString string) {
	g.base.WriteString(`
const (
	%s
)`, constString)
}

func (g *Generator) ValidateArgs(args ...string) {
	args = append(args, "pkg_path")
	for _, a := range args {
		if _, ok := g.base.Param()[a]; !ok {
			g.base.Fatal(fmt.Sprintf("parameter %s is required", a))
		}
	}
}

func (g *Generator) Generate(handler base.Handler) {
	g.base.Generate(handler)
}

func (g *Generator) String() string {
	return g.base.Generator().String()
}

func (g *Generator) Response(filename *string, custom_extension string) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(g.GoFileName(filename, custom_extension)),
		Content: proto.String(g.String()),
	}
}

func (g *Generator) NewFileHandler(extension string, fns ...DescriptorFunc) base.HandlerFunc {
	return func(file *descriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
		for _, f := range fns {
			f(file)
		}
		return g.Response(file.Name, extension), nil

	}
}
