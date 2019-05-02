# driver
--
    import "github.com/autom8ter/api/driver"


## Usage

#### type CallbackFunc

```go
type CallbackFunc func(w io.Writer, r io.Reader) error
```


#### func (CallbackFunc) Callback

```go
func (c CallbackFunc) Callback(w io.Writer, r io.Reader) error
```

#### type Callbacker

```go
type Callbacker interface {
	Callback(w io.Writer, r io.Reader) error
}
```


#### type Categorizer

```go
type Categorizer interface {
	GetCategory() string
}
```


#### type DataMapper

```go
type DataMapper interface {
	DataMap() map[string]interface{}
}
```


#### type Debugger

```go
type Debugger interface {
	Debugf(format string)
}
```


#### type ErrorFunc

```go
type ErrorFunc func(w io.Writer, err error)
```


#### func (ErrorFunc) HandleError

```go
func (e ErrorFunc) HandleError(w io.Writer, err error)
```

#### type ErrorHandler

```go
type ErrorHandler interface {
	HandleError(w io.Writer, err error)
}
```


#### type Grouping

```go
type Grouping interface {
	Categorizer
	Identifier
}
```


#### type Identifier

```go
type Identifier interface {
	GetIdentifier() string
}
```


#### type Message

```go
type Message interface {
	MetaGrouping
	String() string
}
```


#### type MetaGrouping

```go
type MetaGrouping interface {
	Grouping
	Metadata
}
```


#### type Metadata

```go
type Metadata interface {
	GetMeta() map[string]string
}
```


#### type RoundTripperFunc

```go
type RoundTripperFunc func(req *http.Request) (*http.Response, error)
```


#### func (RoundTripperFunc) RoundTrip

```go
func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error)
```

#### type WebTasker

```go
type WebTasker interface {
	http.RoundTripper
	ErrorHandler
	Callbacker
}
```
