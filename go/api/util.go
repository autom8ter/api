package api

import (
	"encoding/json"
	"github.com/autom8ter/objectify"
	"google.golang.org/genproto/googleapis/pubsub/v1"
	"io"
	"os"
)

var Util = objectify.Default()

//////////////////////////////////////////////////////

func AccessFromJSON(j *JSON) *Access {
	a := &Access{}
	_ = json.Unmarshal(j.Data, a)
	return a
}

func AccessFromEnv() *Access {
	return &Access{
		TwilioAccount: os.Getenv("TWILIO_ACCOUNT"),
		TwilioKey:     os.Getenv("TWILIO_KEY"),
		SendgridKey:   os.Getenv("SENDGRID_KEY"),
		SlackKey:      os.Getenv("SLACK_KEY"),
		StripeKey:     os.Getenv("STRIPE_KEY"),
		EmailAddress: &EmailAddress{
			Address: os.Getenv("EMAIL_ADDRESS"),
			Name:    os.Getenv("EMAIL_NAME"),
		},
		LogConfig: &LogConfig{
			Username: os.Getenv("SLACK_LOG_USERNAME"),
			Channel:  os.Getenv("SLACK_LOG_CHANNEL"),
		},
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

func (j *Attachment) MarshalJSON() []byte {
	return Util.MarshalJSON(j)
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

func (j *PubSubMessage) MarshalJSON() []byte {
	return Util.MarshalJSON(j.Message)
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
func (j *JSON) MarshalJSON() []byte {
	return Util.MarshalJSON(j)
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

func (j *Customer) MarshalJSON() []byte {
	return Util.MarshalJSON(j)
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
