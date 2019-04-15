package util

import (
	"context"
	"github.com/autom8ter/api/go/api"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"google.golang.org/grpc"
	"os"
	"time"
)

var ctx = context.Background()

type EnvData struct {
	Vars []string
}

func ServeStringsTest() (*api.String, error) {
	go engine.Default("tcp", "localhost:3000", true).With(
		config.WithPlugins(
			api.NewStringServiceServerFunctions(&api.Msg{
				Id:   api.Util.UUID(),
				Meta: make(map[string]string),
				Data: api.Util.MarshalJSON(EnvData{
					Vars: os.Environ(),
				}),
				PublishTime: time.Now().String(),
			}),
		),
	).Serve()
	c, err := api.NewClientSet(ctx, "localhost:3000", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	resp, err := c.Strings.RenderJSON(ctx, &api.String{
		Text: "Message Data: ID={{.Id}} Data={{.Data}} PublishTime={{.PublishTime}}",
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
