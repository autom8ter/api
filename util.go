package api

import (
	"context"
	"github.com/autom8ter/objectify"
	"os"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", os.Environ())
}

var Util = objectify.Default()

var Context context.Context
