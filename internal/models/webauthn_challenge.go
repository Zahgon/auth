package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

const (
	WebAuthnChallengeTypeSignup         = "signup"
	WebAuthnChallengeTypeRegistration   = "registration"
	WebAuthnChallengeTypeAuthentication = "authentication"
)

// WebAuthnChallenge maps to the webauthn_challenges table.
type WebAuthnChallenge struct {
	ID            uuid.UUID            `json:"id" db:"id"`
	UserID        *uuid.UUID           `json:"user_id,omitempty" db:"user_id"`
	ChallengeType string               `json:"challenge_type" db:"challenge_type"`
	SessionData   *WebAuthnSessionData `json:"session_data" db:"session_data"`
	CreatedAt     time.Time            `json:"created_at" db:"created_at"`
	ExpiresAt     time.Time            `json:"expires_at" db:"expires_at"`
}

func (WebAuthnChallenge) TableName() string { _ = "STUB: not implemented"; return "" }

func NewWebAuthnChallenge(userID *uuid.UUID, challengeType string, sessionData *WebAuthnSessionData, expiresAt time.Time) *WebAuthnChallenge {
	_ = "STUB: not implemented"
	return nil
}

func FindWebAuthnChallengeByID(conn *storage.Connection, id uuid.UUID) (*WebAuthnChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsumeWebAuthnChallengeByID atomically deletes and returns a challenge row
func ConsumeWebAuthnChallengeByID(conn *storage.Connection, id uuid.UUID, challengeType string, userID *uuid.UUID) (*WebAuthnChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *WebAuthnChallenge) IsExpired() bool { _ = "STUB: not implemented"; return false }

func (c *WebAuthnChallenge) Delete(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}
