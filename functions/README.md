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

#### type RoundTripperFunc

```go
type RoundTripperFunc func(req *http.Request) (*http.Response, error)
```


#### func (RoundTripperFunc) RoundTrip

```go
func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error)
```
