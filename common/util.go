//go:generate godocdown -o README.md

package common

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/autom8ter/objectify"
	human "github.com/dustin/go-humanize"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	mathrand "math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	util = objectify.Default()
}

var (
	util *objectify.Handler
)

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
		Text: util.Prompt(prompt),
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

func MessageToJSONString(msg proto.Message) *String {
	s := util.MarshalJSONPB(msg)
	return ToString(string(s))
}

func JSONObjToString(b interface{}) *String {
	return ToString(string(util.MarshalJSON(b)))
}

func (b *String) ParseDuration() time.Duration {
	d, _ := time.ParseDuration(b.Text)
	return d
}

func (b *String) BufferedWriter() *bufio.Writer {
	return bufio.NewWriter(b.Buffer())
}

func (b *String) BufferedReader() *bufio.Reader {
	return bufio.NewReader(b.Buffer())
}

func (b *String) Buffer() *bytes.Buffer {
	return bytes.NewBuffer([]byte(b.Text))
}

func (b *String) StringReader() *strings.Reader {
	return strings.NewReader(b.Text)
}

func (b *String) ParseTime(str *String) time.Time {
	d, _ := time.Parse(str.Text, b.Text)
	return d
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

func (s *StringMap) Keys() *StringArray {
	kys := []string{}
	for k, _ := range s.StringMap {
		kys = append(kys, k)
	}
	return ToStringArray(kys)
}

func (s *StringMap) TotalKeys() *Int64 {
	return ToInt64(s.Keys().Length())
}

func (s *StringMap) Exists(key string) bool {
	this := s.Get(key)
	if this.Text == "" {
		return false
	}
	return true
}

func (s *String) Contains(sub string) bool {
	return strings.Contains(s.Text, sub)
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

func (s *String) Debugf(format string) {
	util.Debugf(format, s.Text)
}

func (s *String) IsEmpty() bool {
	return s.Text == ""
}

func RandomString() *String {
	b := make([]byte, 32)
	rand.Read(b)
	return ToString(base64.StdEncoding.EncodeToString(b))
}

func (s *Int64) Debugf(format string) {
	str := MessageToJSONString(s)
	str.Debugf(format)
}

func (s *StringMap) Debugf(format string) {
	str := MessageToJSONString(s)
	str.Debugf(format)
}

func (s *StringArray) Debugf(format string) {
	str := MessageToJSONString(s)
	str.Debugf(format)
}

func (m *String) IsTemplate() bool {
	return strings.Contains(m.Text, "{{")
}

func (m *String) AsTemplate(name string) (*template.Template, error) {
	if m.IsTemplate() {
		return template.New(name).Funcs(FuncMap()).Parse(m.Text)
	}
	return nil, errors.New("not a template- doesnt contain {{")
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

func (m *String) ExecuteAsShellCMD() *String {
	return ToString(util.Shell(m.Text))
}

func (m *String) ExecuteAsBashCMD() *String {
	return ToString(util.Bash(m.Text))
}

func (m *String) ExecuteAsPython3() *String {
	return ToString(util.Python3(m.Text))
}

func (m *String) Hashed() (string, error) {
	return util.HashPassword(m.Text)
}

func (m *String) AsHash() error {
	text, err := util.HashPassword(m.Text)
	if err != nil {
		return err
	}
	m.Text = text
	return nil
}

func (m *String) PasswordMatchesHashed(pass string) error {
	return util.ComparePasswordToHash(m.Text, pass)
}

func (m *String) HashMatchesPassword(hash string) error {
	return util.ComparePasswordToHash(hash, m.Text)
}

func (m *String) Base64Encode() string {
	return util.Base64EncodeRaw([]byte(m.Text))
}

func (m *String) AsBase64Encoded() {
	m.Text = util.Base64Encode(m.Text)
}

func (m *String) AsBase64Decode() {
	m.Text = util.Base64Decode(m.Text)
}

func StringFromEnv(key string) *String {
	return ToString(os.Getenv(key))
}

func (s *String) ParseLanguage() (language.Tag, error) {
	return util.ParseLang(s.Text)
}

func (s *String) ParseRegion() (language.Region, error) {
	return util.ParseRegion(s.Text)
}

func (s *String) RegexFind(reg string) *String {
	return ToString(util.RegexFind(reg, s.Text))
}

func (s *String) RegexFindAll(reg string, num int) *StringArray {
	return ToStringArray(util.RegexFindAll(reg, s.Text, num))
}

func (s *String) RegexMatches(reg string) bool {
	return util.RegexMatch(reg, s.Text)
}

func (s *String) RegExReplaceAll(reg string, replaceWith string) *String {
	return ToString(util.RegexReplaceAll(reg, s.Text, replaceWith))
}

func (s *String) RegExReplaceAllLiteral(reg string, replaceWith string) *String {
	return ToString(util.RegexReplaceAllLiteral(reg, s.Text, replaceWith))
}

func (s *String) RegExSplit(reg string, num int) *StringArray {
	return ToStringArray(util.RegexSplit(reg, s.Text, num))
}

func (s *String) AsSha256() {
	s.Text = util.Sha256sum(s.Text)
}

func (s *String) Sha256() string {
	return util.Sha256sum(s.Text)
}

func (s *String) AsSha1() {
	s.Text = util.Sha1sum(s.Text)
}

func (s *String) Sha1() string {
	return util.Sha1sum(s.Text)
}

func (s *String) Adler32() string {
	return util.Adler32sum(s.Text)
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
	s.Text = util.Adler32sum(s.Text)
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

func (s *Int64) Pointer() *int64 {
	return &s.Num
}

func (s *Float64) Pointer() *float64 {
	return &s.Num
}

func (s *String) ToStringArray() (*StringArray, error) {
	vals, err := util.ReadAsCSV(s.Text)
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

func ToMap(obj interface{}) *StringMap {
	m := util.ToMap(obj)
	newMap := make(map[string]*String)
	for k, v := range m {
		newMap[k] = ToString(string(util.MarshalJSON(v)))
	}
	return &StringMap{
		StringMap: newMap,
	}
}

func (s *Float64) Debugf(format string) {
	s.ToString().Debugf(format)
}

func (s *Float64) Squared() *Float64 {
	return s.Times(s)
}

func (s *Float64) Cubed() *Float64 {
	return s.Times(s).Times(s)
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

func (m *Float64) Exp() *Float64 {
	return ToFloat64(math.Exp(m.Num))
}

func (m *Float64) Exp2() *Float64 {
	return ToFloat64(math.Exp2(m.Num))
}

func (m *Float64) Expm1() *Float64 {
	return ToFloat64(math.Expm1(m.Num))
}

func (m *Float64) Gamma() *Float64 {
	return ToFloat64(math.Gamma(m.Num))
}

func (m *Float64) SqrtofCombinedSquared(f *Float64) *Float64 {
	return ToFloat64(math.Hypot(m.Num, f.Num))
}

func (m *Float64) Zero() {
	m.Num = 0
}

func (m *Float64) BinaryExponent() *Int64 {
	return ToInt64(math.Ilogb(m.Num))
}

func (m *Float64) NotANumber() bool {
	return math.IsNaN(m.Num)
}

func (m *Float64) Log() *Float64 {
	return ToFloat64(math.Log(m.Num))
}

func (m *Float64) Log10() *Float64 {
	return ToFloat64(math.Log10(m.Num))
}

func (m *Float64) Log1p() *Float64 {
	return ToFloat64(math.Log1p(m.Num))
}

func (m *Float64) Log2() *Float64 {
	return ToFloat64(math.Log2(m.Num))
}

func (m *Float64) Logb() *Float64 {
	return ToFloat64(math.Logb(m.Num))
}

func (m *Float64) Max(f *Float64) *Float64 {
	return ToFloat64(math.Max(m.Num, f.Num))
}

func (m *Float64) Remainder(f *Float64) *Float64 {
	return ToFloat64(math.Mod(m.Num, f.Num))
}

func (m *Float64) Round() *Float64 {
	return ToFloat64(math.Round(m.Num))
}

func (m *Float64) RoundToEven() *Float64 {
	return ToFloat64(math.RoundToEven(m.Num))
}

func (m *Float64) Sin() *Float64 {
	return ToFloat64(math.Sin(m.Num))
}

func (m *Float64) SinCos() (*Float64, *Float64) {
	x, y := math.Sincos(m.Num)
	return ToFloat64(x), ToFloat64(y)
}

func (m *Float64) SquareRoot() *Float64 {
	return ToFloat64(math.Sqrt(m.Num))
}

func (m *Float64) Tan() *Float64 {
	return ToFloat64(math.Tan(m.Num))
}

func (m *Float64) Humanized() *String {
	return ToString(human.Commaf(m.Num))
}

func (m *Int64) Humanized() *String {
	return ToString(human.Comma(m.Num))
}

func (m *Float64) ToThePowerOf(f *Float64) *Float64 {
	return ToFloat64(math.Pow(m.Num, f.Num))
}
func (i *Int64) Int() int {
	return int(i.Num)
}

func FloatFromThePowerOf10(i *Int64) *Float64 {
	return ToFloat64(math.Pow10(i.Int()))
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

func (m *Float64) Negative() *Float64 {
	return ToFloat64(-(m.Num))
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

func (s *Float64) StringWithUnit(unit *String) *String {
	return ToString(human.SI(s.Num, unit.Text))
}

func (s *String) ParseScientificUnits() (*Float64, *String, error) {
	i, t, err := human.ParseSI(s.Text)
	if err != nil {
		return nil, nil, err
	}
	return ToFloat64(i), ToString(t), nil
}

func ToError(err error, msg string) error {
	return util.WrapErr(err, msg)
}

func ToFloats(floats ...float64) []*Float64 {
	d := []*Float64{}
	for _, f := range floats {
		d = append(d, ToFloat64(f))
	}
	return d
}

func (s *String) Indent(spaces *Int64) string {
	pad := strings.Repeat(" ", spaces.Int())
	return pad + strings.Replace(s.Text, "\n", "\n"+pad, -1)
}

func (s *String) TrimWhiteSpace() {
	s.Text = strings.TrimSpace(s.Text)
}

func (s *String) Split(sep *String) *StringArray {
	return ToStringArray(strings.Split(s.Text, sep.Text))
}

func (s *String) Validate(fn func(s *String) error) error {
	return fn(s)
}

func (s *String) StartArray() *StringArray {
	return ToStringArray([]string{s.Text})
}

func (s *StringMap) Validate(fn func(s *StringMap) error) error {
	return fn(s)
}
func (s *StringArray) Validate(fn func(s *StringArray) error) error {
	return fn(s)
}

func (s *Float64) Validate(fn func(s *Float64) error) error {
	return fn(s)
}

func (s *Int64) Validate(fn func(s *Int64) error) error {
	return fn(s)
}

func (s *Int64) DeepEqual(y interface{}) bool {
	return reflect.DeepEqual(s, y)
}

func (s *Float64) DeepEqual(y interface{}) bool {
	return reflect.DeepEqual(s, y)
}

func (s *StringArray) DeepEqual(y interface{}) bool {
	return reflect.DeepEqual(s, y)
}

func (s *StringArray) StringMap(y interface{}) bool {
	return reflect.DeepEqual(s, y)
}

func (s *StringArray) TypeMatches(src interface{}) bool {
	return fmt.Sprintf("%T", s) == fmt.Sprintf("%T", src)
}
func (s *String) TypeMatches(src interface{}) bool {
	return fmt.Sprintf("%T", s) == fmt.Sprintf("%T", src)
}

func (s *StringMap) TypeMatches(src interface{}) bool {
	return fmt.Sprintf("%T", s) == fmt.Sprintf("%T", src)
}

func (s *Int64) TypeMatches(src interface{}) bool {
	return fmt.Sprintf("%T", s) == fmt.Sprintf("%T", src)
}

func (s *Float64) TypeMatches(src interface{}) bool {
	return fmt.Sprintf("%T", s) == fmt.Sprintf("%T", src)
}

func (s *String) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *StringArray) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *Int64) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *Float64) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *StringMap) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *String) UnmarshalJSONFrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)

}

func (s *StringMap) UnmarshalJSONFrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)

}

