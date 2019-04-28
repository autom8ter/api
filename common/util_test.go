package common_test

import (
	"fmt"
	"github.com/autom8ter/api/common"
	"testing"
)

func TestStringArray_Last(t *testing.T) {
	arr := common.ToStringArray([]string{"first", "second", "third"})

	fmt.Println(arr.Last())
}

func TestString_AddLine(t *testing.T) {
	var text = common.ToString("Line1")
	text.AddLine("Line2")
	text.AddLine("Line3")
	text.AddLine("Line4")

	text.Println()
}
