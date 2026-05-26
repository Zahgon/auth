package api

import (
	"net/http"
)

// AdminPasskeyList handles GET /admin/users/{user_id}/passkeys.
// Requires admin credentials. Returns all passkeys for the specified user.
func (a *API) AdminPasskeyList(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// AdminPasskeyDelete handles DELETE /admin/users/{user_id}/passkeys/{passkey_id}.
// Requires admin credentials. Deletes the specified passkey.
func (a *API) AdminPasskeyDelete(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
