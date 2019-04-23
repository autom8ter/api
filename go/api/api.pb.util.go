package api

import (
	"github.com/autom8ter/objectify"
	"html/template"
	"io"
)

var Util = objectify.Default()

func (p *IDToken) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *IDToken) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *AccessToken) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *AccessToken) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *UserMetadata) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *UserMetadata) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}
