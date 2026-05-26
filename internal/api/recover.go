package api

import (
	"net/http"
)

// RecoverParams holds the parameters for a password recovery request
type RecoverParams struct {
	Email               string `json:"email"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
}

func (p *RecoverParams) Validate(a *API) error { _ = "STUB: not implemented"; return nil }

// Recover sends a recovery email
func (a *API) Recover(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
