//go:generate godocdown -o README.md

package driver

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
