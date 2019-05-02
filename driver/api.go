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
	ErrorHandler
	Callbacker
}

type Callbacker interface {
	Callback(w io.Writer, r io.Reader) error
}

type ErrorHandler interface {
	HandleError(w io.Writer, err error)
}

type RoundTripperFunc func(req *http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

type CallbackFunc func(w io.Writer, r io.Reader) error

func (c CallbackFunc) Callback(w io.Writer, r io.Reader) error {
	return c(w, r)
}

type ErrorFunc func(w io.Writer, err error)

func (e ErrorFunc) HandleError(w io.Writer, err error) {
	e(w, err)
}