package static

import (
	"github.com/Masterminds/sprig"
	"github.com/autom8ter/objectify"
	"html/template"
	"io"
)

var util = objectify.Default()

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderAsset(assetName string, data interface{}, w io.Writer) error {
	bits, err := Asset(assetName)
	if err != nil {
		return err
	}
	tmpl, err := template.New("").Funcs(r.FuncMap()).Parse(string(bits))
	return tmpl.Execute(w, data)
}

func (r *Renderer) FuncMap() template.FuncMap {
	mapFuncs := sprig.FuncMap()

	return mapFuncs
}

func (r *Renderer) Functions() []string {
	mapFuncs := r.FuncMap()
	keys := []string{}
	for k, _ := range mapFuncs {
		keys = append(keys, k)
	}
	return keys
}
