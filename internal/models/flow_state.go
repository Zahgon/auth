package models

import (
	"time"

	"github.com/supabase/auth/internal/storage"

	"github.com/gofrs/uuid"
)

type FlowState struct {
	ID                   uuid.UUID  `json:"id" db:"id"`
	UserID               *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	AuthCode             *string    `json:"auth_code,omitempty" db:"auth_code"`
	AuthenticationMethod string     `json:"authentication_method" db:"authentication_method"`
	CodeChallenge        *string    `json:"code_challenge,omitempty" db:"code_challenge"`
	CodeChallengeMethod  *string    `json:"code_challenge_method,omitempty" db:"code_challenge_method"`
	ProviderType         string     `json:"provider_type" db:"provider_type"`
	ProviderAccessToken  string     `json:"provider_access_token" db:"provider_access_token"`
	ProviderRefreshToken string     `json:"provider_refresh_token" db:"provider_refresh_token"`
	AuthCodeIssuedAt     *time.Time `json:"auth_code_issued_at" db:"auth_code_issued_at"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`

	// OAuth context fields (previously stored in JWT state parameter)
	InviteToken        *string    `json:"invite_token,omitempty" db:"invite_token"`
	Referrer           *string    `json:"referrer,omitempty" db:"referrer"`
	OAuthClientStateID *uuid.UUID `json:"oauth_client_state_id,omitempty" db:"oauth_client_state_id"`
	LinkingTargetID    *uuid.UUID `json:"linking_target_id,omitempty" db:"linking_target_id"`
	EmailOptional      bool       `json:"email_optional" db:"email_optional"`
}

// FlowStateParams contains all parameters for creating a flow state
type FlowStateParams struct {
	ProviderType         string
	AuthenticationMethod AuthenticationMethod
	CodeChallenge        string // empty for implicit flow
	CodeChallengeMethod  string // empty for implicit flow
	UserID               *uuid.UUID
	InviteToken          string
	Referrer             string
	OAuthClientStateID   *uuid.UUID
	LinkingTargetID      *uuid.UUID
	EmailOptional        bool
}

type CodeChallengeMethod int

const (
	SHA256 CodeChallengeMethod = iota
	Plain
)

func (codeChallengeMethod CodeChallengeMethod) String() string {
	_ = "STUB: not implemented"
	return ""
}

func ParseCodeChallengeMethod(codeChallengeMethod string) (CodeChallengeMethod, error) {
	_ = "STUB: not implemented"
	return *new(CodeChallengeMethod), nil
}

type FlowType int

const (
	PKCEFlow FlowType = iota
	ImplicitFlow
)

func (flowType FlowType) String() string { _ = "STUB: not implemented"; return "" }

func (FlowState) TableName() string { _ = "STUB: not implemented"; return "" }

// NewFlowState creates a flow state for both PKCE and implicit flows.
// PKCE fields (AuthCode, CodeChallenge, CodeChallengeMethod) are only set
// if CodeChallenge is provided in params.
// Returns an error if CodeChallenge is provided but CodeChallengeMethod is invalid.
func NewFlowState(params FlowStateParams) (*FlowState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set PKCE fields only if code_challenge is provided

// Set optional context fields

// IsPKCE returns true if this flow state represents a PKCE flow
func (f *FlowState) IsPKCE() bool { _ = "STUB: not implemented"; return false }

func FindFlowStateByAuthCode(tx *storage.Connection, authCode string) (*FlowState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindFlowStateByID(tx *storage.Connection, id string) (*FlowState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindFlowStateByIDForUpdate finds a flow state by ID and locks the row with
// FOR UPDATE SKIP LOCKED to prevent concurrent modifications. If the row is
// already locked by another transaction, SKIP LOCKED causes the query to
// return no rows instead of blocking, which surfaces as FlowStateNotFoundError.
// The lock is held until the transaction commits or rolls back.
func FindFlowStateByIDForUpdate(tx *storage.Connection, id string) (*FlowState, error) {
	_ = "STUB: not implemented"

	// Pop does not provide a way to execute FOR UPDATE queries,
	// so we use a raw query to lock the row first.
	return nil, nil
}

func FindFlowStateByUserID(tx *storage.Connection, id string, authenticationMethod AuthenticationMethod) (*FlowState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FlowState) VerifyPKCE(codeVerifier string) error { _ = "STUB: not implemented"; return nil }

func (f *FlowState) IsExpired(expiryDuration time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FlowState) RecordAuthCodeIssuedAtTime(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}
