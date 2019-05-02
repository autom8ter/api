# driver
--
    import "github.com/autom8ter/api/driver"


## Usage

#### type CallbackFunc

```go
type CallbackFunc func(resp *http.Response) error
```


#### func (CallbackFunc) Callback

```go
func (c CallbackFunc) Callback(r *http.Response) error
```

#### func (CallbackFunc) CallbackWithRoundTrip

```go
func (c CallbackFunc) CallbackWithRoundTrip(req *http.Request, roundTrip http.RoundTripper) error
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


#### type Messenger

```go
type Messenger interface {
	AsMessenger(meta map[string]string) Message
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

#### func (RoundTripperFunc) RoundTripWithCallback

```go
func (r RoundTripperFunc) RoundTripWithCallback(req *http.Request, callback CallbackFunc) error
```

#### type WebTasker

```go
type WebTasker interface {
	http.RoundTripper
	Callback(r *http.Response) error
}
```
