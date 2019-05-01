# driver
--
    import "github.com/autom8ter/api/driver"


## Usage

#### type Categorizer

```go
type Categorizer interface {
	Category() string
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
	Identifier() string
}
```


#### type JSON

```go
type JSON interface {
	JSONString() string
}
```


#### type JSONMessage

```go
type JSONMessage interface {
	MetaGrouping
	JSON
}
```


#### type JSONTask

```go
type JSONTask interface {
	JSON
	Identifier
	URL() string
	Method() string
	Headers() map[string]string
	Queue
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
	Meta() map[string]string
}
```


#### type ProtoMessage

```go
type ProtoMessage interface {
	MetaGrouping
	proto.Message
}
```


#### type Queue

```go
type Queue interface {
	GetQueLocation() string
	GetQueID() string
	ExecuteAtUnix() int64
}
```
