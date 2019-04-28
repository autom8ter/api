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
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/text/language"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	mathrand "math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", ToStringArray(os.Environ()))
	Util = objectify.Default()
	httpClient = http.DefaultClient
	store = sessions.NewCookieStore([]byte(os.Getenv(SESSION_SECRET_ENV_KEY)))
	gob.Register(map[string]interface{}{})
}

var (
	Util                   *objectify.Handler
	Context                context.Context
	httpClient             *http.Client
	store                  *sessions.CookieStore
	AUTH_SESSION           = "auth-session"
	SESSION_SECRET_ENV_KEY = "SECRET"
)

func GetAuthSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, AUTH_SESSION)
}

func GetStateSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "state")
}

func ENVFromContext(ctx context.Context) *StringArray {
	return ctx.Value("env").(*StringArray)
}

func StringArrayFromEnv() *StringArray {
	return ToStringArray(os.Environ())
}

func StringFromFile(filename string) (*String, error) {
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ToString(string(bits)), nil
}

func ToString(s string) *String {
	return &String{
		Text: s,
	}
}

func StringFromPrompt(prompt string) *String {
	return &String{
		Text: Util.Prompt(prompt),
	}
}

func ToBool(b bool) *Bool {
	return &Bool{
		Answer: b,
	}
}

func ToStringArray(s []string) *StringArray {
	news := []*String{}
	for _, v := range s {
		news = append(news, ToString(v))
	}
	return &StringArray{
		Strings: news,
	}
}

func ToStringMap(s map[string]string) *StringMap {
	news := map[string]*String{}
	for k, v := range s {
		news[k] = ToString(v)
	}
	return &StringMap{
		StringMap: news,
	}
}

func ToBytes(b []byte) *Bytes {
	return &Bytes{
		Bits: b,
	}
}

func ToInt64(i int) *Int64 {
	return &Int64{
		Num: int64(i),
	}
}

func ToFloat64(i float64) *Float64 {
	return &Float64{
		Num: i,
	}
}

func ObjToBytes(b interface{}) *Bytes {
	return ToBytes(Util.MarshalJSON(b))
}

func ObjToString(b interface{}) *String {
	return ToString(string(Util.MarshalJSON(b)))
}

func (b *Bytes) ToString() *String {
	return ToString(string(b.Bits))
}

