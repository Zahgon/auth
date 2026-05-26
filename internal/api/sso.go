package api

import (
	"net/http"

	"github.com/gofrs/uuid"
)

type SingleSignOnParams struct {
	ProviderID          uuid.UUID `json:"provider_id"`
	Domain              string    `json:"domain"`
	RedirectTo          string    `json:"redirect_to"`
	SkipHTTPRedirect    *bool     `json:"skip_http_redirect"`
	CodeChallenge       string    `json:"code_challenge"`
	CodeChallengeMethod string    `json:"code_challenge_method"`
}

type SingleSignOnResponse struct {
	URL string `json:"url"`
}

func (p *SingleSignOnParams) validate() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// SingleSignOn handles the single-sign-on flow for a provided SSO domain or provider.
func (a *API) SingleSignOn(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Always create flow state for all SSO flows

/* <- idpInitiated */

// Some IdPs do not support the use of the `persistent` NameID format,
// and require a different format to be sent to work.
