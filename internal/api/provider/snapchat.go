package provider

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const IssuerSnapchat = "https://accounts.snapchat.com/accounts/oauth2/token"

const (
	defaultSnapchatAuthBase  = "accounts.snapchat.com"
	defaultSnapchatTokenBase = "accounts.snapchat.com"
	defaultSnapchatAPIBase   = "kit.snapchat.com"
)

type snapchatProvider struct {
	*oauth2.Config
	ProfileURL string
}

type snapchatUser struct {
	Data struct {
		Me struct {
			ExternalID  string `json:"externalId"`
			DisplayName string `json:"displayName"`
			Bitmoji     struct {
				Avatar string `json:"avatar"`
			} `json:"bitmoji"`
			Email string `json:"email"`
		} `json:"me"`
	} `json:"data"`
}

// NewSnapchatProvider creates a Snapchat account provider.
func NewSnapchatProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (p snapchatProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p snapchatProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p snapchatProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"

	// Create a URL with the GraphQL query parameter
	return nil, nil
}

// Add the GraphQL query parameter

// To be deprecated

func parseSnapchatIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
