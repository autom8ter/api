package util_test

import (
	"github.com/autom8ter/api/go/api"
	"github.com/autom8ter/api/go/util"

	"testing"
)

func TestStringService(t *testing.T) {
	resp, err := util.ServeStringsTest()
	if err != nil {
		t.Fatal(err.Error())
	}
	api.Util.Entry().Println(resp.Text)
}
