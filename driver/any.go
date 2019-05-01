package driver

import "github.com/golang/protobuf/ptypes/any"

type Any struct {
	*any.Any
}

func (a *Any) JSONString() string {
	return string(util.MarshalJSONPB(a))
}

func (a *Any) Category() string {
	return a.TypeUrl
}
