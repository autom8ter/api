package api

import (
	"context"
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/objectify"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"os"
	"reflect"
)

type Messenger interface {
	Attributes() map[string]string
	DataBytes() []byte
	ID() string
}

func AsMessage(m Messenger) *pubsub.PubsubMessage {
	return &pubsub.PubsubMessage{
		Data:       m.DataBytes(),
		Attributes: m.Attributes(),
		MessageId:  m.ID(),
	}
}

func DataTo(msg *pubsub.PubsubMessage, obj interface{}) error {
	return json.Unmarshal(msg.Data, obj)
}

var Util = objectify.Default()

type ClientSet struct {
	UserSet UserServiceClient
}

func NewClientSet(ctx context.Context, addr string, opts ...grpc.DialOption) (*ClientSet, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientSet{
		UserSet: NewUserServiceClient(conn),
	}, nil
}

type UserServer struct {
	UserServiceServer
	driver.PluginFunc
}

func NewUserServer(addr string, server UserServiceServer) *UserServer {
	s := &UserServer{
		UserServiceServer: server,
	}
	s.PluginFunc = func(server *grpc.Server) {
		RegisterUserServiceServer(server, s)
	}
	return s
}

func Serve(addr string, debug bool, plugin driver.Plugin) error {
	return engine.Serve(addr, debug, plugin)
}

//////////////////////////////////////////////////////

func AccessFromJSON(j *JSON) *Access {
	a := &Access{}
	_ = json.Unmarshal(j.Data, a)
	return a
}

//TWILIO_ACCOUNT TWILIO_KEY SENDGRID_ACCOUNT SENDGRID_KEY STRIPE_ACCOUNT STRIPE_KEY SLACK_ACCOUNT SLACK_KEY GCP_PROJECT GCP_KEY
func AccessFromEnv() *Access {
	return &Access{
		Autom8TerAccount: os.Getenv("AUTOM8TER_ACCOUNT"),
		Autom8TerKey:     os.Getenv("AUTOM8TER_KEY"),
		TwilioAccount:    os.Getenv("TWILIO_ACCOUNT"),
		TwilioKey:        os.Getenv("TWILIO_KEY"),
		SendgridAccount:  os.Getenv("SENDGRID_ACCOUNT"),
		SendgridKey:      os.Getenv("SENDGRID_KEY"),
		StripeAccount:    os.Getenv("STRIPE_ACCOUNT"),
		StripeKey:        os.Getenv("STRIPE_KEY"),
		SlackAccount:     os.Getenv("SLACK_ACCOUNT"),
		SlackKey:         os.Getenv("SLACK_KEY"),
		GcpProject:       os.Getenv("GCP_PROJECT"),
		GcpKey:           os.Getenv("GCP_KEY"),
	}
}

func (a *Access) Validate() error {
	if a.TwilioKey == "" {
		return errors.New("empty twilio key")
	}
	if a.TwilioAccount == "" {
		return errors.New("empty twilio account")
	}
	if a.Autom8TerAccount == "" {
		return errors.New("empty autom8ter account")
	}
	if a.Autom8TerKey == "" {
		return errors.New("empty autom8ter key")
	}
	if a.SendgridKey == "" {
		return errors.New("empty sendgrid key")
	}
	if a.StripeKey == "" {
		return errors.New("empty stripe key")
	}
	if a.SlackKey == "" {
		return errors.New("empty slack key")
	}
	return nil
}

//////////////////////////////////////////////////////

type AttachmentFunc func(a *Attachment)

type AttachmentActionFunc func(a *AttachmentAction)

func NewAttachment(opts ...AttachmentFunc) *Attachment {
	a := &Attachment{}
	for _, o := range opts {
		o(a)
	}
	return a
}

func NewAttachmentAction(opts ...AttachmentActionFunc) *AttachmentAction {
	a := &AttachmentAction{}
	for _, o := range opts {
		o(a)
	}
	return a
}

//

func (j *Refund) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Refund) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Refund) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Refund) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Refund) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Refund) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Refund) ID() string {
	return reflect.TypeOf(s).String()
}

//

func (j *Product) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Product) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Product) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Product) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Product) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Product) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Product) ID() string {
	return reflect.TypeOf(s).String()
}

//

func (j *Charge) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Charge) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Charge) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Charge) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Charge) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Charge) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Charge) ID() string {
	return reflect.TypeOf(s).String()
}

//

func (j *BankAccount) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *BankAccount) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *BankAccount) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *BankAccount) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *BankAccount) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *BankAccount) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *BankAccount) ID() string {
	return reflect.TypeOf(s).String()
}

//
func (j *File) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *File) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *File) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *File) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *File) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *File) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *File) ID() string {
	return reflect.TypeOf(s).String()
}

//
func (j *StandardClaims) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *StandardClaims) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *StandardClaims) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *StandardClaims) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *StandardClaims) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *StandardClaims) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *StandardClaims) ID() string {
	return reflect.TypeOf(s).String()
}

//

func (j *Access) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Access) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Access) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Access) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Access) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Access) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Access) ID() string {
	return reflect.TypeOf(s).String()
}

//
func (j *Card) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Card) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Card) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Card) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *Card) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *Card) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Card) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Card) ID() string {
	return reflect.TypeOf(s).String()
}

//
func (j *Address) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Address) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Address) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Address) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *Address) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *Address) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Address) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Address) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////

func (j *Attachment) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Attachment) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Attachment) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Attachment) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Attachment) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Attachment) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Attachment) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////

