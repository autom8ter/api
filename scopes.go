package api

func (s Scope) Normalize() string {
	switch s {
	case Scope_PROFILE:
		return "profile"
	case Scope_EMAIL:
		return "email"
	case Scope_OPENID:
		return "openid"
	case Scope_READ_USERS:
		return "read:users"
	case Scope_READ_USER_IDP_TOKENS:
		return "read:user_idp_tokens"
	case Scope_CREATE_USERS:
		return "create:users"
	case Scope_READ_STATS:
		return "read:stats"
	case Scope_CREATE_EMAIL_TEMPLATES:
		return "create:email_templates"
	case Scope_UPDATE_EMAIL_TEMPLATES:
		return "update:email_templates"
	case Scope_READ_EMAIL_TEMPLATES:
		return "read:email_templates"
	case Scope_CREATE_RULES:
		return "create:rules"
	case Scope_DELETE_RULES:
		return "delete:rules"
	case Scope_UPDATE_RULES:
		return "update:rules"
	case Scope_READ_RULES:
		return "read:rules"
	case Scope_CREATE_ROLES:
		return "create:rules"
	case Scope_DELETE_ROLES:
		return "delete:rules"
	case Scope_UPDATE_ROLES:
		return "update:rules"
	case Scope_READ_ROLES:
		return "read:rules"
	case Scope_READ_LOGS:
		return "read:logs"
	default:
		return ""
	}
}

func NormalizeScopes(scopes ...Scope) []string {
	s := []string{}
	for _, scope := range scopes {
		s = append(s, scope.Normalize())
	}
	return s
}
