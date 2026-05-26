package api

import (
	"net/http"
	"time"

	"github.com/supabase/auth/internal/models"
)

// PasskeyListItem is the response shape for a single passkey in the list and management endpoints.
type PasskeyListItem struct {
	ID           string     `json:"id"`
	FriendlyName string     `json:"friendly_name,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	LastUsedAt   *time.Time `json:"last_used_at,omitempty"`
}

// PasskeyUpdateParams is the request body for PATCH /passkeys/{passkey_id}.
type PasskeyUpdateParams struct {
	FriendlyName string `json:"friendly_name"`
}

// TODO(fm): we should not allow any of the following operations on credentials used for
// MFA webauthn factors — in particular, the deletion operation.

// PasskeyList handles GET /passkeys/.
// Requires authentication. Returns all passkeys for the authenticated user.
func (a *API) PasskeyList(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// PasskeyUpdate handles PATCH /passkeys/{passkey_id}.
// Requires authentication. Updates the friendly_name of a passkey owned by the authenticated user.
func (a *API) PasskeyUpdate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// PasskeyDelete handles DELETE /passkeys/{passkey_id}.
// Requires authentication. Deletes a passkey owned by the authenticated user.
func (a *API) PasskeyDelete(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func toPasskeyListItem(cred *models.WebAuthnCredential) PasskeyListItem {
	_ = "STUB: not implemented"
	return *new(PasskeyListItem)
}
