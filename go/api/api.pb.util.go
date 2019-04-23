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
		Home:        "/",
		Dashboard:   "/dashboard",
		Settings:    "/dashboard/settings",
		Logout:      "/logout",
		Callback:    "/callback",
		Login:       "/login",
		Subscribe:   "/subscribe",
		Unsubscribe: "/unsubscribe",
		Faq:         "/faq",
		Support:     "/support",
		Terms:       "/terms",
		Privacy:     "/privacy",
		Debug:       "/debug",
		Blog:        "/blog",
	}
}

func NewAuth0() *Paths {
	return &Paths{
		Home:        "/",
		Dashboard:   "/dashboard",
		Settings:    "/dashboard/settings",
		Logout:      "/logout",
		Callback:    "/callback",
		Login:       "/login",
		Subscribe:   "/subscribe",
		Unsubscribe: "/unsubscribe",
		Faq:         "/faq",
		Support:     "/support",
		Terms:       "/terms",
		Privacy:     "/privacy",
		Debug:       "/debug",
		Blog:        "/blog",
	}
}

func NewAuth0(debug bool, domain string, clientID string, clientSecret string, redirectURL string, scopes ...string) (*Auth0, error) {
	a := &Auth0{
		Domain:       domain,
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Redirect:     redirectURL,
		Scopes:       scopes,
	}
	if a.Domain == "" {
		return nil, errors.New("empty domain")
	}
	if a.ClientId == "" {
		return nil, errors.New("empty client id")
	}
	if a.ClientSecret == "" {
		return nil, errors.New("empty client secret")
	}
	if a.Redirect == "" {
		return nil, errors.New("empty redirect")
	}
	if a.Scopes == nil {
		a.Scopes = []string{"userid", "profile", "email"}
	}
	return a, nil
}
