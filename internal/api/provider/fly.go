package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultFlyAPIBase = "api.fly.io"
)

type flyProvider struct {
	*oauth2.Config
	APIPath string
}

type flyUser struct {
	ResourceOwnerID string `json:"resource_owner_id"`
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	Organizations   []struct {
		ID   string `json:"id"`
		Role string `json:"role"`
	} `json:"organizations"`
	Scope       []string          `json:"scope"`
	Application map[string]string `json:"application"`
	ExpiresIn   int               `json:"expires_in"`
	CreatedAt   int               `json:"created_at"`
}

// NewFlyProvider creates a Fly oauth provider.
func NewFlyProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// Fly only provides the "read" scope.
// https://fly.io/docs/reference/extensions_api/#single-sign-on-flow

func (p flyProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p flyProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p flyProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
