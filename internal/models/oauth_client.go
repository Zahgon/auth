package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

// OAuth client type constants
const (
	OAuthServerClientTypePublic       = "public"
	OAuthServerClientTypeConfidential = "confidential"
)

// OAuth token endpoint authentication method constants
const (
	TokenEndpointAuthMethodNone              = "none"
	TokenEndpointAuthMethodClientSecretBasic = "client_secret_basic"
	TokenEndpointAuthMethodClientSecretPost  = "client_secret_post"
)

// OAuthServerClient represents an OAuth client application registered with this OAuth server
type OAuthServerClient struct {
	ID                      uuid.UUID `json:"client_id" db:"id"`
	ClientSecretHash        string    `json:"-" db:"client_secret_hash"`
	RegistrationType        string    `json:"registration_type" db:"registration_type"`
	ClientType              string    `json:"client_type" db:"client_type"`
	TokenEndpointAuthMethod string    `json:"token_endpoint_auth_method" db:"token_endpoint_auth_method"`

	RedirectURIs string     `json:"-" db:"redirect_uris"`
	GrantTypes   string     `json:"grant_types" db:"grant_types"`
	ClientName   *string    `json:"client_name,omitempty" db:"client_name"`
	ClientURI    *string    `json:"client_uri,omitempty" db:"client_uri"`
	LogoURI      *string    `json:"logo_uri,omitempty" db:"logo_uri"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// TableName returns the table name for the OAuthServerClient model
func (OAuthServerClient) TableName() string { _ = "STUB: not implemented"; return "" }

// BeforeSave is invoked before the OAuth client is saved to the database
func (c *OAuthServerClient) BeforeSave(tx *pop.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate performs basic validation on the OAuth client
func (c *OAuthServerClient) Validate() error { _ = "STUB: not implemented"; return nil }

// Confidential clients must have a client secret

// Public clients should not have a client secret (enforce PKCE instead)

// Apply default token_endpoint_auth_method per RFC 7591:
// "If unspecified or omitted, the default is 'client_secret_basic'"
// For public clients, the default is 'none' since they don't have a client secret

// Validate token_endpoint_auth_method

// Public clients must use 'none'

// Confidential clients cannot use 'none'

// GetRedirectURIs returns the redirect URIs as a slice
func (c *OAuthServerClient) GetRedirectURIs() []string { _ = "STUB: not implemented"; return nil }

// SetRedirectURIs sets the redirect URIs from a slice
func (c *OAuthServerClient) SetRedirectURIs(uris []string) { _ = "STUB: not implemented"; return }

// GetGrantTypes returns the grant types as a slice
func (c *OAuthServerClient) GetGrantTypes() []string { _ = "STUB: not implemented"; return nil }

// SetGrantTypes sets the grant types from a slice
func (c *OAuthServerClient) SetGrantTypes(types []string) { _ = "STUB: not implemented"; return }

// IsPublic returns true if the client is a public client
func (c *OAuthServerClient) IsPublic() bool { _ = "STUB: not implemented"; return false }

// IsConfidential returns true if the client is a confidential client
func (c *OAuthServerClient) IsConfidential() bool { _ = "STUB: not implemented"; return false }

// GetTokenEndpointAuthMethod returns the token endpoint auth method
func (c *OAuthServerClient) GetTokenEndpointAuthMethod() string {
	_ = "STUB: not implemented"
	return ""
}

// IsGrantTypeAllowed returns true if the client is allowed to use the specified grant type
func (c *OAuthServerClient) IsGrantTypeAllowed(grantType string) bool {
	_ = "STUB: not implemented"
	return false
}

// validateRedirectURI validates a single redirect URI according to OAuth 2.1 spec
func validateRedirectURI(uri string) error { _ = "STUB: not implemented"; return nil }

// Allow localhost for development, otherwise require HTTPS

// Error types for OAuth client operations
type OAuthServerClientNotFoundError struct{}

func (e OAuthServerClientNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e OAuthServerClientNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false
}

type InvalidRedirectURIError struct {
	URI string
}

func (e InvalidRedirectURIError) Error() string { _ = "STUB: not implemented"; return "" }

// Query functions for OAuth clients

// FindOAuthServerClientByID finds an OAuth client by ID
func FindOAuthServerClientByID(tx *storage.Connection, id uuid.UUID) (*OAuthServerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOAuthServerClient creates a new OAuth client in the database
func CreateOAuthServerClient(tx *storage.Connection, client *OAuthServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateOAuthServerClient updates an existing OAuth client in the database
func UpdateOAuthServerClient(tx *storage.Connection, client *OAuthServerClient) error {
	_ = "STUB: not implemented"
	return nil
}