func (j *SMS) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *SMS) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *SMS) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *SMS) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *SMS) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *SMS) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *SMS) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *SMS) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////

func (j *Email) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Email) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Email) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Email) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *Email) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *Email) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Email) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Email) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////

func (j *Call) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Call) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Call) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Call) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *Call) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *Call) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Call) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Call) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////

func (j *AttachmentAction) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *AttachmentAction) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *AttachmentAction) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *AttachmentAction) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (j *AttachmentAction) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

func (s *AttachmentAction) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *AttachmentAction) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *AttachmentAction) ID() string {
	return reflect.TypeOf(s).String()
}

//////////////////////////////////////////////////////
func (j *JSON) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *JSON) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(j.Data, obj)
}

func (j *JSON) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j.Data, w)
}

func (j *JSON) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j.Data, w)
}

func (s *JSON) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *JSON) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *JSON) ID() string {
	return reflect.TypeOf(s).String()
}

///////////////////////////////////////////////////////////////

func (j *Customer) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *Customer) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *Customer) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *Customer) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}

func (s *Customer) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *Customer) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *Customer) ID() string {
	return reflect.TypeOf(s).String()
}

///////////////////////////////////////////////////////////////

func (j *SignedKey) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j), nil
}

func (j *SignedKey) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j), obj)
}

func (j *SignedKey) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j, w)
}

func (j *SignedKey) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j, w)
}
func (s *SignedKey) Attributes() map[string]string {
	return Util.Attributes(s)

}

func (s *SignedKey) DataBytes() []byte {
	return Util.MarshalJSON(s)

}

func (s *SignedKey) ID() string {
	return reflect.TypeOf(s).String()
}

///////////////////////////////////////////////////////////////

type Claims struct {
	standardClaims map[string]jwt.StandardClaims
}

func (c *Claims) SignedTwilioToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["twilio"])
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
		"cty": "twilio-fpa;v=1",
	}
	for k, v := range headers {
		token.Header[k] = v
	}
	tok, err := token.SignedString([]byte(secret))
	return &SignedKey{SignedKey: tok}, err
}

func (c *Claims) SignedSendGridToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["sendgrid"])
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
		"cty": "sendgrid",
	}
	for k, v := range headers {
		token.Header[k] = v
	}
	tok, err := token.SignedString([]byte(secret))
	return &SignedKey{SignedKey: tok}, err
}

func (c *Claims) SignedSlackToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["slack"])
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
		"cty": "slack",
	}
	for k, v := range headers {
		token.Header[k] = v
	}
	tok, err := token.SignedString([]byte(secret))
	return &SignedKey{SignedKey: tok}, err
}

func (c *Claims) SignedStripeToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["stripe"])
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
		"cty": "stripe",
	}
	for k, v := range headers {
		token.Header[k] = v
	}
	tok, err := token.SignedString([]byte(secret))
	return &SignedKey{SignedKey: tok}, err
}

func (c *Claims) SignedGCPToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["gcp"])
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
		"cty": "gcp",
	}

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &SignedKey{
		SignedKey: signed,
	}, nil
}

func (t *StandardClaims) Claims() *Claims {
	claims := make(map[string]jwt.StandardClaims)
	claims["twilio"] = jwt.StandardClaims{
		Audience:  t.Audience,
		ExpiresAt: t.ExpiresAt,
		Id:        t.Id,
		IssuedAt:  t.IssuedAt,
		Issuer:    t.Access.TwilioKey,
		NotBefore: t.NotBefore,
		Subject:   t.Access.TwilioAccount,
	}
	claims["sendgrid"] = jwt.StandardClaims{
		Audience:  t.Audience,
		ExpiresAt: t.ExpiresAt,
		Id:        t.Id,
		IssuedAt:  t.IssuedAt,
		Issuer:    t.Access.SendgridKey,
		NotBefore: t.NotBefore,
		Subject:   t.Access.SendgridAccount,
	}
	claims["stripe"] = jwt.StandardClaims{
		Audience:  t.Audience,
		ExpiresAt: t.ExpiresAt,
		Id:        t.Id,
		IssuedAt:  t.IssuedAt,
		Issuer:    t.Access.StripeKey,
		NotBefore: t.NotBefore,
		Subject:   t.Access.StripeAccount,
	}
	claims["slack"] = jwt.StandardClaims{
		Audience:  t.Audience,
		ExpiresAt: t.ExpiresAt,
		Id:        t.Id,
		IssuedAt:  t.IssuedAt,
		Issuer:    t.Access.SlackKey,
		NotBefore: t.NotBefore,
		Subject:   t.Access.SlackAccount,
	}
	claims["gcp"] = jwt.StandardClaims{
		Audience:  t.Audience,
		ExpiresAt: t.ExpiresAt,
		Id:        t.Id,
		IssuedAt:  t.IssuedAt,
		Issuer:    t.Access.GcpKey,
		NotBefore: t.NotBefore,
		Subject:   t.Access.GcpProject,
	}
	return &Claims{
		standardClaims: claims,
	}
}

type JWTMiddleware struct {
	middleware *jwtmiddleware.JWTMiddleware
}

func NewJWTMiddleware(key SignedKey) *JWTMiddleware {
	return &JWTMiddleware{jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key.SignedKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})}
}

func (j *JWTMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j.middleware.HandlerWithNext(w, r, next)
	}
}

func (j *JWTMiddleware) Check(w http.ResponseWriter, r *http.Request) error {
	return j.middleware.CheckJWT(w, r)
}

/*
Attributes() map[string]string
	Data()[]byte
	ID() string
*/
