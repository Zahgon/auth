package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

const DefaultQRSize = 3

type EnrollFactorParams struct {
	FriendlyName string `json:"friendly_name"`
	FactorType   string `json:"factor_type"`
	Issuer       string `json:"issuer"`
	Phone        string `json:"phone"`
}

type TOTPObject struct {
	QRCode string `json:"qr_code,omitempty"`
	Secret string `json:"secret,omitempty"`
	URI    string `json:"uri,omitempty"`
}

type EnrollFactorResponse struct {
	ID           uuid.UUID   `json:"id"`
	Type         string      `json:"type"`
	FriendlyName string      `json:"friendly_name"`
	TOTP         *TOTPObject `json:"totp,omitempty"`
	Phone        string      `json:"phone,omitempty"`
}

type ChallengeFactorParams struct {
	Channel  string          `json:"channel"`
	WebAuthn *WebAuthnParams `json:"webauthn,omitempty"`
}

type VerifyFactorParams struct {
	ChallengeID uuid.UUID       `json:"challenge_id"`
	Code        string          `json:"code"`
	WebAuthn    *WebAuthnParams `json:"webauthn,omitempty"`
}

type ChallengeFactorResponse struct {
	ID        uuid.UUID              `json:"id"`
	Type      string                 `json:"type"`
	ExpiresAt int64                  `json:"expires_at,omitempty"`
	WebAuthn  *WebAuthnChallengeData `json:"webauthn,omitempty"`
}

type WebAuthnChallengeData struct {
	Type              string      `json:"type"` // "create" or "request"
	CredentialOptions interface{} `json:"credential_options"`
}

type WebAuthnParams struct {
	Type               string          `json:"type"` // "create" or "request"
	CredentialResponse json.RawMessage `json:"credential_response"`
}

type UnenrollFactorResponse struct {
	ID uuid.UUID `json:"id"`
}

func (a *API) getWebAuthnMFA() (*webauthn.WebAuthn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	QRCodeGenerationErrorMessage = "Error generating QR Code"
)

func validateFactors(db *storage.Connection, user *models.User, newFactorName string, config *conf.GlobalConfiguration, session *models.Session) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) enrollPhoneFactor(w http.ResponseWriter, r *http.Request, params *EnrollFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) enrollWebAuthnFactor(w http.ResponseWriter, r *http.Request, params *EnrollFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) enrollTOTPFactor(w http.ResponseWriter, r *http.Request, params *EnrollFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// See: https://css-tricks.com/probably-dont-base64-svg/

func (a *API) EnrollFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) challengePhoneFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// We omit messageID for now, can consider reinstating if there are requests.

func (a *API) challengeTOTPFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) challengeWebAuthnFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Get existing WebAuthn credentials to exclude duplicates

func (a *API) validateChallenge(r *http.Request, db *storage.Connection, factor *models.Factor, challengeID uuid.UUID) (*models.Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *API) ChallengeFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) verifyTOTPFactor(w http.ResponseWriter, r *http.Request, params *VerifyFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Send MFA factor enrolled notification email if enabled and the factor was just verified

// Log the error but don't fail the verification

func (a *API) verifyPhoneFactor(w http.ResponseWriter, r *http.Request, params *VerifyFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Send MFA factor enrolled notification email if enabled and the factor was just verified

// Log the error but don't fail the verification

func (a *API) verifyWebAuthnFactor(w http.ResponseWriter, r *http.Request, params *VerifyFactorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Once the challenge is validated, we consume the challenge

// Challenge verification not needed as the challenge is destroyed on use

// Send MFA factor enrolled notification email if enabled and the factor was just verified

// Log the error but don't fail the verification

func (a *API) VerifyFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) UnenrollFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Send MFA factor unenrolled notification email if enabled

// Log the error but don't fail the unenrollment
