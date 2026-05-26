package provider

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultVercelMarketplaceAPIBase = "api.vercel.com"
	IssuerVercelMarketplace         = "https://marketplace.vercel.com"
)

type vercelMarketplaceProvider struct {
	*oauth2.Config
	oidc    *oidc.Provider
	APIPath string
}

// NewVercelMarketplaceProvider creates a VercelMarketplace account provider via OIDC.
func NewVercelMarketplaceProvider(ctx context.Context, ext conf.OAuthProviderConfiguration, scopes string, cache *OIDCProviderCache) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g vercelMarketplaceProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g vercelMarketplaceProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g vercelMarketplaceProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
