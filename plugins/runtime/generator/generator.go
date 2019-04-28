package generator

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"io/ioutil"
	"os"
)

type Script func(g *Generator)

type Generator struct {
	gen *generator.Generator
}

func New() *Generator {
	return &Generator{
		generator.New(),
	}
}

func (g *Generator) Err(err error, args ...string) {
	if g.gen == nil {
		g.gen = generator.New()
	}
	g.gen.Error(err, args...)
}

func (g *Generator) Exit(args ...string) {
	if g.gen == nil {
		g.gen = generator.New()
	}
	g.gen.Fail(args...)
}

func (g *Generator) Request() *plugin_go.CodeGeneratorRequest {
	if g.gen == nil {
		g.gen = generator.New()
	}
	return g.gen.Request
}

func (g *Generator) Response() *plugin_go.CodeGeneratorResponse {
	if g.gen == nil {
		g.gen = generator.New()
	}
	return g.gen.Response
}

func (g *Generator) CommandLineParameters() {
	g.gen.CommandLineParameters(g.gen.Request.GetParameter())
}

func (g *Generator) Run(scripts ...Script) {
	setup()(g)
	for _, s := range scripts {
		s(g)
	}
	closer()(g)
}

func setup() Script {
	return func(g *Generator) {
		if g.gen == nil {
			g.gen = generator.New()
		}
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			g.Err(err, "reading input")
		}

		if err = proto.Unmarshal(data, g.Request()); err != nil {
			g.Err(err, "parsing input proto")
		}

		if len(g.Request().FileToGenerate) == 0 {
			g.Exit("no files to generate")
		}
	}
}

func closer() Script {
	return func(g *Generator) {
		// Generate the protobufs
		g.gen.GenerateAllFiles()

		data, err := proto.Marshal(g.Response())
		if err != nil {
			g.Err(err, "failed to marshal output proto")
		}

		_, err = os.Stdout.Write(data)
		if err != nil {
			g.Err(err, "failed to write output proto")
		}
	}
}
