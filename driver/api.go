//go:generate godocdown -o README.md

package driver

import (
	"github.com/autom8ter/objectify"
	"github.com/golang/protobuf/proto"
)

var util = objectify.Default()

type Grouping interface {
	Categorizer
	Identifier
}

type Identifier interface {
	Identifier() string
}

type Categorizer interface {
	Category() string
}

type JSON interface {
	JSONString() string
}

type Metadata interface {
	Meta() map[string]string
}

type JSONMessage interface {
	MetaGrouping
	JSON
}

type Message interface {
	MetaGrouping
	String() string
}

type ProtoMessage interface {
	MetaGrouping
	proto.Message
}

type MetaGrouping interface {
	Grouping
	Metadata
}
