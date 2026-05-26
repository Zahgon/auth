package models

import (
	"database/sql/driver"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

// WebAuthnTransports is a JSON-serializable slice of authenticator transports.
type WebAuthnTransports []protocol.AuthenticatorTransport

func (t *WebAuthnTransports) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (t WebAuthnTransports) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// WebAuthnCredential maps to the webauthn_credentials table.
type WebAuthnCredential struct {
	ID              uuid.UUID          `json:"id" db:"id"`
	UserID          uuid.UUID          `json:"user_id" db:"user_id"`
	CredentialID    []byte             `json:"-" db:"credential_id"`
	PublicKey       []byte             `json:"-" db:"public_key"`
	AttestationType string             `json:"attestation_type" db:"attestation_type"`
	AAGUID          *uuid.UUID         `json:"aaguid,omitempty" db:"aaguid"`
	SignCount       uint32             `json:"sign_count" db:"sign_count"`
	Transports      WebAuthnTransports `json:"transports" db:"transports"`
	BackupEligible  bool               `json:"backup_eligible" db:"backup_eligible"`
	BackedUp        bool               `json:"backed_up" db:"backed_up"`
	FriendlyName    string             `json:"friendly_name" db:"friendly_name"`
	CreatedAt       time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" db:"updated_at"`
	LastUsedAt      *time.Time         `json:"last_used_at,omitempty" db:"last_used_at"`
}

func (WebAuthnCredential) TableName() string { _ = "STUB: not implemented"; return "" }

func NewWebAuthnCredential(userID uuid.UUID, cred *webauthn.Credential, friendlyName string) *WebAuthnCredential {
	_ = "STUB: not implemented"
	return nil
}

// ToWebAuthnCredential converts back to the library's Credential type for verification.
func (pc *WebAuthnCredential) ToWebAuthnCredential() webauthn.Credential {
	_ = "STUB: not implemented"
	return *new(webauthn.Credential)
}

func FindWebAuthnCredentialsByUserID(conn *storage.Connection, userID uuid.UUID) ([]*WebAuthnCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindWebAuthnCredentialByCredentialID(conn *storage.Connection, credentialID []byte) (*WebAuthnCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindWebAuthnCredentialByID(conn *storage.Connection, id uuid.UUID) (*WebAuthnCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindWebAuthnCredentialByIDAndUserID(conn *storage.Connection, id, userID uuid.UUID) (*WebAuthnCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CountWebAuthnCredentialsByUserID(conn *storage.Connection, userID uuid.UUID) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pc *WebAuthnCredential) UpdateSignCount(tx *storage.Connection, signCount uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *WebAuthnCredential) UpdateLastUsedAt(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *WebAuthnCredential) UpdateLastUsedWithSignCount(tx *storage.Connection, signCount uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *WebAuthnCredential) UpdateFriendlyName(tx *storage.Connection, friendlyName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *WebAuthnCredential) Delete(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteWebAuthnCredentialsByUserID(tx *storage.Connection, userID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}
