//go:generate godocdown -o README.md

package functions

import (
	"io"
	"net/http"
)

type RoundTripperFunc func(req *http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

type CallbackFunc func(w io.Writer, r io.Reader) error

func (c CallbackFunc) Callback(w io.Writer, r io.Reader) error {
	return c(w, r)
}

type ErrorWriterFunc func(w io.Writer, err error)

func (e ErrorWriterFunc) HandleError(w io.Writer, err error) {
	e(w, err)
}

type ErrorLoggerFunc func(err error)

func (e ErrorLoggerFunc) Err(err error) {
	e(err)
}
