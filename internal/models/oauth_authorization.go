package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

// OAuthServerAuthorizationStatus represents the status of an OAuth server authorization request
type OAuthServerAuthorizationStatus string

const (
	OAuthServerAuthorizationPending  OAuthServerAuthorizationStatus = "pending"
	OAuthServerAuthorizationApproved OAuthServerAuthorizationStatus = "approved"
	OAuthServerAuthorizationDenied   OAuthServerAuthorizationStatus = "denied"
	OAuthServerAuthorizationExpired  OAuthServerAuthorizationStatus = "expired"
)

func (s OAuthServerAuthorizationStatus) String() string {
	_ = "STUB: not implemented"

	// OAuthServerResponseType represents the OAuth server response type
	return ""
}

type OAuthServerResponseType string

const (
	OAuthServerResponseTypeCode OAuthServerResponseType = "code"
)

func (rt OAuthServerResponseType) String() string {
	_ = "STUB: not implemented"

	// OAuthServerAuthorization represents an OAuth 2.1 server authorization request
	return ""
}

type OAuthServerAuthorization struct {
	ID                  uuid.UUID                      `json:"-" db:"id"`
	AuthorizationID     string                         `json:"authorization_id" db:"authorization_id"`
	ClientID            uuid.UUID                      `json:"-" db:"client_id"`
	UserID              *uuid.UUID                     `json:"user_id" db:"user_id"`
	RedirectURI         string                         `json:"redirect_uri" db:"redirect_uri"`
	Scope               string                         `json:"scope" db:"scope"`
	State               *string                        `json:"state,omitempty" db:"state"`
	Resource            *string                        `json:"resource,omitempty" db:"resource"`
	CodeChallenge       *string                        `json:"code_challenge,omitempty" db:"code_challenge"`
	CodeChallengeMethod *string                        `json:"code_challenge_method,omitempty" db:"code_challenge_method"`
	Nonce               *string                        `json:"nonce,omitempty" db:"nonce"` // OIDC nonce parameter
	ResponseType        OAuthServerResponseType        `json:"response_type" db:"response_type"`
	Status              OAuthServerAuthorizationStatus `json:"status" db:"status"`
	AuthorizationCode   *string                        `json:"-" db:"authorization_code"`
	CreatedAt           time.Time                      `json:"created_at" db:"created_at"`
	ExpiresAt           time.Time                      `json:"expires_at" db:"expires_at"`
	ApprovedAt          *time.Time                     `json:"approved_at" db:"approved_at"`

	// Relations with OAuth clients
	Client *OAuthServerClient `json:"client,omitempty" db:"-"`
}

// TableName returns the table name for the OAuthServerAuthorization model
func (OAuthServerAuthorization) TableName() string { _ = "STUB: not implemented"; return "" }

// NewOAuthServerAuthorizationParams contains parameters for creating a new OAuth server authorization
type NewOAuthServerAuthorizationParams struct {
	ClientID            uuid.UUID
	RedirectURI         string
	Scope               string
	State               string
	Resource            string
	CodeChallenge       string
	CodeChallengeMethod string
	TTL                 time.Duration
	Nonce               string
}

// NewOAuthServerAuthorization creates a new OAuth server authorization request without user (for initial flow)
func NewOAuthServerAuthorization(params NewOAuthServerAuthorizationParams) *OAuthServerAuthorization {
	_ = "STUB: not implemented"
	return nil
}

// Generate random ID for frontend

// No user yet

// Normalize code challenge method to lowercase for database storage
// Database enum expects 's256' and 'plain' (lowercase)

// IsExpired checks if the authorization request has expired
func (auth *OAuthServerAuthorization) IsExpired() bool { _ = "STUB: not implemented"; return false }

// SetUser sets the user ID for the authorization request (after login)
func (auth *OAuthServerAuthorization) SetUser(tx *storage.Connection, userID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetScopeList returns the scopes as a slice
func (auth *OAuthServerAuthorization) GetScopeList() []string {
	_ = "STUB: not implemented"
	return nil
}

// GenerateAuthorizationCode generates a new authorization code if not already set
func (auth *OAuthServerAuthorization) GenerateAuthorizationCode() string {
	_ = "STUB: not implemented"
	return ""
}

// Approve approves the authorization request and generates an authorization code
func (auth *OAuthServerAuthorization) Approve(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Deny denies the authorization request
func (auth *OAuthServerAuthorization) Deny(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// MarkExpired marks the authorization request as expired
func (auth *OAuthServerAuthorization) MarkExpired(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate performs basic validation on the OAuth authorization
func (auth *OAuthServerAuthorization) Validate() error { _ = "STUB: not implemented"; return nil }

// UserID can be nil initially for unauthenticated authorization requests
// It will be set when user authenticates

// VerifyPKCE verifies the PKCE code verifier against the stored challenge
func (auth *OAuthServerAuthorization) VerifyPKCE(codeVerifier string) error {
	_ = "STUB: not implemented"
	return nil
}

// No PKCE challenge stored, verification passes

// Use the shared PKCE verification function

// Query functions for OAuth authorizations

// FindOAuthServerAuthorizationByID finds an OAuth authorization by authorization_id
func FindOAuthServerAuthorizationByID(tx *storage.Connection, authorizationID string) (*OAuthServerAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindOAuthServerAuthorizationByIDForUpdate finds an OAuth authorization by
// authorization_id and locks the row with FOR UPDATE SKIP LOCKED.
// Must be called inside a transaction.
func FindOAuthServerAuthorizationByIDForUpdate(tx *storage.Connection, authorizationID string) (*OAuthServerAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindOAuthServerAuthorizationByCode finds an OAuth authorization by authorization code
func FindOAuthServerAuthorizationByCode(tx *storage.Connection, code string) (*OAuthServerAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load client relationship (always present)

// CreateOAuthServerAuthorization creates a new OAuth authorization in the database
func CreateOAuthServerAuthorization(tx *storage.Connection, auth *OAuthServerAuthorization) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanupExpiredOAuthServerAuthorizations marks expired authorizations as expired
func CleanupExpiredOAuthServerAuthorizations(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Error types for OAuth authorization operations

type OAuthServerAuthorizationNotFoundError struct{}

func (e OAuthServerAuthorizationNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e OAuthServerAuthorizationNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false
}
