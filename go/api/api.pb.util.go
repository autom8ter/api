package api

import (
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"html/template"
	"io"
)

var Util = objectify.Default()

func (p *IDToken) Validate() error {
	if p.Name == "" {
		return errors.New("Validate: a profile name is required")
	}
	return nil
}

func (p *IDToken) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *IDToken) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}


func (p *AccessToken) Validate() error {
	if p.Sub == "" {
		return errors.New("Validate: sub is required")
	}
	return nil
}

func (p *AccessToken) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *AccessToken) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}