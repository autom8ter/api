# authorizer
--
    import "github.com/autom8ter/api/authorizer"


## Usage

#### type Authorizer

```go
type Authorizer struct {
	Providers map[string]oauth2.TokenSource
}
```


#### func (*Authorizer) DeepEqual

```go
func (s *Authorizer) DeepEqual(y interface{}) bool
```

#### func (*Authorizer) Do

```go
func (s *Authorizer) Do(ctx context.Context, key string, req *http.Request, object ToUnmarshal) error
```

#### func (*Authorizer) Exists

```go
func (s *Authorizer) Exists(key string) bool
```

#### func (*Authorizer) GetClient

```go
func (s *Authorizer) GetClient(ctx context.Context, key string) *http.Client
```

#### func (*Authorizer) Length

```go
func (s *Authorizer) Length() int
```

#### func (*Authorizer) Put

```go
func (s *Authorizer) Put(key string, tok oauth2.TokenSource)
```

#### func (*Authorizer) Validate

```go
func (s *Authorizer) Validate(fn func(a *Authorizer) error) error
```

#### type ToUnmarshal

```go
type ToUnmarshal interface {
	UnmarshalJSONFrom(b []byte) error
}
```
