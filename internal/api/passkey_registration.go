package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
)

// PasskeyRegistrationOptionsParams is the request body for POST /passkeys/registration/options.
type PasskeyRegistrationOptionsParams struct{}

// PasskeyRegistrationOptionsResponse is the response body for POST /passkeys/registration/options.
type PasskeyRegistrationOptionsResponse struct {
	ChallengeID string                                       `json:"challenge_id"`
	Options     *protocol.PublicKeyCredentialCreationOptions `json:"options"`
	ExpiresAt   int64                                        `json:"expires_at"`
}

// PasskeyRegistrationVerifyParams is the request body for POST /passkeys/registration/verify.
type PasskeyRegistrationVerifyParams struct {
	ChallengeID string          `json:"challenge_id"`
	Credential  json.RawMessage `json:"credential"`
}

// PasskeyMetadataResponse is the response body for successful passkey creation.
type PasskeyMetadataResponse struct {
	ID           string    `json:"id"`
	FriendlyName string    `json:"friendly_name,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// PasskeyRegistrationOptions handles POST /passkeys/registration/options.
// Requires authentication. Generates WebAuthn registration options for adding a passkey to an existing account.
func (a *API) PasskeyRegistrationOptions(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Check passkey limit

// Load existing passkeys to build exclusion list

// PasskeyRegistrationVerify handles POST /passkeys/registration/verify.
// Requires authentication. Verifies the WebAuthn credential and creates a passkey for the authenticated user.
func (a *API) PasskeyRegistrationVerify(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Atomically consume the challenge to prevent replay/race conditions

// Parse the credential creation response from the JSON params

// Load existing passkeys for the user adapter

// parseCredentialCreationResponse parses a WebAuthn credential creation response from raw JSON.
func parseCredentialCreationResponse(raw json.RawMessage) (*protocol.ParsedCredentialCreationData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
