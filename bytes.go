package api

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewBytes() *Bytes {
	return &Bytes{
		Bits: []byte{},
	}
}

func BytesFromString(str string) *Bytes {
	return &Bytes{
		Bits: []byte(str),
	}
}

func BytesFromBytes(bits []byte) *Bytes {
	return &Bytes{
		Bits: bits,
	}
}

func BytesFromReader(r io.Reader) (*Bytes, error) {
	b := NewBytes()
	_, err := io.Copy(b, r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func BytesFromFile(fileName string) (*Bytes, error) {
	bits, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Bytes{
		Bits: bits,
	}, nil
}

func (m *Bytes) Read(p []byte) (n int, err error) {
	before := len(m.Bits)
	m.Bits = append(m.Bits, p...)
	after := len(m.Bits)
	return after - before, nil
}

func (m *Bytes) Write(p []byte) (n int, err error) {
	before := len(m.Bits)
	m.Bits = append(m.Bits, p...)
	after := len(m.Bits)
	return after - before, nil
}

func (m *Bytes) YAML() []byte {
	return Util.MarshalYAML(m.Bits)
}

func (m *Bytes) JSON() []byte {
	return Util.MarshalJSON(m.Bits)
}

func (m *Bytes) XML() []byte {
	return Util.MarshalXML(m.Bits)
}

func (m *Bytes) Length() int {
	return len(m.Bits)
}

func (m *Bytes) Compile(w io.Writer, t *Template) error {
	return t.Render(w, m)
}

func (m *Bytes) Contains(str string) bool {
	return strings.Contains(string(m.Bits), str)
}

func (m *Bytes) BitString() string {
	return string(m.Bits)
}

func (m *Bytes) CompileHTTP(t *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.Compile(w, t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (m *Bytes) WriteHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, m.BitString())
		return
	}
}

func (m *Bytes) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(m.JSON(), obj)
}

func (m *Bytes) UnMarshalProto(obj interface{}) error {
	return proto.Unmarshal(m.Bits, obj.(proto.Message))
}

func AsBytes(obj interface{}) *Bytes {
	return &Bytes{
		Bits: Util.MarshalJSON(obj),
	}
}
