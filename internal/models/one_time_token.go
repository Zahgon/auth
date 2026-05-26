package models

import (
	"database/sql/driver"
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type OneTimeTokenType int

const (
	ConfirmationToken OneTimeTokenType = iota
	ReauthenticationToken
	RecoveryToken
	EmailChangeTokenNew
	EmailChangeTokenCurrent
	PhoneChangeToken
)

func (t OneTimeTokenType) String() string { _ = "STUB: not implemented"; return "" }

func ParseOneTimeTokenType(s string) (OneTimeTokenType, error) {
	_ = "STUB: not implemented"
	return *new(OneTimeTokenType), nil
}

func (t OneTimeTokenType) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (t *OneTimeTokenType) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

type OneTimeTokenNotFoundError struct {
}

func (e OneTimeTokenNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e OneTimeTokenNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

type OneTimeToken struct {
	ID uuid.UUID `json:"id" db:"id"`

	UserID    uuid.UUID        `json:"user_id" db:"user_id"`
	TokenType OneTimeTokenType `json:"token_type" db:"token_type"`

	TokenHash string `json:"token_hash" db:"token_hash"`
	RelatesTo string `json:"relates_to" db:"relates_to"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (OneTimeToken) TableName() string { _ = "STUB: not implemented"; return "" }

func ClearAllOneTimeTokensForUser(tx *storage.Connection, userID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

func ClearOneTimeTokenForUser(tx *storage.Connection, userID uuid.UUID, tokenType OneTimeTokenType) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateOneTimeToken(tx *storage.Connection, userID uuid.UUID, relatesTo, tokenHash string, tokenType OneTimeTokenType) error {
	_ = "STUB: not implemented"
	return nil
}

func FindOneTimeToken(tx *storage.Connection, tokenHash string, tokenTypes ...OneTimeTokenType) (*OneTimeToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G602

// FindUserByConfirmationToken finds users with the matching confirmation token.
func FindUserByConfirmationOrRecoveryToken(tx *storage.Connection, token string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByConfirmationToken finds users with the matching confirmation token.
func FindUserByConfirmationToken(tx *storage.Connection, token string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByRecoveryToken finds a user with the matching recovery token.
func FindUserByRecoveryToken(tx *storage.Connection, token string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByEmailChangeToken finds a user with the matching email change token.
func FindUserByEmailChangeToken(tx *storage.Connection, token string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByEmailChangeCurrentAndAudience finds a user with the matching email change and audience.
func FindUserByEmailChangeCurrentAndAudience(tx *storage.Connection, email, token, aud string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByEmailChangeNewAndAudience finds a user with the matching email change and audience.
func FindUserByEmailChangeNewAndAudience(tx *storage.Connection, email, token, aud string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserForEmailChange finds a user requesting for an email change
func FindUserForEmailChange(tx *storage.Connection, email, token, aud string, secureEmailChangeEnabled bool) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
