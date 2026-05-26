package provider

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultLinkedinOIDCAPIBase = "api.linkedin.com"
	IssuerLinkedin             = "https://www.linkedin.com/oauth"
)

type linkedinOIDCProvider struct {
	*oauth2.Config
	oidc    *oidc.Provider
	APIPath string
}

// NewLinkedinOIDCProvider creates a Linkedin account provider via OIDC.
func NewLinkedinOIDCProvider(ctx context.Context, ext conf.OAuthProviderConfiguration, scopes string, cache *OIDCProviderCache) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g linkedinOIDCProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g linkedinOIDCProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g linkedinOIDCProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
