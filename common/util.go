//go:generate godocdown -o README.md

package common

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/autom8ter/objectify"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", ToStringArray(os.Environ()))
	Util = objectify.Default()
	httpClient = http.DefaultClient

}

var (
	Util         *objectify.Handler
	Context      context.Context
	httpClient   *http.Client
	store        *sessions.CookieStore
	AUTH_SESSION = "auth-session"
)

func ENVFromContext(ctx context.Context) *StringArray {
	return ctx.Value("env").(*StringArray)
}

func ToString(s string) *String {
	return &String{
		Text: s,
	}
}

func ToBool(b bool) *Bool {
	return &Bool{
		Answer: b,
	}
}

func ToStringArray(s []string) *StringArray {
	return &StringArray{
		Strings: s,
	}
}

func ToStringMap(s map[string]string) *StringMap {
	return &StringMap{
		StringMap: s,
	}
}

func ToBytes(b []byte) *Bytes {
	return &Bytes{
		Bits: b,
	}
}

func (h HTTPMethod) Normalize() *String {
	switch h {
	case HTTPMethod_POST:
		return ToString("POST")
	case HTTPMethod_PATCH:
		return ToString("PATCH")
	default:
		return ToString("GET")
	}
}

func (h *HTTPRequest) Do(token *Token) (*Bytes, error) {
	u, err := url.Parse(h.Url.Text)
	if err != nil {
		return nil, err
	}
	r := &http.Request{
		Method: h.Method.Normalize().Text,
		URL:    u,
	}
	if token.TokenType.Text == "" {
		token.TokenType.Text = "Bearer"
	}
	if token.AccessToken.Text != "" {
		r.Header.Set("Authorization", token.TokenType.Text+" "+token.AccessToken.Text)
	}

	if len(h.Form.StringMap) != 0 {
		for k, v := range h.Form.StringMap {
			r.Form.Set(k, v)
		}
	}
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, Util.WrapErr(err, resp.Status)
	}
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Util.WrapErr(err, "reading http response body")
	}
	return &Bytes{
		Bits: bits,
	}, nil

}

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

func (b *Bytes) UnmarshalProto(obj interface{}) error {
	return proto.Unmarshal(b.Bits, obj.(proto.Message))
}

func (b *Bytes) UnmarshalJSON(obj interface{}) error {
	return json.Unmarshal(b.Bits, obj)
}

func (m *Bytes) CompileHTTP(t *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.Compile(w, t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (m *Bytes) WriteString() http.HandlerFunc {
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

func (m *Bytes) Clear() {
	m.Bits = []byte{}
}

func AsBytes(obj interface{}) *Bytes {
	return &Bytes{
		Bits: Util.MarshalJSON(obj),
	}
}

func (s *StringMap) Get(key string) string {
	return s.StringMap[key]
}

func (s *StringMap) Put(key string, val string) {
	s.StringMap[key] = val
}

func (s *StringMap) Clear(key string) {
	s.StringMap[key] = ""
}

func (s *StringMap) Keys() []string {
	kys := []string{}
	for k, _ := range s.StringMap {
		kys = append(kys, k)
	}
	return kys
}

func (s *StringMap) TotalKeys() int {
	return len(s.Keys())
}

func (s *StringMap) Exists(key string) bool {
	this := s.Get(key)
	if this == "" {
		return false
	}
	return true
}

func (s *String) Contains(sub string) *Bool {
	return ToBool(strings.Contains(s.Text, sub))
}

func (s *String) TrimPrefix(sub string) {
	s.Text = strings.TrimPrefix(s.Text, sub)
}

func (s *String) TrimSuffix(sub string) {
	s.Text = strings.TrimSuffix(s.Text, sub)
}

func (s *String) Index(sub string) *Int64 {
	return &Int64{Num: int64(strings.Index(s.Text, sub))}
}

func (s *String) ToLower() {
	s.Text = strings.ToLower(s.Text)
}

func (s *String) ToUpper() {
	s.Text = strings.ToUpper(s.Text)
}

func (s *String) ToTitle() {
	s.Text = strings.ToTitle(s.Text)
}

func (s *String) Replace(oldNew ...string) {
	r := strings.NewReplacer(oldNew...)
	s.Text = r.Replace(s.Text)
}

func (s *String) Println() {
	fmt.Println(s.Text)
}

func (s *String) IsEmpty() bool {
	if s.Text == "" {
		return true
	}
	return false
}

func SecretFromEnv() *Secret {
	return &Secret{
		Text: os.Getenv("SECRET"),
	}
}

func (s *Secret) InitSessions() error {
	store = sessions.NewCookieStore([]byte(s.Text))
	gob.Register(map[string]interface{}{})
	return nil
}

func GetAuthSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, AUTH_SESSION)
}

func GetStateSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "state")
}

func Random() *String {
	b := make([]byte, 32)
	rand.Read(b)
	return ToString(base64.StdEncoding.EncodeToString(b))
}

func FromSession(key string, session *sessions.Session) interface{} {
	return session.Values[key]
}

func AuthSessionValues(session *sessions.Session, key string, data map[string]interface{}) {
	session.Values[key] = data
}

func (t *Token) ToSession(session *sessions.Session) {
	session.Values["access_token"] = t.AccessToken
	session.Values["token_type"] = t.TokenType
	session.Values["refresh_token"] = t.RefreshToken
	session.Values["expiry"] = t.Expiry
	session.Values["id_token"] = t.IdToken

}
func TokenFromAuthSession(session *sessions.Session) (*Token, error) {
	return &Token{
		AccessToken:  ToString(session.Values["access_token"].(string)),
		TokenType:    ToString(session.Values["token_type"].(string)),
		RefreshToken: ToString(session.Values["refresh_token"].(string)),
		Expiry:       ToString(session.Values["expiry"].(string)),
		IdToken:      ToString(session.Values["id_token"].(string)),
	}, nil

}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	if err := sessions.Save(r, w); err != nil {
		http.Error(w, Util.WrapErr(err, "saving session").Error(), http.StatusInternalServerError)
	}
}

func TokenFromOAuthToken(tok *oauth2.Token) *Token {
	return &Token{
		AccessToken:  ToString(tok.AccessToken),
		TokenType:    ToString(tok.TokenType),
		RefreshToken: ToString(tok.RefreshToken),
		Expiry:       ToString(tok.Expiry.String()),
		IdToken:      ToString(tok.Extra("id_token").(string)),
	}
}

func NewTemplateFromFile(filename string) (*Template, error) {
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Template{
		Text: ToString(string(bits)),
		Name: ToString(filename),
	}, nil
}

func (m *Template) IsTemplate() *Bool {
	return ToBool(strings.Contains(m.Text.Text, "{{"))
}

func (m *Template) AsTemplate() (*template.Template, error) {
	return (template.New(m.Name.Text).Funcs(FuncMap())).Parse(m.Text.Text)
}

func (m *Template) RenderBytes(w io.Writer, bits *Bytes) error {
	templ, err := template.New(m.Name.Text).Funcs(FuncMap()).Parse(m.Text.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, bits.Bits)
}

func (m *Template) Render(w io.Writer, data interface{}) error {
	templ, err := template.New(m.Name.Text).Funcs(FuncMap()).Parse(m.Text.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, data)
}

func FuncMap() template.FuncMap {
	m := make(map[string]interface{})
	for k, v := range sprig.GenericFuncMap() {
		m[k] = v
	}
	return m
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

func (m *StringArray) Append(vals ...*String) {
	for _, v := range vals {
		m.Strings = append(m.Strings, v.Text)
	}
}
