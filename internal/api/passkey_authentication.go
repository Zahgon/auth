package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
)

// PasskeyAuthenticationOptionsResponse is the response body for POST /passkeys/authentication/options.
type PasskeyAuthenticationOptionsResponse struct {
	ChallengeID string                                      `json:"challenge_id"`
	Options     *protocol.PublicKeyCredentialRequestOptions `json:"options"`
	ExpiresAt   int64                                       `json:"expires_at"`
}

// PasskeyAuthenticationVerifyParams is the request body for POST /passkeys/authentication/verify.
type PasskeyAuthenticationVerifyParams struct {
	ChallengeID string          `json:"challenge_id"`
	Credential  json.RawMessage `json:"credential"`
}

// PasskeyAuthenticationOptions handles POST /passkeys/authentication/options.
// Generates WebAuthn authentication options for discoverable credential login.
func (a *API) PasskeyAuthenticationOptions(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Discoverable flow: empty allowCredentials, no user binding

// no user_id for discoverable flow

// PasskeyAuthenticationVerify handles POST /passkeys/authentication/verify.
// Validates the WebAuthn assertion and issues tokens for discoverable credential login.
func (a *API) PasskeyAuthenticationVerify(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Discoverable login: resolve user from userHandle in the assertion

// Look up the authenticated user from the validated assertion's userHandle

// Find the matching WebAuthnCredential record to update

// parseCredentialAssertionResponse parses a WebAuthn credential assertion response from raw JSON.
func parseCredentialAssertionResponse(raw json.RawMessage) (*protocol.ParsedCredentialAssertionData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
