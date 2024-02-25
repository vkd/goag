package specification

import "github.com/getkin/kin-openapi/openapi3"

type SecurityScheme struct {
	NoRef[SecurityScheme]

	// Any
	Type        SecuritySchemeType
	Description string

	// apiKey
	Name string
	In   SecuritySchemeIn

	// http
	Scheme       string
	BearerFormat string

	// oauth2
	Flows OAuthFlows

	// openIdConnect
	OpenIDConnectURL string
}

func NewSecurityScheme(s *openapi3.SecurityScheme) *SecurityScheme {
	return &SecurityScheme{
		Type:        SecuritySchemeType(s.Type),
		Description: s.Description,

		Name: s.Name,
		In:   SecuritySchemeIn(s.In),

		Scheme:       s.Scheme,
		BearerFormat: s.BearerFormat,

		Flows: NewOAuthFlows(s.Flows),

		OpenIDConnectURL: s.OpenIdConnectUrl,
	}
}

var _ Ref[SecurityScheme] = (*SecurityScheme)(nil)

func (s *SecurityScheme) Value() *SecurityScheme { return s }

type SecuritySchemeType string

const (
	SecuritySchemeTypeApiKey        SecuritySchemeType = "apiKey"
	SecuritySchemeTypeHTTP          SecuritySchemeType = "http"
	SecuritySchemeTypeMutualTLS     SecuritySchemeType = "mutualTLS"
	SecuritySchemeTypeOAuth2        SecuritySchemeType = "oauth2"
	SecuritySchemeTypeOpenIDConnect SecuritySchemeType = "openIdConnect"
)

type SecuritySchemeIn string

const (
	SecuritySchemeInQuery  SecuritySchemeIn = "query"
	SecuritySchemeInHeader SecuritySchemeIn = "header"
	SecuritySchemeInCookie SecuritySchemeIn = "cookie"
)

type OAuthFlows struct {
	Implicit          Optional[OAuthFlow]
	Password          Optional[OAuthFlow]
	ClientCredentials Optional[OAuthFlow]
	AuthorizationCode Optional[OAuthFlow]
}

func NewOAuthFlows(s *openapi3.OAuthFlows) OAuthFlows {
	var out OAuthFlows
	if s == nil {
		return out
	}
	if s.Implicit != nil {
		out.Implicit = NewOptional(NewOAuthFlow(s.Implicit))
	}
	if s.Password != nil {
		out.Password = NewOptional(NewOAuthFlow(s.Password))
	}
	if s.ClientCredentials != nil {
		out.ClientCredentials = NewOptional(NewOAuthFlow(s.ClientCredentials))
	}
	if s.AuthorizationCode != nil {
		out.AuthorizationCode = NewOptional(NewOAuthFlow(s.AuthorizationCode))
	}
	return out
}

type OAuthFlow struct {
	AuthorizationURL string
	TokenURL         string
	RefreshURL       string
	Scopes           Map[string]
}

func NewOAuthFlow(s *openapi3.OAuthFlow) OAuthFlow {
	return OAuthFlow{
		AuthorizationURL: s.AuthorizationURL,
		TokenURL:         s.TokenURL,
		RefreshURL:       s.RefreshURL,
		Scopes: NewMap[string, string](s.Scopes, func(s string) string {
			return s
		}),
	}
}
