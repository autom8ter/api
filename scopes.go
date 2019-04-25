package api

type Scope int

const (
	OPENID Scope = iota
	PROFILE
	EMAIL
)

func (s Scope) String() string {
	switch s {
	case PROFILE:
		return "profile"
	case EMAIL:
		return "email"
	case OPENID:
		return "openid"
	default:
		return "openid"
	}
}
