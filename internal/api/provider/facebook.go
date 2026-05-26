package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const IssuerFacebook = "https://www.facebook.com"

const (
	defaultFacebookAuthBase  = "www.facebook.com"
	defaultFacebookTokenBase = "graph.facebook.com" //#nosec G101 -- Not a secret value.
	defaultFacebookAPIBase   = "graph.facebook.com"
)

type facebookProvider struct {
	*oauth2.Config
	ProfileURL string
}

type facebookUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Alias     string `json:"name"`
	Avatar    struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}

// NewFacebookProvider creates a Facebook account provider.
func NewFacebookProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (p facebookProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p facebookProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p facebookProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated
