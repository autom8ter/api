package api

import (
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
)

var Util = objectify.Default()

func (p *Profile) Validate() error {
	if p.Name == "" {
		return errors.New("Validate: a profile name is required")
	}
	return nil
}