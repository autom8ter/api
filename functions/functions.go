//go:generate godocdown -o README.md

package functions

import (
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
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

type HTTPHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (h HTTPHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func (h HTTPHandlerFunc) Chain(handlers ...http.Handler) HTTPHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		for _, this := range handlers {
			this.ServeHTTP(w, r)
		}
	}
}

type GRPCFunc func(s *grpc.Server)

func (g GRPCFunc) RegisterWithServer(s *grpc.Server) {
	g(s)
}

func (g GRPCFunc) Chain(plugs ...driver.Plugin) GRPCFunc {
	return func(s *grpc.Server) {
		g.RegisterWithServer(s)
		for _, p := range plugs {
			p.RegisterWithServer(s)
		}
	}
}
