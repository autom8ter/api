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

func (c *Identity) ToCommon(id string, meta map[string]string) (*common.Common, error) {
	return common.ToCommon(id, meta, c)
}

func UserMetadataFromCommon(c *common.Common) (*UserMetadata, error) {
	meta := &UserMetadata{}
	err := c.Unmarshal(meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

func AppMetadataFromCommon(c *common.Common) (*AppMetadata, error) {
	meta := &AppMetadata{}
	err := c.Unmarshal(meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

func AddressFromCommon(c *common.Common) (*AppMetadata, error) {
	meta := &AppMetadata{}
	err := c.Unmarshal(meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

func UserFromCommon(c *common.Common) (*User, error) {
	u := &User{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CardFromCommon(c *common.Common) (*Card, error) {
	u := &Card{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ProductFromCommon(c *common.Common) (*Product, error) {
	u := &Product{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func RoleFromCommon(c *common.Common) (*Role, error) {
	u := &Role{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func IdentityFromCommon(c *common.Common) (*Identity, error) {
	u := &Identity{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PlanFromCommon(c *common.Common) (*Plan, error) {
	u := &Plan{}
	err := c.Unmarshal(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *Address) Validate(fn func(c *Address) error) error {
	return fn(c)
}

func (c *Address) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *UserMetadata) Validate(fn func(c *UserMetadata) error) error {
	return fn(c)
}

func (c *UserMetadata) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *User) Validate(fn func(c *User) error) error {
	return fn(c)
}

func (c *User) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *Card) Validate(fn func(c *Card) error) error {
	return fn(c)
}

func (c *Card) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *AppMetadata) Validate(fn func(c *AppMetadata) error) error {
	return fn(c)
}

func (c *AppMetadata) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *Product) Validate(fn func(c *Product) error) error {
	return fn(c)
}

func (c *Product) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *Plan) Validate(fn func(c *Plan) error) error {
	return fn(c)
}

func (c *Plan) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *Identity) Validate(fn func(c *Identity) error) error {
	return fn(c)
}

func (c *Identity) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}

func (c *Role) Validate(fn func(c *Role) error) error {
	return fn(c)
}

func (c *Role) DataMap() map[string]interface{} {
	return common.Util.ToMap(c)
}
