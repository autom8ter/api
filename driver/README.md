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


#### func  AsCategorizer

```go
func AsCategorizer(str string) Categorizer
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


#### func  AsIdentifier

```go
func AsIdentifier(str string) Identifier
```

#### type JSON

```go
type JSON interface {
	JSONString() string
}
```


#### func  AsJSON

```go
func AsJSON(str string) JSON
```

#### type JSONMessage

```go
type JSONMessage interface {
	MetaGrouping
	JSON
}
```


#### type Map

```go
type Map map[string]string
```


#### func (Map) Meta

```go
func (m Map) Meta() map[string]string
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


#### func  AsMeta

```go
func AsMeta(m map[string]string) Metadata
```

#### type ProtoMessage

```go
type ProtoMessage interface {
	MetaGrouping
	proto.Message
}
```


#### type String

```go
type String string
```


#### func (String) Bytes

```go
func (s String) Bytes(str string) []byte
```

#### func (String) Category

```go
func (s String) Category() string
```

#### func (String) Contains

```go
func (s String) Contains(str string) bool
```

#### func (String) Identifier

```go
func (s String) Identifier() string
```

#### func (String) JSONString

```go
func (s String) JSONString() string
```

#### func (String) ParseURL

```go
func (s String) ParseURL() (*url.URL, error)
```

#### func (String) Split

```go
func (s String) Split(sep string) []string
```

#### func (String) String

```go
func (s String) String() string
```
