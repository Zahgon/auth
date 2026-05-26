package models

import (
	"database/sql/driver"

	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type Challenge struct {
	ID                  uuid.UUID            `json:"challenge_id" db:"id"`
	FactorID            uuid.UUID            `json:"factor_id" db:"factor_id"`
	CreatedAt           time.Time            `json:"created_at" db:"created_at"`
	VerifiedAt          *time.Time           `json:"verified_at,omitempty" db:"verified_at"`
	IPAddress           string               `json:"ip_address" db:"ip_address"`
	Factor              *Factor              `json:"factor,omitempty" belongs_to:"factor"`
	OtpCode             string               `json:"otp_code,omitempty" db:"otp_code"`
	WebAuthnSessionData *WebAuthnSessionData `json:"web_authn_session_data,omitempty" db:"web_authn_session_data"`
}

type WebAuthnSessionData struct {
	*webauthn.SessionData
}

func (s *WebAuthnSessionData) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Handle byte and string as a precaution, in postgres driver, json/jsonb should be returned as []byte

func (s *WebAuthnSessionData) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (ws *WebAuthnSessionData) ToChallenge(factorID uuid.UUID, ipAddress string) *Challenge {
	_ = "STUB: not implemented"
	return nil
}

func (Challenge) TableName() string { _ = "STUB: not implemented"; return "" }

// Update the verification timestamp
func (c *Challenge) Verify(tx *storage.Connection) error { _ = "STUB: not implemented"; return nil }

func (c *Challenge) HasExpired(expiryDuration float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Challenge) GetExpiryTime(expiryDuration float64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *Challenge) SetOtpCode(otpCode string, encrypt bool, encryptionKeyID, encryptionKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Challenge) GetOtpCode(decryptionKeys map[string]string, encrypt bool, encryptionKeyID string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}
