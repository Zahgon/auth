package api

import (
	"net/http"
)

// MagicLinkParams holds the parameters for a magic link request
type MagicLinkParams struct {
	Email               string                 `json:"email"`
	Data                map[string]interface{} `json:"data"`
	CodeChallengeMethod string                 `json:"code_challenge_method"`
	CodeChallenge       string                 `json:"code_challenge"`
}

func (p *MagicLinkParams) Validate(a *API) error { _ = "STUB: not implemented"; return nil }

// MagicLink sends a recovery email
func (a *API) MagicLink(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// User either doesn't exist or hasn't completed the signup process.
// Sign them up with temporary password.

// SignupParams must always be marshallable

// signups are autoconfirmed, send magic link after signup

// SignupParams must always be marshallable

// otherwise confirmation email already contains 'magic link'

// responseStub only implement http responsewriter for ignoring
// incoming data from methods where it passed
type responseStub struct {
}

func (rw *responseStub) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (rw *responseStub) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rw *responseStub) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }
