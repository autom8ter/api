package api

import (
	"context"
	"github.com/autom8ter/objectify"
	"net/http"
	"os"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", os.Environ())
	HTTPClient = http.DefaultClient
}

var Util = objectify.Default()

var Context context.Context

var HTTPClient *http.Client

func ENVFromContext() []string {
	return Context.Value("env").([]string)
}
