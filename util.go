//go:generate godocdown -o README.md

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/api/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"io"
	"net/http"
)

var (
	DEFAULT_OAUTH_REDIRECT = common.ToString("http://localhost:8080/callback")
	DEFAULT_OAUTH_SCOPES   = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}
)

type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
	Payment PaymentServiceClient
	User    UserServiceClient
	Admin   AdminServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility: NewUtilityServiceClient(conn),
		Contact: NewContactServiceClient(conn),
		Payment: NewPaymentServiceClient(conn),
		User:    NewUserServiceClient(conn),
		Admin:   NewAdminServiceClient(conn),
	}
}

func SearchUSPhoneNumbersURL(account *common.String) *common.String {
	return common.ToString(fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/US/Local.json", account.Text))
}

func IncomingPhoneNumbersURL(account *common.String) *common.String {
	return common.ToString(fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json", account.Text))
}

func ChatServiceURL() *common.String {
	return common.ToString("https://chat.twilio.com/v2/Services")
}

func (s Scope) Normalize() *common.String {
	switch s {
	case Scope_PROFILE:
		return common.ToString("profile")
	case Scope_EMAIL:
		return common.ToString("email")
	case Scope_OPENID:
		return common.ToString("openid")
	case Scope_READ_USERS:
		return common.ToString("read:users")
	case Scope_READ_USER_IDP_TOKENS:
		return common.ToString("read:user_idp_tokens")
	case Scope_CREATE_USERS:
		return common.ToString("create:users")
	case Scope_READ_STATS:
		return common.ToString("read:stats")
	case Scope_CREATE_EMAIL_TEMPLATES:
		return common.ToString("create:email_templates")
	case Scope_UPDATE_EMAIL_TEMPLATES:
		return common.ToString("update:email_templates")
	case Scope_READ_EMAIL_TEMPLATES:
		return common.ToString("read:email_templates")
	case Scope_CREATE_RULES:
		return common.ToString("create:rules")
	case Scope_DELETE_RULES:
		return common.ToString("delete:rules")
	case Scope_UPDATE_RULES:
		return common.ToString("update:rules")
	case Scope_READ_RULES:
		return common.ToString("read:rules")
	case Scope_CREATE_ROLES:
		return common.ToString("create:rules")
	case Scope_DELETE_ROLES:
		return common.ToString("delete:rules")
	case Scope_UPDATE_ROLES:
		return common.ToString("update:rules")
	case Scope_READ_ROLES:
		return common.ToString("read:rules")
	case Scope_READ_LOGS:
		return common.ToString("read:logs")
	default:
		return common.ToString("")
	}
}

func NormalizeScopes(scopes ...Scope) *common.StringArray {
	s := &common.StringArray{}
	for _, scope := range scopes {
		s.Append(scope.Normalize())
	}
	return s
}

func (a *Auth) DefaultIfEmpty() {
	if len(a.Scopes) == 0 {
		a.Scopes = DEFAULT_OAUTH_SCOPES
	}
	if a.Redirect.IsEmpty() {
		a.Redirect = DEFAULT_OAUTH_REDIRECT
	}
}
func (a *Auth) Validate() error {
	if len(a.Scopes) == 0 {
		return common.ToError(errors.New("validation error: empty scope"), "")
	}
	if a.Redirect.IsEmpty() {
		return common.ToError(errors.New("validation error: empty redirect"), "")
	}
	if a.ClientId.IsEmpty() {
		return common.ToError(errors.New("validation error: empty clientid"), "")
	}
	if a.Domain.IsEmpty() {
		return common.ToError(errors.New("validation error: empty domain"), "")
	}
	if a.ClientSecret.IsEmpty() {
		return common.ToError(errors.New("validation error: empty secret"), "")
	}

	return nil
}

func UserFromSession(session *sessions.Session) (*User, error) {
	return session.Values["user"].(*User), nil
}

func (a *Auth) audienceAuthCodeOption(u URL) oauth2.AuthCodeOption {
	return oauth2.SetAuthURLParam("audience", u.Normalize(a.Domain).Text)

}

func (a *Auth) AuthCodeURL(state string, u URL) string {
	return a.config().AuthCodeURL(state, a.audienceAuthCodeOption(u))
}

func (a *Auth) config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     a.ClientId.Text,
		ClientSecret: a.ClientSecret.Text,
		RedirectURL:  a.Redirect.Text,
		Scopes:       NormalizeScopes(a.Scopes...).Array(),
		Endpoint: oauth2.Endpoint{
			AuthURL:  URL_AUTHORIZEURL.Normalize(a.Domain).Text,
			TokenURL: URL_TOKENURL.Normalize(a.Domain).Text,
		},
	}
}

