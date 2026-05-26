package api

import (
	"net/http"

	jwk "github.com/lestrrat-go/jwx/v2/jwk"
)

type JwksResponse struct {
	Keys []jwk.Key `json:"keys"`
}

func (a *API) WellKnownJwks(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// don't expose hmac jwk in endpoint

// OpenIDConfigurationResponse represents both OIDC Discovery and OAuth 2.0 Authorization Server Metadata
// This unified response serves both:
// - /.well-known/openid-configuration (OIDC Discovery per OpenID Connect Discovery 1.0)
// - /.well-known/oauth-authorization-server (OAuth Authorization Server Metadata per RFC 8414)
//
// Since OIDC Discovery extends RFC 8414, a single response structure satisfies both specifications.
type OpenIDConfigurationResponse struct {
	// Core Discovery Fields (Required by both OIDC and OAuth 2.0)
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	JWKSURL               string `json:"jwks_uri"`
	UserInfoEndpoint      string `json:"userinfo_endpoint,omitempty"` // OIDC-specific
	RegistrationEndpoint  string `json:"registration_endpoint,omitempty"`

	// Supported Parameters
	ScopesSupported                   []string `json:"scopes_supported,omitempty"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	ResponseModesSupported            []string `json:"response_modes_supported,omitempty"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`               // OIDC-specific
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"` // OIDC-specific
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported                   []string `json:"claims_supported,omitempty"`       // OIDC-specific
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"` // OAuth 2.1/PKCE
}

// WellKnownOpenID handles both OIDC Discovery and OAuth 2.0 Authorization Server Metadata endpoints
// This unified handler serves:
// - GET /.well-known/openid-configuration (OIDC Discovery)
// - GET /.well-known/oauth-authorization-server (RFC 8414)
//
// Both endpoints return the same comprehensive metadata since OIDC Discovery is a superset of OAuth 2.0 metadata
func (a *API) WellKnownOpenID(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure issuer doesn't end with a slash to avoid double slashes in URLs

// OAuth 2.1 / OIDC Supported Features

// TODO :: should create this based on signing key config?

// OIDC Standard Claims

// Include registration endpoint if dynamic registration is enabled
