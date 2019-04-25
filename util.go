package api

import (
	"context"
	"github.com/autom8ter/objectify"
	"net/http"
	"os"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", os.Environ())
}

var Util = objectify.Default()

var Context context.Context

var httpClient *http.Client

func ENVFromContext() []string {
	return Context.Value("env").([]string)
}
