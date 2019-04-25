package api

import (
	"encoding/json"
	"github.com/Masterminds/sprig"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func FuncMap() template.FuncMap {
	m := make(map[string]interface{})
	for k, v := range sprig.GenericFuncMap() {
		m[k] = v
	}
	return m
}

func NewTemplateFromFile(filename string) (*Template, error) {
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Template{
		Text: string(bits),
		Name: filename,
	}, nil
}

func (m *Template) IsTemplate() bool {
	return strings.Contains(m.Text, "{{")
}

func (m *Template) AsTemplate() (*template.Template, error) {
	return (template.New(m.Name).Funcs(FuncMap())).Parse(m.Text)
}

func (m *Template) RenderBytes(w io.Writer, bits *Bytes) error {
	templ, err := template.New(m.Name).Funcs(FuncMap()).Parse(m.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, bits.Bits)
}

func (m *Template) Render(w io.Writer, data interface{}) error {
	templ, err := template.New(m.Name).Funcs(FuncMap()).Parse(m.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, data)
}

func (t *Template) RenderUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if t.IsTemplate() {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			obj := session.Values["user"]
			bits, err := json.Marshal(obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u := &User{}
			err = json.Unmarshal(bits, u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.RenderBytes(w, AsBytes(u))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		io.WriteString(w, t.Text)
		return
	}
}

func RenderFile(name string, data []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bits, err := ioutil.ReadFile(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bitstring := string(bits)
		if strings.Contains(bitstring, "{{") {
			templ, err := template.New("").Funcs(FuncMap()).Parse(string(bits))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = templ.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		_, err = io.WriteString(w, bitstring)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

}

func WriteFile(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bits, err := ioutil.ReadFile(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bitstring := string(bits)
		_, err = io.WriteString(w, bitstring)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

}
