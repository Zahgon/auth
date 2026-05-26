package api

import (
	"context"
	"net/http"
	"regexp"
)

// RefreshTokenGrantParams are the parameters the RefreshTokenGrant method accepts
type RefreshTokenGrantParams struct {
	RefreshToken string `json:"refresh_token"`
}

var legacyRefreshTokenPattern = regexp.MustCompile("^[a-z0-9]{12}$")

func (p *RefreshTokenGrantParams) Validate() error { _ = "STUB: not implemented"; return nil }

// RefreshTokenGrant implements the refresh_token grant type flow
func (a *API) RefreshTokenGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
