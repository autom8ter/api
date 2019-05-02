# driver
--
    import "github.com/autom8ter/api/driver"


## Usage

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


#### type ErrorLogger

```go
type ErrorLogger interface {
	Err(err error)
}
```


#### type ErrorWriter

```go
type ErrorWriter interface {
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


#### type WebTasker

```go
type WebTasker interface {
	http.RoundTripper
	ErrorWriter
	Callbacker
}
```
