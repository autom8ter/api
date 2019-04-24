package api

import (
	"encoding/json"
	"github.com/Masterminds/sprig"
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

func RenderFileWithUserInfo(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bits, err := ioutil.ReadFile(filename)
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
			obj := session.Values["userinfo"]
			bits, err := json.Marshal(obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u := &UserInfo{}
			err = json.Unmarshal(bits, u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = templ.Execute(w, u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		io.WriteString(w, bitstring)
		return
	}
}

func RenderFileWithData(name string, data interface{}) http.HandlerFunc {
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
