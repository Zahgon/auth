package api

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/supabase/auth/internal/models"
)

// getPasskeyWebAuthn creates a *webauthn.WebAuthn instance from the shared server-side WebAuthn configuration.
func (a *API) getPasskeyWebAuthn() (*webauthn.WebAuthn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// required to support discoverable credentials

// TODO(fm): webAuthnUser is a thin adapter that wraps a *models.User and returns passkey
// credentials (from the webauthn_credentials table) instead of MFA factor
// credentials. This is necessary because the existing User.WebAuthnCredentials()
// method returns MFA WebAuthn factor credentials until they are consolidated.
type webAuthnUser struct {
	user        *models.User
	credentials []webauthn.Credential
}

func newWebAuthnUser(user *models.User, passkeyCredentials []*models.WebAuthnCredential) *webAuthnUser {
	_ = "STUB: not implemented"
	return nil
}

func (u *webAuthnUser) WebAuthnID() []byte { _ = "STUB: not implemented"; return nil }

func (u *webAuthnUser) WebAuthnName() string { _ = "STUB: not implemented"; return "" }

func (u *webAuthnUser) WebAuthnDisplayName() string { _ = "STUB: not implemented"; return "" }

func (u *webAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	_ = "STUB: not implemented"
	return nil
}
