package api

import (
	"context"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

// IdTokenGrantParams are the parameters the IdTokenGrant method accepts
type IdTokenGrantParams struct {
	IdToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	Nonce        string `json:"nonce"`
	Provider     string `json:"provider"`
	ClientID     string `json:"client_id"`
	Issuer       string `json:"issuer"`
	LinkIdentity bool   `json:"link_identity"`
}

func (p *IdTokenGrantParams) getProvider(ctx context.Context, db *storage.Connection, config *conf.GlobalConfiguration, r *http.Request, cache *provider.OIDCProviderCache) (*oidc.Provider, bool, string, []string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, "", nil, false, nil
}

// Facebook (Limited Login) nonce check is not supported

// Custom OIDC provider - identifier already includes 'custom:' prefix

// Ensure it's an OIDC provider

// already checked above

// IdTokenGrant implements the id_token grant type flow
func (a *API) IdTokenGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// set it so linkIdentityToUser works below

// verify nonce to mitigate replay attacks