func (b *String) ToBytes() *Bytes {
	return ToBytes([]byte(b.Text))
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
			r.Form.Set(k, v.Text)
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

func (m *Bytes) Compile(name string, w io.Writer, t *String) error {
	return t.Render(name, w, m)
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

func (m *Bytes) CompileHTTP(name string, t *String) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.Compile(name, w, t); err != nil {
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

func (s *StringMap) Get(key string) *String {
	return s.StringMap[key]
}

func (s *StringMap) Put(key string, val *String) {
	s.StringMap[key] = val
}

func (s *StringMap) Clear(key string) {
	s.StringMap[key] = nil
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
	if this.Text == "" {
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

func (s *String) Matches(this string) bool {
	return s.Text == this
}

func (s *String) Println() {
	fmt.Println(s.Text)
}

func (s *String) Debugln() {
	Util.Entry().Debugln(s.Text)
}

func (s *String) IsEmpty() bool {
	return s.Text == ""
}

func RandomString() *String {
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

func (s *Token) Debugln() {
	ObjToString(s).Debugln()
}

func (t *Token) JSON() []byte {
	return Util.MarshalJSON(t)
}

func (t *Token) YAML() []byte {
	return Util.MarshalYAML(t)
}

func (t *Token) XML() []byte {
	return Util.MarshalXML(t)
}

func (s *StringMap) Debugln() {
	ObjToString(s).Debugln()
}

func (s *Bytes) Debugln() {
	ObjToString(s).Debugln()
}

func (s *StringArray) Debugln() {
	ObjToString(s).Debugln()
}

func (t *StringMap) JSON() []byte {
	return Util.MarshalJSON(t)
}

func (t *StringMap) YAML() []byte {
	return Util.MarshalYAML(t)
}

func (t *StringMap) XML() []byte {
	return Util.MarshalXML(t)
}

func (t *StringArray) JSON() []byte {
	return Util.MarshalJSON(t)
}

func (t *StringArray) YAML() []byte {
	return Util.MarshalYAML(t)
}

func (t *StringArray) XML() []byte {
	return Util.MarshalXML(t)
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

func (m *String) IsTemplate() *Bool {
	return ToBool(strings.Contains(m.Text, "{{"))
}

func (m *String) AsTemplate(name string) (*template.Template, error) {
	if m.IsTemplate().Answer {
		return template.New(name).Funcs(FuncMap()).Parse(m.Text)
	}
	return nil, errors.New("not a template- doesnt contain {{")
}

func (m *String) RenderBytes(name string, w io.Writer, bits *Bytes) error {
	templ, err := m.AsTemplate(name)
	if err != nil {
		return err
	}
	return templ.Execute(w, bits.Bits)
}

func (m *String) Render(name string, w io.Writer, data interface{}) error {
	templ, err := m.AsTemplate(name)
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

func (m *StringArray) Append(vals ...*String) {
	for _, v := range vals {
		m.Strings = append(m.Strings, v)
	}
}

func (m *StringArray) Length() int {
	return len(m.Strings)
}

func (m *StringArray) IsEmpty() bool {
	return len(m.Strings) == 0
}

func (m *StringArray) SelectRandom() *String {
	mathrand.Seed(time.Now().Unix())
	return m.Strings[mathrand.Intn(len(m.Strings))]
}

func (m *StringArray) Get(index int) *String {
	return m.Strings[index]
}

func (m *StringArray) First() *String {
	return m.Strings[0]
}

func (m *StringArray) Last() *String {
	return m.Strings[m.Length()-1]
}

func (m *String) WriteString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, m.Text)
		return
	}
}

func (m *String) JSON() []byte {
	return Util.MarshalJSON(m.Text)
}

func (m *String) YAML() []byte {
	return Util.MarshalYAML(m.Text)
}

func (m *String) XML() []byte {
	return Util.MarshalXML(m.Text)
}

func (m *String) ExecuteAsShellCMD() *Bytes {
	return ToBytes([]byte(Util.Shell(m.Text)))
}

func (m *String) ExecuteAsBashCMD() *Bytes {
	return ToBytes([]byte(Util.Bash(m.Text)))
}

func (m *String) ExecuteAsPython3() *Bytes {
	return ToBytes([]byte(Util.Python3(m.Text)))
}

func (m *String) Hashed() (string, error) {
	return Util.HashPassword(m.Text)
}

func (m *String) AsHash() error {
	text, err := Util.HashPassword(m.Text)
	if err != nil {
		return err
	}
	m.Text = text
	return nil
}

func (m *String) PasswordMatchesHashed(pass string) error {
	return Util.ComparePasswordToHash(m.Text, pass)
}

func (m *String) HashMatchesPassword(hash string) error {
	return Util.ComparePasswordToHash(hash, m.Text)
}

func (m *String) Base64Encode() string {
	return Util.Base64EncodeRaw([]byte(m.Text))
}

func (m *String) AsBase64Encoded() {
	m.Text = Util.Base64Encode(m.Text)
}

func (m *String) AsBase64Decode() {
	m.Text = Util.Base64Decode(m.Text)
}

func StringFromEnv(key string) *String {
	return ToString(os.Getenv(key))
}

func (s *String) ParseLanguage() (language.Tag, error) {
	return Util.ParseLang(s.Text)
}

func (s *String) ParseRegion() (language.Region, error) {
	return Util.ParseRegion(s.Text)
}

func (s *String) RegexFind(reg string) *String {
	return ToString(Util.RegexFind(reg, s.Text))
}

func (s *String) RegexFindAll(reg string, num int) *StringArray {
	return ToStringArray(Util.RegexFindAll(reg, s.Text, num))
}

func (s *String) RegexMatches(reg string) *Bool {
	return ToBool(Util.RegexMatch(reg, s.Text))
}

func (s *String) RegExReplaceAll(reg string, replaceWith string) *String {
	return ToString(Util.RegexReplaceAll(reg, s.Text, replaceWith))
}

func (s *String) RegExReplaceAllLiteral(reg string, replaceWith string) *String {
	return ToString(Util.RegexReplaceAllLiteral(reg, s.Text, replaceWith))
}

func (s *String) RegExSplit(reg string, num int) *StringArray {
	return ToStringArray(Util.RegexSplit(reg, s.Text, num))
}

func (s *String) AsSha256() {
	s.Text = Util.Sha256sum(s.Text)
}

func (s *String) Sha256() string {
	return Util.Sha256sum(s.Text)
}

func (s *String) AsSha1() {
	s.Text = Util.Sha1sum(s.Text)
}

func (s *String) Sha1() string {
	return Util.Sha1sum(s.Text)
}

func (s *String) Adler32() string {
	return Util.Adler32sum(s.Text)
}

func (s *String) ParseURL() (*url.URL, error) {
	return url.Parse(s.Text)
}

func (s *String) AddLine(text string) {
	s.Text = fmt.Sprintln(s.Text) + text
}

func (s *String) Append(text string) {
	s.Text = s.Text + text
}

func (s *String) PrePend(text string) {
	s.Text = text + s.Text
}

func (s *String) SetEnv(key string) error {
	return os.Setenv(key, s.Text)
}

func (s *String) AsAdler32() {
	s.Text = Util.Adler32sum(s.Text)
}

func (s *String) ToInt64() (*Int64, error) {
	i, err := strconv.Atoi(s.Text)
	if err != nil {
		return nil, err
	}
	return ToInt64(i), nil
}

func (s *String) Pointer() *string {
	return &s.Text
}

func (s *Bool) Pointer() *bool {
	return &s.Answer
}

func (s *Int64) Pointer() *int64 {
	return &s.Num
}

func (s *Float64) Pointer() *float64 {
	return &s.Num
}

func (s *String) ToStringArray() (*StringArray, error) {
	vals, err := Util.ReadAsCSV(s.Text)
	if err != nil {
		return nil, err
	}
	return ToStringArray(vals), nil
}

func (s *StringArray) Pointer() []*string {
	out := make([]*string, len(s.Strings))
	for i := range s.Strings {
		out[i] = &s.Strings[i].Text
	}
	return out
}

func (s *StringArray) ToString() *String {
	return ToString(string(Util.MarshalJSON(s.Strings)))
}

func StructToMap(obj interface{}) *StringMap {
	m := Util.ToMap(obj)
	newMap := make(map[string]*String)
	for k, v := range m {
		newMap[k] = ToString(string(Util.MarshalJSON(v)))
	}
	return &StringMap{
		StringMap: newMap,
	}
}

func (s *StringMap) ToString() *String {
	return ToString(string(Util.MarshalJSON(s.StringMap)))
}

func (s *Float64) Debugln() {
	s.ToString().Debugln()
}

func (s *Int64) Debugln() {
	s.ToString().Debugln()
}

func (s *Int64) ToString() *String {
	return ToString(strconv.Itoa(int(s.Num)))
}

func (s *Float64) ToString() *String {
	return ToString(fmt.Sprintf("%v", s.Num))
}

func (s *StringArray) Array() []string {
	out := make([]string, len(s.Strings))
	for i := range s.Strings {
		out[i] = s.Strings[i].Text
	}
	return out
}

//Abs returns the absolute value of Float64.
func (m *Float64) Abs() *Float64 {
	return ToFloat64(math.Abs(m.Num))
}

//CubeRoot returns the cube root of x.
func (m *Float64) CubeRoot() *Float64 {
	return ToFloat64(math.Cbrt(m.Num))
}

//Ceil returns the least integer value greater than or equal to x.
func (m *Float64) Ceiling() *Float64 {
	return ToFloat64(math.Ceil(m.Num))
}

func (m *Float64) Minus(n *Float64) *Float64 {
	return ToFloat64(m.Num - n.Num)
}

func (m *Float64) Plus(n *Float64) *Float64 {
	return ToFloat64(m.Num + n.Num)
}

func (m *Float64) Times(n *Float64) *Float64 {
	return ToFloat64(m.Num * n.Num)
}

func (m *Float64) DividedBy(n *Float64) *Float64 {
	return ToFloat64(m.Num / n.Num)
}

func (m *Int64) Minus(n *Int64) *Int64 {
	return ToInt64(int(m.Num - n.Num))
}

func (m *Int64) Plus(n *Int64) *Int64 {
	return ToInt64(int(m.Num + n.Num))
}

func (m *Int64) Times(n *Int64) *Int64 {
	return ToInt64(int(m.Num * n.Num))
}

func (m *Int64) DividedBy(n *Int64) *Int64 {
	return ToInt64(int(m.Num / n.Num))
}

func (m *Int64) Remainder(n *Int64) *Int64 {
	return ToInt64(int(m.Num % n.Num))
}
