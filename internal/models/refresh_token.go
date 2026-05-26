package models

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

// RefreshToken is the database model for refresh tokens.
type RefreshToken struct {
	ID int64 `db:"id"`

	Token string `db:"token"`

	UserID uuid.UUID `db:"user_id"`

	Parent    storage.NullString `db:"parent"`
	SessionId *uuid.UUID         `db:"session_id"`

	Revoked   bool      `db:"revoked"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	DONTUSEINSTANCEID uuid.UUID `json:"-" db:"instance_id"`
}

func (RefreshToken) TableName() string { _ = "STUB: not implemented"; return "" }

// GrantParams is used to pass session-specific parameters when issuing a new
// refresh token to authenticated users.
type GrantParams struct {
	FactorID *uuid.UUID

	SessionNotAfter *time.Time
	SessionTag      *string

	OAuthClientID *uuid.UUID
	Scopes        *string

	UserAgent string
	IP        string
}

func (g *GrantParams) FillGrantParams(r *http.Request) { _ = "STUB: not implemented"; return }

// GrantAuthenticatedUser creates a refresh token for the provided user.
func GrantAuthenticatedUser(tx *storage.Connection, user *User, params GrantParams) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GrantRefreshTokenSwap swaps a refresh token for a new one, revoking the provided token.
func GrantRefreshTokenSwap(config conf.AuditLogConfiguration, r *http.Request, tx *storage.Connection, user *User, token *RefreshToken) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RevokeTokenFamily revokes all refresh tokens that descended from the provided token.
func RevokeTokenFamily(tx *storage.Connection, token *RefreshToken) error {
	_ = "STUB: not implemented"
	return nil
}

func FindTokenBySessionID(tx *storage.Connection, sessionId *uuid.UUID) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Session) ApplyGrantParams(params *GrantParams) { _ = "STUB: not implemented"; return }

func (s *Session) SetupRefreshTokenData(dbEncryption conf.DatabaseEncryptionConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Session) UpdateRefreshTokenCounterAndHmacKey(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func createRefreshToken(tx *storage.Connection, user *User, oldToken *RefreshToken, params *GrantParams) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
