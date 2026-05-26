package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// Figma
// Reference: https://www.figma.com/developers/api#oauth2

const (
	defaultFigmaAuthBase = "www.figma.com"
	defaultFigmaAPIBase  = "api.figma.com"
)

type figmaProvider struct {
	*oauth2.Config
	APIHost string
}

type figmaUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"handle"`
	AvatarURL string `json:"img_url"`
}

// NewFigmaProvider creates a Figma account provider.
func NewFigmaProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (p figmaProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p figmaProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p figmaProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated
