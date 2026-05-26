package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

// OAuthServerConsent represents user consent for an OAuth server client's access to specific scopes
type OAuthServerConsent struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	ClientID  uuid.UUID  `json:"-" db:"client_id"`
	Scopes    string     `json:"scopes" db:"scopes"`
	GrantedAt time.Time  `json:"granted_at" db:"granted_at"`
	RevokedAt *time.Time `json:"revoked_at" db:"revoked_at"`
}

// TableName returns the table name for the OAuthConsent model
func (OAuthServerConsent) TableName() string { _ = "STUB: not implemented"; return "" }

// NewOAuthConsent creates a new OAuth consent record
func NewOAuthServerConsent(userID uuid.UUID, clientID uuid.UUID, scopes []string) *OAuthServerConsent {
	_ = "STUB: not implemented"
	return nil
}

// GetScopeList returns the granted scopes as a slice
func (consent *OAuthServerConsent) GetScopeList() []string { _ = "STUB: not implemented"; return nil }

// HasScope checks if the consent includes a specific scope
func (consent *OAuthServerConsent) HasScope(scope string) bool {
	_ = "STUB: not implemented"
	return false
}

// HasAllScopes checks if the consent includes all of the requested scopes
func (consent *OAuthServerConsent) HasAllScopes(requestedScopes []string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsRevoked checks if the consent has been revoked
func (consent *OAuthServerConsent) IsRevoked() bool { _ = "STUB: not implemented"; return false }

// Revoke revokes the consent
func (consent *OAuthServerConsent) Revoke(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateScopes updates the granted scopes for this consent
func (consent *OAuthServerConsent) UpdateScopes(tx *storage.Connection, scopes []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Update granted time to reflect the change

// Validate performs basic validation on the OAuth consent
func (consent *OAuthServerConsent) Validate() error { _ = "STUB: not implemented"; return nil }

// Query functions for OAuth consents

// FindOAuthServerConsentByUserAndClient finds an OAuth consent by user and client
func FindOAuthServerConsentByUserAndClient(tx *storage.Connection, userID uuid.UUID, clientID uuid.UUID) (*OAuthServerConsent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No consent found (not an error)

// FindActiveOAuthServerConsentByUserAndClient finds an active (non-revoked) OAuth consent
func FindActiveOAuthServerConsentByUserAndClient(tx *storage.Connection, userID uuid.UUID, clientID uuid.UUID) (*OAuthServerConsent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No active consent found (not an error)

// FindOAuthServerConsentsByUser finds all OAuth consents for a user
func FindOAuthServerConsentsByUser(tx *storage.Connection, userID uuid.UUID, includeRevoked bool) ([]*OAuthServerConsent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpsertOAuthServerConsent creates or updates an OAuth consent
func UpsertOAuthServerConsent(tx *storage.Connection, consent *OAuthServerConsent) error {
	_ = "STUB: not implemented"
	return nil
}

// Update existing consent

// Un-revoke if previously revoked

// Create new consent

// RevokeOAuthServerConsentsByClient revokes all consents for a specific client
func RevokeOAuthServerConsentsByClient(tx *storage.Connection, clientID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// RevokeOAuthServerConsentsByUser revokes all consents for a specific user
func RevokeOAuthServerConsentsByUser(tx *storage.Connection, userID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}
