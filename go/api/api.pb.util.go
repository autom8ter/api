package api

import (
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"html/template"
	"io"
)

var Util = objectify.Default()

func (p *UserInfo) Validate() error {
	if p.Name == "" {
		return errors.New("Validate: UserInfo requires a name")
	}
	return Util.Validate(p)
}

func (p *UserInfo) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *UserInfo) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *UserMetadata) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *UserMetadata) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *AppMetadata) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *AppMetadata) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *Tokens) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *Tokens) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *Paths) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *Paths) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func DefaultPaths() *Paths {
	return &Paths{
		Home:                 "/",
		Dashboard:            "/dashboard",
		Settings:             "/dashboard/settings",
		Logout:               "/logout",
		Callback:             "/callback",
		Login:                "/login",
		Subscribe:            "/subscribe",
		Unsubscribe:          "/unsubscribe",
		Faq:                  "/faq",
		Support:              "/support",
		Terms:                "/terms",
		Privacy:              "/privacy",
		Debug:                "/debug",
		Blog:                 "/blog",
	}
}
