package api

import "github.com/autom8ter/api/common"

func (c *Card) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *User) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *Product) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *Role) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *Plan) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *Address) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *AppMetadata) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func (c *UserMetadata) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}
