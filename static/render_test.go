package static_test

import (
	"fmt"
	"github.com/autom8ter/api/static"
	"testing"
)

func TestRenderer_Functions(t *testing.T) {
	s := static.NewRenderer()
	fmt.Println(s.Functions())
}
