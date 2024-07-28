package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

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

func NewSecurityScheme(s *openapi3.SecurityScheme) (*SecurityScheme, error) {
	flows, err := NewOAuthFlows(s.Flows)
	if err != nil {
		return nil, fmt.Errorf("new flows: %w", err)
	}
	return &SecurityScheme{
		Type:        SecuritySchemeType(s.Type),
		Description: s.Description,

		Name: s.Name,
		In:   SecuritySchemeIn(s.In),

		Scheme:       s.Scheme,
		BearerFormat: s.BearerFormat,

		Flows: flows,

		OpenIDConnectURL: s.OpenIdConnectUrl,
	}, nil
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
	Implicit          Maybe[OAuthFlow]
	Password          Maybe[OAuthFlow]
	ClientCredentials Maybe[OAuthFlow]
	AuthorizationCode Maybe[OAuthFlow]
}

func NewOAuthFlows(s *openapi3.OAuthFlows) (zero OAuthFlows, _ error) {
	var out OAuthFlows
	if s == nil {
		return out, nil
	}
	if s.Implicit != nil {
		flow, err := NewOAuthFlow(s.Implicit)
		if err != nil {
			return zero, fmt.Errorf("new implicit flow: %w", err)
		}
		out.Implicit = Just(flow)
	}
	if s.Password != nil {
		flow, err := NewOAuthFlow(s.Password)
		if err != nil {
			return zero, fmt.Errorf("new password flow: %w", err)
		}
		out.Password = Just(flow)
	}
	if s.ClientCredentials != nil {
		flow, err := NewOAuthFlow(s.ClientCredentials)
		if err != nil {
			return zero, fmt.Errorf("new client credentials flow: %w", err)
		}
		out.ClientCredentials = Just(flow)
	}
	if s.AuthorizationCode != nil {
		flow, err := NewOAuthFlow(s.AuthorizationCode)
		if err != nil {
			return zero, fmt.Errorf("new authorization code flow: %w", err)
		}
		out.AuthorizationCode = Just(flow)
	}
	return out, nil
}

type OAuthFlow struct {
	AuthorizationURL string
	TokenURL         string
	RefreshURL       string
	Scopes           Map[string]
}

func NewOAuthFlow(s *openapi3.OAuthFlow) (zero OAuthFlow, _ error) {
	scopesMap, err := NewMap[string, string](s.Scopes, func(s string) (string, error) {
		return s, nil
	})
	if err != nil {
		return zero, fmt.Errorf("new scopes map: %w", err)
	}
	return OAuthFlow{
		AuthorizationURL: s.AuthorizationURL,
		TokenURL:         s.TokenURL,
		RefreshURL:       s.RefreshURL,
		Scopes:           scopesMap,
	}, nil
}