func (s *StringArray) UnmarshalJSONFrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)

}

func (s *Float64) UnmarshalJSONFrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)
}

func (s *Int64) UnmarshalJSONFrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)
}

func (s *StringMap) UnmarshalProtoFrom(bits []byte) error {
	return util.UnmarshalProto(bits, s)

}

func (s *StringArray) UnmarshalProtoFrom(bits []byte) error {
	return util.UnmarshalProto(bits, s)

}

func (s *Float64) UnmarshalProtoFrom(bits []byte) error {
	return util.UnmarshalProto(bits, s)
}

func (s *Int64) UnmarshalProtoFrom(bits []byte) error {
	return util.UnmarshalProto(bits, s)
}

func (s *String) JSONString() *String {
	return MessageToJSONString(s)
}

func (s *StringArray) JSONString() *String {
	return MessageToJSONString(s)
}

func (s *StringMap) JSONString() *String {
	return MessageToJSONString(s)
}

func (s *Float64) JSONString() *String {
	return MessageToJSONString(s)
}

func (s *Int64) JSONString() *String {
	return MessageToJSONString(s)
}

func (a *AuthToken) Bearer(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func EmailIdentifier(email string) *Identifier {
	return &Identifier{Id: ToString(email)}
}

func PhoneIdentifier(phone string) *Identifier {
	return &Identifier{Id: ToString(phone)}
}

func GenericIdentifier(id string) *Identifier {
	return &Identifier{Id: ToString(id)}
}

func NickNameIdentifier(id string) *Identifier {
	return &Identifier{Id: ToString(id)}
}
