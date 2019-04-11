package api

import (
	"context"
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/objectify"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"os"
)

var Util = objectify.Default()

func Cmd(name, description string, fn func() error) *cobra.Command {
	return &cobra.Command{
		Use:  name,
		Long: description,
		Run: func(cmd *cobra.Command, args []string) {
			if err := fn(); err != nil {
				Util.Fatalln(err.Error())
			}
		},
	}
}

type ClientSet struct {
	UserSet UserServiceClient
}

type Server struct {
	Addr string
	UserServiceServer
	driver.PluginFunc
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

func NewServer(addr string, server UserServiceServer) *Server {
	s := &Server{
		Addr:              addr,
		UserServiceServer: server,
	}
	s.PluginFunc = func(server *grpc.Server) {
		RegisterUserServiceServer(server, s)
	}
	return s
}

func (s *Server) Serve(debug bool) error {
	return engine.Serve(s.Addr, debug, s)
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

func (j *Attachment) AsMap() map[string]interface{} {
	return Util.ToMap(j)
}

//////////////////////////////////////////////////////

func (p *PubSubMessage) PubSubMessage() *pubsub.PubsubMessage {
	return p.Message
}

func (p *PubSubTopic) PubSubTopic() *pubsub.Topic {
	return p.Topic
}

type PubSubFunc func(a *PubSubMessage)

func NewPubSub(opts ...PubSubFunc) *PubSubMessage {
	a := &PubSubMessage{}
	for _, o := range opts {
		o(a)
	}
	return a
}

type PubSubTopicFunc func(a *PubSubTopic)

func NewPubSubTopic(opts ...PubSubTopicFunc) *PubSubTopic {
	a := &PubSubTopic{}
	for _, o := range opts {
		o(a)
	}
	return a
}

func (j *PubSubMessage) MarshalJSON() ([]byte, error) {
	return Util.MarshalJSON(j.Message), nil
}

func (j *PubSubMessage) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(Util.MarshalJSON(j.Message), obj)
}

func (j *PubSubMessage) CompileHTML(text string, w io.Writer) error {
	return Util.RenderHTML(text, j.Message, w)
}

func (j *PubSubMessage) CompileTXT(text string, w io.Writer) error {
	return Util.RenderTXT(text, j.Message, w)
}

func (j *PubSubMessage) AsMap() map[string]interface{} {
	return Util.ToMap(j)
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

func (j *JSON) AsMap() map[string]interface{} {
	return Util.ToMap(j)
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

func (j *Customer) AsMap() map[string]interface{} {
	return Util.ToMap(j)
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
		token.Header[k]=v
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
		token.Header[k]=v
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
		token.Header[k]=v
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
		token.Header[k]=v
	}
	tok, err := token.SignedString([]byte(secret))
	return &SignedKey{SignedKey: tok}, err
}


func (c *Claims) SignedGCPToken(secret string, headers map[string]string) (*SignedKey, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.standardClaims["gcp"])
	token.Header = map[string]interface{}{}
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
