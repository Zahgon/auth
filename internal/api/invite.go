package api

import (
	"net/http"
)

// InviteParams are the parameters the Signup endpoint accepts
type InviteParams struct {
	Email string                 `json:"email"`
	Data  map[string]interface{} `json:"data"`
}

// Invite is the endpoint for inviting a new user
func (a *API) Invite(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// because params above sets no password, this method
// is not computationally hard so it can be used within
// a database transaction
/* <- isSSOUser */
