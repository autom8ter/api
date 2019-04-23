package static_test

import (
	"fmt"
	"github.com/autom8ter/static"
	"net/http/httptest"
	"testing"
)

func TestRenderer_Functions(t *testing.T) {
	s := static.NewRenderer()
	fmt.Println(s.Functions())
}

func TestHandler(t *testing.T) {
	httptest.NewServer()
}
