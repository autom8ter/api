# functions
--
    import "github.com/autom8ter/api/functions"


## Usage

#### type CallbackFunc

```go
type CallbackFunc func(w io.Writer, r io.Reader) error
```


#### func (CallbackFunc) Callback

```go
func (c CallbackFunc) Callback(w io.Writer, r io.Reader) error
```

#### type ErrorLoggerFunc

```go
type ErrorLoggerFunc func(err error)
```


#### func (ErrorLoggerFunc) Err

```go
func (e ErrorLoggerFunc) Err(err error)
```

#### type ErrorWriterFunc

```go
type ErrorWriterFunc func(w io.Writer, err error)
```


#### func (ErrorWriterFunc) HandleError

```go
func (e ErrorWriterFunc) HandleError(w io.Writer, err error)
```

#### type GRPCFunc

```go
type GRPCFunc func(s *grpc.Server)
```


#### func (GRPCFunc) Chain

```go
func (g GRPCFunc) Chain(plugs ...driver.Plugin) GRPCFunc
```

#### func (GRPCFunc) RegisterWithServer

```go
func (g GRPCFunc) RegisterWithServer(s *grpc.Server)
```

#### type HTTPHandlerFunc

```go
type HTTPHandlerFunc func(w http.ResponseWriter, r *http.Request)
```


#### func (HTTPHandlerFunc) Chain

```go
func (h HTTPHandlerFunc) Chain(handlers ...http.Handler) HTTPHandlerFunc
```

#### func (HTTPHandlerFunc) ServeHTTP

```go
func (h HTTPHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

#### type RoundTripperFunc

```go
type RoundTripperFunc func(req *http.Request) (*http.Response, error)
```


#### func (RoundTripperFunc) RoundTrip

```go
func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error)
```
