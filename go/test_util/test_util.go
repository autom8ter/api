package test_util

import (
	"github.com/autom8ter/api/go/api"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
)

func ServeStrings(strings api.StringServiceServerFunctions) {
	if err := engine.Default("tcp", "localhost:3000", true).With(
		config.WithPlugins(
			strings,
		),
	).Serve(); err != nil {
		api.Util.Fatalln("functions:", err.Error())
	}
}

