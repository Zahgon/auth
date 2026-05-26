package oauthserver

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/models"
)

// validateRedirectURIList validates a list of redirect URIs
func validateRedirectURIList(redirectURIs []string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

// validateGrantTypeList validates a list of grant types
func validateGrantTypeList(grantTypes []string) error { _ = "STUB: not implemented"; return nil }

// validateClientName validates a client name
func validateClientName(clientName string) error { _ = "STUB: not implemented"; return nil }

// validateClientURI validates a client URI
func validateClientURI(clientURI string) error { _ = "STUB: not implemented"; return nil }

// validateLogoURI validates a logo URI
func validateLogoURI(logoURI string) error { _ = "STUB: not implemented"; return nil }

// OAuthServerClientRegisterParams contains parameters for registering a new OAuth client
type OAuthServerClientRegisterParams struct {
	// Required fields
	RedirectURIs []string `json:"redirect_uris"`

	// Client type can be explicitly provided or inferred from token_endpoint_auth_method
	ClientType              string `json:"client_type,omitempty"`                // models.OAuthServerClientTypePublic or models.OAuthServerClientTypeConfidential
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"` // "none", "client_secret_basic", or "client_secret_post"

	GrantTypes []string `json:"grant_types,omitempty"`
	ClientName string   `json:"client_name,omitempty"`
	ClientURI  string   `json:"client_uri,omitempty"`
	LogoURI    string   `json:"logo_uri,omitempty"`

	// Internal field
	RegistrationType string `json:"-"`
}

// validate validates the OAuth client registration parameters
func (p *OAuthServerClientRegisterParams) validate() error {
	_ = "STUB: not implemented"
	// Validate redirect URIs (required for registration)
	return nil
}

// Validate grant types if provided

// Validate client name

// Validate client URI

// Validate logo URI

// Validate client_type if provided (defaults to confidential if not specified)

// Validate token_endpoint_auth_method if provided

// Validate consistency between client_type and token_endpoint_auth_method

// validateRedirectURI validates OAuth 2.1 redirect URIs as specific in
//
// * https://tools.ietf.org/html/rfc6749#section-3.1.2
//   - The redirection endpoint URI MUST be an absolute URI as defined by [RFC3986] Section 4.3.
//   - The endpoint URI MUST NOT include a fragment component.
//   - https://tools.ietf.org/html/rfc3986#section-4.3
//     absolute-URI  = scheme ":" hier-part [ "?" query ]
//   - https://tools.ietf.org/html/rfc6819#section-5.1.1
func validateRedirectURI(uri string) error { _ = "STUB: not implemented"; return nil }

// Must have scheme and host

// Block dangerous URI schemes that can lead to XSS or token leakage

// Only restrict HTTP (not HTTPS or custom schemes)
// HTTP is only allowed for localhost/loopback addresses

// All other schemes (https, custom schemes like myapp://* etc.) are allowed

// Must not have fragment

// generateClientSecret generates a secure random client secret
func generateClientSecret() string { _ = "STUB: not implemented"; return "" }

// This should never happen, but fallback to panic for security

// hashClientSecret hashes a client secret using SHA-256
func hashClientSecret(secret string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ValidateClientSecret validates a client secret against its hash using constant-time comparison
func ValidateClientSecret(providedSecret, storedHash string) bool {
	_ = "STUB: not implemented"
	return false
}

// registerOAuthServerClient creates a new OAuth server client with generated credentials
func (s *Server) registerOAuthServerClient(ctx context.Context, params *OAuthServerClientRegisterParams) (*models.OAuthServerClient, string, error) {
	_ = "STUB: not implemented"
	// Validate all parameters
	return nil, "", nil
}

// Set defaults

// Determine client type using centralized logic

// Determine token_endpoint_auth_method
// If explicitly provided, use it; otherwise set default based on client type
// Per RFC 7591: "If unspecified or omitted, the default is 'client_secret_basic'"
// For public clients, the default is 'none' since they don't have a client secret

// Only generate client secret for confidential clients

// getOAuthServerClient retrieves an OAuth client by ID
func (s *Server) getOAuthServerClient(ctx context.Context, clientID uuid.UUID) (*models.OAuthServerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteOAuthServerClient soft-deletes an OAuth client
func (s *Server) deleteOAuthServerClient(ctx context.Context, clientID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// Soft delete by setting deleted_at

// regenerateOAuthServerClientSecret regenerates a client secret for confidential clients
func (s *Server) regenerateOAuthServerClientSecret(ctx context.Context, clientID uuid.UUID) (*models.OAuthServerClient, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Only confidential clients can have their secrets regenerated

// Generate new client secret

// Update client with new secret hash

// OAuthServerClientUpdateParams contains parameters for updating an OAuth client
type OAuthServerClientUpdateParams struct {
	RedirectURIs            *[]string `json:"redirect_uris,omitempty"`
	GrantTypes              *[]string `json:"grant_types,omitempty"`
	ClientName              *string   `json:"client_name,omitempty"`
	ClientURI               *string   `json:"client_uri,omitempty"`
	LogoURI                 *string   `json:"logo_uri,omitempty"`
	TokenEndpointAuthMethod *string   `json:"token_endpoint_auth_method,omitempty"`
}

// isEmpty returns true if no fields are set for update
func (p *OAuthServerClientUpdateParams) isEmpty() bool { _ = "STUB: not implemented"; return false }

// validate validates the OAuth client update parameters
func (p *OAuthServerClientUpdateParams) validate() error {
	_ = "STUB: not implemented"
	// Validate redirect URIs if provided
	return nil
}

// Validate grant types if provided

// Validate client name if provided

// Validate client URI if provided

// Validate logo URI if provided

// Validate token endpoint auth method if provided

// updateOAuthServerClient updates an existing OAuth client
func (s *Server) updateOAuthServerClient(ctx context.Context, clientID uuid.UUID, params *OAuthServerClientUpdateParams) (*models.OAuthServerClient, error) {
	_ = "STUB: not implemented"
	// Validate all parameters
	return nil, nil
}

// Update only the provided fields
