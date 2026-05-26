package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type FactorState int

const (
	FactorStateUnverified FactorState = iota
	FactorStateVerified
)

func (factorState FactorState) String() string { _ = "STUB: not implemented"; return "" }

const TOTP = "totp"
const Phone = "phone"
const WebAuthn = "webauthn"

type AuthenticationMethod int

const (
	OAuth AuthenticationMethod = iota
	PasswordGrant
	OTP
	TOTPSignIn
	MFAPhone
	MFAWebAuthn
	SSOSAML
	Recovery
	Invite
	MagicLink
	EmailSignup
	EmailChange
	TokenRefresh
	Anonymous
	Web3
	OAuthProviderAuthorizationCode
	PasskeyLogin
)

func (authMethod AuthenticationMethod) IsRecovery() bool { _ = "STUB: not implemented"; return false }

func (authMethod AuthenticationMethod) String() string { _ = "STUB: not implemented"; return "" }

func ParseAuthenticationMethod(authMethod string) (AuthenticationMethod, error) {
	_ = "STUB: not implemented"
	return *new(AuthenticationMethod), nil
}

type Factor struct {
	ID uuid.UUID `json:"id" db:"id"`
	// TODO: Consider removing this nested user field. We don't use it.
	User                      User                       `json:"-" belongs_to:"user"`
	UserID                    uuid.UUID                  `json:"-" db:"user_id"`
	CreatedAt                 time.Time                  `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time                  `json:"updated_at" db:"updated_at"`
	Status                    string                     `json:"status" db:"status"`
	FriendlyName              string                     `json:"friendly_name,omitempty" db:"friendly_name"`
	Secret                    string                     `json:"-" db:"secret"`
	FactorType                string                     `json:"factor_type" db:"factor_type"`
	Challenge                 []Challenge                `json:"-" has_many:"challenges"`
	Phone                     storage.NullString         `json:"phone" db:"phone"`
	LastChallengedAt          *time.Time                 `json:"last_challenged_at" db:"last_challenged_at"`
	WebAuthnCredential        *MFAWebAuthnCredential     `json:"-" db:"web_authn_credential"`
	WebAuthnAAGUID            *uuid.UUID                 `json:"web_authn_aaguid,omitempty" db:"web_authn_aaguid"`
	LastWebAuthnChallengeData *LastWebAuthnChallengeData `json:"last_webauthn_challenge_data,omitempty" db:"last_webauthn_challenge_data"`
}

type MFAWebAuthnCredential struct {
	webauthn.Credential
}

func (wc *MFAWebAuthnCredential) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

type LastWebAuthnChallengeData struct {
	Challenge          Challenge       `json:"challenge"`
	Type               string          `json:"type"`
	CredentialResponse json.RawMessage `json:"credential_response"`
}

func (lwcd *LastWebAuthnChallengeData) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (lwcd *LastWebAuthnChallengeData) Scan(value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *MFAWebAuthnCredential) Scan(value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle byte and string as a precaution, in postgres driver, json/jsonb should be returned as []byte

func (Factor) TableName() string { _ = "STUB: not implemented"; return "" }

func NewFactor(user *User, friendlyName string, factorType string, state FactorState) *Factor {
	_ = "STUB: not implemented"
	return nil
}

func NewTOTPFactor(user *User, friendlyName string) *Factor { _ = "STUB: not implemented"; return nil }

func NewPhoneFactor(user *User, phone, friendlyName string) *Factor {
	_ = "STUB: not implemented"
	return nil
}

func NewWebAuthnFactor(user *User, friendlyName string) *Factor {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) SetSecret(secret string, encrypt bool, encryptionKeyID, encryptionKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) GetSecret(decryptionKeys map[string]string, encrypt bool, encryptionKeyID string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (f *Factor) SaveWebAuthnCredential(tx *storage.Connection, credential *webauthn.Credential) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) UpdateLastWebAuthnChallenge(tx *storage.Connection, challenge *Challenge, challengeType string, credentialResponse interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func FindFactorByFactorID(conn *storage.Connection, factorID uuid.UUID) (*Factor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteUnverifiedFactors(tx *storage.Connection, user *User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) CreateChallenge(ipAddress string) *Challenge {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) WriteChallengeToDatabase(tx *storage.Connection, challenge *Challenge) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) CreatePhoneChallenge(ipAddress string, otpCode string, encrypt bool, encryptionKeyID, encryptionKey string) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateFriendlyName changes the friendly name
func (f *Factor) UpdateFriendlyName(tx *storage.Connection, friendlyName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) UpdatePhone(tx *storage.Connection, phone string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateStatus modifies the factor status
func (f *Factor) UpdateStatus(tx *storage.Connection, state FactorState) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) DowngradeSessionsToAAL1(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) IsVerified() bool { _ = "STUB: not implemented"; return false }

func (f *Factor) IsUnverified() bool { _ = "STUB: not implemented"; return false }

func (f *Factor) IsPhoneFactor() bool { _ = "STUB: not implemented"; return false }

func (f *Factor) FindChallengeByID(conn *storage.Connection, challengeID uuid.UUID) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteFactorsByUserId(tx *storage.Connection, userId uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteExpiredFactors(tx *storage.Connection, validityDuration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factor) FindLatestUnexpiredChallenge(tx *storage.Connection, expiryDuration float64) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
