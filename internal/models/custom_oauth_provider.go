package models

import (
	"database/sql/driver"
	"time"

	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

// ProviderType represents the type of OAuth provider
type ProviderType string

const (
	ProviderTypeOAuth2 ProviderType = "oauth2"
	ProviderTypeOIDC   ProviderType = "oidc"
)

// CustomOAuthProvider represents a custom OAuth2 or OIDC provider configuration
type CustomOAuthProvider struct {
	ID           uuid.UUID    `db:"id" json:"id"`
	ProviderType ProviderType `db:"provider_type" json:"provider_type"`

	// Common fields for both OAuth2 and OIDC
	Identifier          string        `db:"identifier" json:"identifier"`
	Name                string        `db:"name" json:"name"`
	ClientID            string        `db:"client_id" json:"client_id"`
	ClientSecret        string        `db:"client_secret" json:"-"` // Encrypted via EncryptedString, never expose in JSON
	AcceptableClientIDs slices.String `db:"acceptable_client_ids" json:"acceptable_client_ids"`
	Scopes              slices.String `db:"scopes" json:"scopes"`
	PKCEEnabled         bool          `db:"pkce_enabled" json:"pkce_enabled"`
	AttributeMapping    slices.Map    `db:"attribute_mapping" json:"attribute_mapping"`
	AuthorizationParams slices.Map    `db:"authorization_params" json:"authorization_params"`
	Enabled             bool          `db:"enabled" json:"enabled"`
	EmailOptional       bool          `db:"email_optional" json:"email_optional"`

	// OIDC-specific fields (null for OAuth2 providers)
	Issuer            *string        `db:"issuer" json:"issuer,omitempty"`
	DiscoveryURL      *string        `db:"discovery_url" json:"discovery_url,omitempty"`
	SkipNonceCheck    bool           `db:"skip_nonce_check" json:"skip_nonce_check"`
	CachedDiscovery   *OIDCDiscovery `db:"cached_discovery" json:"discovery_document,omitempty"`
	DiscoveryCachedAt *time.Time     `db:"discovery_cached_at" json:"-"` // Internal caching, not exposed in API

	// OAuth2-specific fields (null for OIDC providers)
	AuthorizationURL *string `db:"authorization_url" json:"authorization_url,omitempty"`
	TokenURL         *string `db:"token_url" json:"token_url,omitempty"`
	UserinfoURL      *string `db:"userinfo_url" json:"userinfo_url,omitempty"`
	JwksURI          *string `db:"jwks_uri" json:"jwks_uri,omitempty"`

	// Timestamps
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (p CustomOAuthProvider) TableName() string { _ = "STUB: not implemented"; return "" }

// SetClientSecret encrypts and stores the client secret using the configured
// database encryption settings. If encryption is disabled, the secret is
// stored in plaintext (temporary fallback for now)
func (p *CustomOAuthProvider) SetClientSecret(secret string, dbEncryption conf.DatabaseEncryptionConfiguration) error {
	_ = "STUB: not implemented"
	return nil

	// Fallback: store in plaintext when encryption is not enabled.
}

// GetClientSecret decrypts and returns the client secret using the configured
// database decryption keys. It expects the client secret to be stored in
// encrypted form when encryption is enabled, but will also handle plaintext
// secrets (for deployments where encryption is not yet configured).
func (p *CustomOAuthProvider) GetClientSecret(dbEncryption conf.DatabaseEncryptionConfiguration) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Not an encrypted string – treat as plaintext.

// IsOIDC returns true if this is an OIDC provider
func (p *CustomOAuthProvider) IsOIDC() bool { _ = "STUB: not implemented"; return false }

// IsOAuth2 returns true if this is an OAuth2 provider
func (p *CustomOAuthProvider) IsOAuth2() bool { _ = "STUB: not implemented"; return false }

// GetProviderName returns the provider identifier (which already includes "custom:" prefix)
func (p *CustomOAuthProvider) GetProviderName() string { _ = "STUB: not implemented"; return "" }

// GetDiscoveryURL returns the discovery URL for OIDC providers
// If discovery_url is set, use that; otherwise construct from issuer
func (p *CustomOAuthProvider) GetDiscoveryURL() string { _ = "STUB: not implemented"; return "" }

// SetDiscoveryCache stores a validated OIDC discovery document and records the cache time.
func (p *CustomOAuthProvider) SetDiscoveryCache(discovery *OIDCDiscovery) {
	_ = "STUB: not implemented"
	return
}

// ClearDiscoveryCache removes the cached discovery document.
func (p *CustomOAuthProvider) ClearDiscoveryCache() { _ = "STUB: not implemented"; return }

// OIDCDiscovery represents cached OIDC discovery document
type OIDCDiscovery struct {
	Issuer                 string   `json:"issuer"`
	AuthorizationEndpoint  string   `json:"authorization_endpoint"`
	TokenEndpoint          string   `json:"token_endpoint"`
	UserinfoEndpoint       string   `json:"userinfo_endpoint,omitempty"`
	JwksURI                string   `json:"jwks_uri"`
	ScopesSupported        []string `json:"scopes_supported,omitempty"`
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`
	GrantTypesSupported    []string `json:"grant_types_supported,omitempty"`
	SubjectTypesSupported  []string `json:"subject_types_supported,omitempty"`
}

func (d *OIDCDiscovery) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (d *OIDCDiscovery) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// CRUD operations for CustomOAuthProvider

// FindCustomOAuthProviderByID finds a custom OAuth provider by ID
func FindCustomOAuthProviderByID(tx *storage.Connection, id uuid.UUID) (*CustomOAuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindCustomOAuthProviderByIdentifier finds a custom OAuth provider by identifier
func FindCustomOAuthProviderByIdentifier(tx *storage.Connection, identifier string) (*CustomOAuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAllCustomOAuthProviders finds all custom OAuth providers
func FindAllCustomOAuthProviders(tx *storage.Connection) ([]*CustomOAuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAllCustomOAuthProvidersByType finds all custom OAuth providers of a specific type
func FindAllCustomOAuthProvidersByType(tx *storage.Connection, providerType ProviderType) ([]*CustomOAuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountCustomOAuthProviders counts all custom OAuth providers
func CountCustomOAuthProviders(tx *storage.Connection) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// CreateCustomOAuthProvider creates a new custom OAuth provider
func CreateCustomOAuthProvider(tx *storage.Connection, provider *CustomOAuthProvider) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure text[] fields are never nil (Postgres NOT NULL constraint)

// UpdateCustomOAuthProvider updates an existing custom OAuth provider
func UpdateCustomOAuthProvider(tx *storage.Connection, provider *CustomOAuthProvider) error {
	_ = "STUB: not implemented"
	// Set updated_at timestamp explicitly in application code
	return nil
}

// DeleteCustomOAuthProvider deletes a custom OAuth provider
func DeleteCustomOAuthProvider(tx *storage.Connection, id uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}
