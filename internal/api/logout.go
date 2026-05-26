package api

import (
	"net/http"
)

type LogoutBehavior string

const (
	LogoutGlobal LogoutBehavior = "global"
	LogoutLocal  LogoutBehavior = "local"
	LogoutOthers LogoutBehavior = "others"
)

// Logout is the endpoint for logging out a user and thereby revoking any refresh tokens
func (a *API) Logout(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// default mode, log out everywhere
