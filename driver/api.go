//go:generate godocdown -o README.md

package driver

import (
	"io"
	"net/http"
)

type Grouping interface {
	Categorizer
	Identifier
}

type Identifier interface {
	GetIdentifier() string
}

type Categorizer interface {
	GetCategory() string
}

type Metadata interface {
	GetMeta() map[string]string
}

type Message interface {
	MetaGrouping
	String() string
}

type MetaGrouping interface {
	Grouping
	Metadata
}

type DataMapper interface {
	DataMap() map[string]interface{}
}

type Debugger interface {
	Debugf(format string)
}

type WebTasker interface {
	http.RoundTripper
	ErrorWriter
	Callbacker
}

type Callbacker interface {
	Callback(w io.Writer, r io.Reader) error
}

type ErrorWriter interface {
	HandleError(w io.Writer, err error)
}

type ErrorLogger interface {
	Err(err error)
}
