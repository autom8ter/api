package driver

import (
	"net/url"
	"strings"
)

type String string

func (s String) Category() string {
	return string(s)
}

func AsIdentifier(str string) Identifier {
	return String(str)
}

func AsJSON(str string) JSON {
	return String(str)
}

func AsCategorizer(str string) Categorizer {
	return String(str)
}

func (s String) Identifier() string {
	return s.String()
}

func (s String) JSONString() string {
	return s.String()
}

func (s String) String() string {
	return string(s)
}

func (s String) ParseURL() (*url.URL, error) {
	return url.Parse(s.String())
}

func (s String) Contains(str string) bool {
	return strings.Contains(s.String(), str)
}

func (s String) Bytes(str string) []byte {
	return []byte(s.String())
}

func (s String) Split(sep string) []string {
	return strings.Split(s.String(), sep)
}
