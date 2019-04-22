package api

import (
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"html/template"
	"io"
)

var Util = objectify.Default()

func (p *Profile) Validate() error {
	if p.Name == "" {
		return errors.New("Validate: a profile name is required")
	}
	return nil
}

func (p *Profile) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *Profile) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}