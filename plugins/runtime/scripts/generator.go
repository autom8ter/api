package scripts

import (
	"github.com/autom8ter/api/plugins/runtime/encoder"
	"github.com/autom8ter/api/plugins/runtime/generator"
	"github.com/autom8ter/api/util"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	ggdescriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"io/ioutil"
	"moul.io/protoc-gen-gotemplate/helpers"
	"os"
	"strings"
)

var (
	registry *ggdescriptor.Registry // some helpers need access to registry
)

const (
	boolTrue  = "true"
	boolFalse = "false"
)

type Config struct {
	Templates     string
	Destination   string
	Debug         bool
	All           bool
	SinglePackage bool
}

func (c *Config) UnmarshalRequest() generator.Script {
	return func(g *generator.Generator) {
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

func (c *Config) ParseArgs() generator.Script {
	return func(g *generator.Generator) {
		g.CommandLineParameters()
		// Parse parameters
		if parameter := g.Request().GetParameter(); parameter != "" {
			for _, param := range strings.Split(parameter, ",") {
				parts := strings.Split(param, "=")
				if len(parts) != 2 {
					util.Util.Entry().Warnf("Err: invalid parameter: %q", param)
					continue
				}
				switch parts[0] {
				case "templates":
					c.Templates = parts[1]
				case "destination":
					c.Destination = parts[1]
				case "single-package":
					switch strings.ToLower(parts[1]) {
					case boolTrue, "t":
						c.SinglePackage = true
					case boolFalse, "f":
					default:
						util.Util.Entry().Warnf("Err: invalid value for single-package-mode: %q", parts[1])
					}
				case "debug":
					switch strings.ToLower(parts[1]) {
					case boolTrue, "t":
						c.Debug = true
					case boolFalse, "f":
					default:
						util.Util.Entry().Warnf("Err: invalid value for debug: %q", parts[1])
					}
				case "all":
					switch strings.ToLower(parts[1]) {
					case boolTrue, "t":
						c.All = true
					case boolFalse, "f":
					default:
						util.Util.Entry().Warnf("Err: invalid value for debug: %q", parts[1])
					}
				default:
					util.Util.Entry().Warnf("Err: unknown parameter: %q", param)
				}
			}
		}
	}
}

func (c *Config) l() generator.Script {
	return func(g *generator.Generator) {
		tmplMap := make(map[string]*plugin_go.CodeGeneratorResponse_File)
		concatOrAppend := func(file *plugin_go.CodeGeneratorResponse_File) {
			if val, ok := tmplMap[file.GetName()]; ok {
				*val.Content += file.GetContent()
			} else {
				tmplMap[file.GetName()] = file
				g.Response().File = append(g.Response().File, file)
			}
		}

		if c.SinglePackage {
			registry = ggdescriptor.NewRegistry()
			pgghelpers.SetRegistry(registry)
			if err := registry.Load(g.Request()); err != nil {
				g.Err(err, "registry: failed to load the request")
			}
		}

		// Generate the encoders
		for _, file := range g.Request().GetProtoFile() {
			if c.All {
				if c.SinglePackage {
					if _, err := registry.LookupFile(file.GetName()); err != nil {
						g.Err(err, "registry: failed to lookup file %q", file.GetName())
					}
				}
				e := encoder.NewGenericTemplateBasedEncoder(c.Templates, file, c.Debug, c.Destination)
				for _, tmpl := range e.Files() {
					concatOrAppend(tmpl)
				}

				continue
			}

			for _, service := range file.GetService() {
				e := encoder.NewGenericServiceTemplateBasedEncoder(c.Templates, service, file, c.Debug, c.Destination)
				for _, tmpl := range e.Files() {
					concatOrAppend(tmpl)
				}
			}
		}
	}
}
