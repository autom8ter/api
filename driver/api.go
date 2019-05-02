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

type Messenger interface {
	AsMessenger(meta map[string]string) Message
}

type WebTasker interface {
	http.RoundTripper
	Callback(w io.Writer, r *http.Response) error
}

type RoundTripperFunc func(req *http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

func (r RoundTripperFunc) RoundTripWithCallback(w io.Writer, req *http.Request, callback CallbackFunc) error {
	resp, err := r(req)
	if err != nil {
		return err
	}
	return callback(w, resp)
}

type CallbackFunc func(w io.Writer, resp *http.Response) error

func (c CallbackFunc) Callback(w io.Writer, r *http.Response) error {
	return c(w, r)
}

func (c CallbackFunc) CallbackWithRoundTrip(w io.Writer, req *http.Request, roundTrip http.RoundTripper) error {
	resp, err := roundTrip.RoundTrip(req)
	if err != nil {
		return err
	}
	return c(w, resp)
}
