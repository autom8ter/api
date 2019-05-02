# driver
--
    import "github.com/autom8ter/api/driver"


## Usage

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
