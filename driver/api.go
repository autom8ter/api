//go:generate godocdown -o README.md

package driver

import "net/http"

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
	Callback(r *http.Response) error
}

type RoundTripperFunc func(req *http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

func (r RoundTripperFunc) RoundTripWithCallback(req *http.Request, callback CallbackFunc) error {
	resp, err := r(req)
	if err != nil {
		return err
	}
	return callback(resp)
}

type CallbackFunc func(resp *http.Response) error

func (c CallbackFunc) Callback(r *http.Response) error {
	return c(r)
}

func (c CallbackFunc) CallbackWithRoundTrip(req *http.Request, roundTrip http.RoundTripper) error {
	resp, err := roundTrip.RoundTrip(req)
	if err != nil {
		return err
	}
	return c(resp)
}