func (u URL) Normalize(domain *common.String) *common.String {
	switch u {
	case URL_TOKENURL:
		return common.ToString(tokenURL(domain.Text))
	case URL_USER_INFOURL:
		return common.ToString(userInfoURL(domain.Text))
	case URL_USERSURL:
		return common.ToString(usersURL(domain.Text))
	case URL_AUTHORIZEURL:
		return common.ToString(authURL(domain.Text))
	case URL_SEARCH_USERSURL:
		return common.ToString(searchUsersURL(domain.Text))
	case URL_ROLESURL:
		return common.ToString(rolesURL(domain.Text))
	case URL_LOGSURL:
		return common.ToString(logsURL(domain.Text))
	case URL_GRANTSURL:
		return common.ToString(grantsURL(domain.Text))
	case URL_STATSURL:
		return common.ToString(statsURL(domain.Text))
	case URL_CLIENTSURL:
		return common.ToString(clientsURL(domain.Text))
	case URL_CONNECTIONSURL:
		return common.ToString(connectionsURL(domain.Text))
	case URL_RULESURL:
		return common.ToString(rulesURL(domain.Text))
	case URL_TENANTSURL:
		return common.ToString(tenantsURL(domain.Text))
	case URL_JWKSURL:
		return common.ToString(jWKSURL(domain.Text))
	case URL_DEVICEURL:
		return common.ToString(deviceCredentials(domain.Text))
	case URL_EMAILURL:
		return common.ToString(emailURL(domain.Text))
	case URL_EMAIL_TEMPLATEURL:
		return common.ToString(emailTemplateURL(domain.Text))
	default:
		return common.ToString("")
	}
}

func tokenURL(domain string) string {
	return "https://" + domain + "/oauth/token"
}

func userInfoURL(domain string) string {
	return "https://" + domain + "/userinfo"
}

func usersURL(domain string) string {
	return "https://" + domain + "/api/v2/users"
}

func authURL(domain string) string {
	return "https://" + domain + "/authorize"
}

func searchUsersURL(domain string) string {
	return "https://" + domain + "/api/v2/users-by-email"
}

func rolesURL(domain string) string {
	return "https://" + domain + "/api/v2/roles"
}

func logsURL(domain string) string {
	return "https://" + domain + "/api/v2/logs"
}

func grantsURL(domain string) string {
	return "https://" + domain + "/api/v2/grants"
}

func statsURL(domain string) string {
	return "https://" + domain + "/api/v2/stats/active-users"
}

func clientsURL(domain string) string {

	return "https://" + domain + "/api/v2/clients"
}

func clientGrantsURL(domain string) string {
	return "https://" + domain + "/api/v2/client-grants"
}

func connectionsURL(domain string) string {
	return "https://" + domain + "/api/v2/connections"
}

func customDomainsURL(domain string) string {
	return "https://" + domain + "/api/v2/custom-domains"
}

func rulesURL(domain string) string {
	return "https://" + domain + "/api/v2/rules"
}

func tenantsURL(domain string) string {
	return "https://" + domain + "/api/v2/tenants/settings"
}

func emailURL(domain string) string {
	return "https://" + domain + "/api/v2/emails/provider"
}

func emailTemplateURL(domain string) string {
	return "https://" + domain + "/api/v2/email-templates"
}

func deviceCredentials(domain string) string {
	return "https://" + domain + "/api/v2/device-credentials"
}

func jWKSURL(domain string) string {
	return "https://" + domain + "/.well-known/jwks.json"
}

func (c *Jwks) TokenCert(token *jwt.Token) (string, error) {
	var cert string
	for k, _ := range c.Keys {
		if token.Header["kid"] == c.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + c.Keys[k].X5C.Strings[0].Text + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := common.ToError(errors.New("creating certificate"), "Unable to find appropriate key.")
		return cert, err
	}
	return cert, nil
}

func (a *Auth) Token(ctx context.Context, code string) (*common.Token, error) {
	token, err := a.config().Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return common.TokenFromOAuthToken(token), nil
}

func RenderUserFunc(t *common.String) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := common.GetAuthSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if t.IsTemplate() {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			obj := session.Values["user"]
			bits, err := json.Marshal(obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u := &User{}
			err = json.Unmarshal(bits, u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.RenderBytes("", w, common.AsBytes(u))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		io.WriteString(w, t.Text)
		return
	}
}

func (p Plan) Normalize() *common.String {
	switch p {
	case Plan_BASIC:
		return common.ToString("basic")
	case Plan_PREMIUM:
		return common.ToString("premium")
	default:
		return common.ToString("free")
	}
}

func (p *SubscriptionResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *SMSResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *CallResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *User) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumberResource) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumber) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *JSONWebKeys) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *TokenQuery) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
func (p *Identity) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
func (p *Auth) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *Card) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *Jwks) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
