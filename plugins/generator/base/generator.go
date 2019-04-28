package base

import (
	"github.com/autom8ter/objectify"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")

type Handler interface {
	HandleFile(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type HandlerFunc func(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)

func (h HandlerFunc) HandleFile(r *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	return h(r)
}

func NewHandlerFunc(fn func(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)) HandlerFunc {
	return fn
}

var util = objectify.Default()

// Default creates a new base generator
func Default() *Generator {
	return &Generator{
		gen:    gen.New(),
		reader: os.Stdin,
		writer: os.Stdout,
	}
}

// New creates a new base generator
func New(r io.Reader, w io.Writer) *Generator {
	return &Generator{
		gen:    gen.New(),
		reader: r,
		writer: w,
	}
}

type Generator struct {
	gen    *gen.Generator
	indent string
	reader io.Reader
	writer io.Writer
}

// Write prints the arguments to the generated output.  It handles strings and int32s, plus
// handling indirections because they may be *string, etc.
func (g *Generator) WriteString(str ...string) {
	g.gen.WriteString(g.indent)
	for _, v := range str {
		g.gen.WriteString(v)
	}
	g.gen.WriteByte('\n')
}

func (g *Generator) init() {
	g.gen.CommandLineParameters(g.gen.Request.GetParameter())

	g.gen.WrapTypes()

	g.gen.SetPackageNames()
	g.gen.BuildTypeNameMap()

	g.gen.GenerateAllFiles()
	g.gen.Reset()

}

func (g *Generator) generate(h Handler, request *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	response := new(plugin.CodeGeneratorResponse)
	for _, protoFile := range request.ProtoFile {
		file, err := h.HandleFile(protoFile)
		if err != nil {
			return response, err
		}
		response.File = append(response.File, file)
	}
	return response, nil
}

func (g *Generator) Generate(handlers ...Handler) {
	for _, h := range handlers {
		input, err := ioutil.ReadAll(g.reader)
		if err != nil {
			util.Entry().Fatalln(err.Error(), "reading input")
		}

		request := g.gen.Request
		if err := proto.Unmarshal(input, request); err != nil {
			util.Entry().Fatalln(err.Error(), "parsing input proto")
		}

		if len(request.FileToGenerate) == 0 {
			util.Entry().Fatalln("no files to generate")
		}

		g.init()

		response, err := g.generate(h, request)
		if err != nil {
			util.Entry().Fatalln(err, "failed to generate files from proto")
		}

		output, err := proto.Marshal(response)
		if err != nil {
			util.Entry().Fatalln(err, "failed to marshal output proto")
		}
		_, err = g.writer.Write(output)
		if err != nil {
			util.Entry().Fatalln(err, "failed to write output proto")
		}
	}
}

func (g *Generator) ProtoFileBaseName(name string) string {
	if ext := path.Ext(name); ext == ".proto" {
		name = name[:len(name)-len(ext)]
	}
	return name

}

func (g *Generator) Object(str string) gen.Object {
	g.gen.RecordTypeUse(str)
	return g.gen.ObjectNamed(str)
}

func (g *Generator) TypeName(str string) string {
	g.gen.RecordTypeUse(str)
	return g.gen.TypeName(g.Object(str))
}

func (g *Generator) Indent() { g.indent += "\t" }

func (g *Generator) UnIndent() {
	if len(g.indent) > 0 {
		g.indent = g.indent[1:]
	}
}

func (g *Generator) FatalIfErr(err error, args ...string) {
	if err != nil {
		util.Entry().Fatalln(err.Error(), args)
	}
}

func (g *Generator) Fatal(args ...string) {
	util.Entry().Fatalln(args)
}

func (g *Generator) Underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

func (g *Generator) ReceiverName(typeName string) string {
	if typeName == "" {
		return ""
	}
	return string([]rune(g.Underscore(typeName))[0])
}

func (g *Generator) Param() map[string]string {
	return g.gen.Param
}

func (g *Generator) Generator() *gen.Generator {
	return g.gen
}
