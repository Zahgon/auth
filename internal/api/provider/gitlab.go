package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// Gitlab

const defaultGitLabAuthBase = "gitlab.com"

type gitlabProvider struct {
	*oauth2.Config
	Host string
}

type gitlabUser struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url"`
	ConfirmedAt string `json:"confirmed_at"`
	ID          int    `json:"id"`
}

type gitlabUserEmail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// NewGitlabProvider creates a Gitlab account provider.
func NewGitlabProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g gitlabProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g gitlabProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g gitlabProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// additional emails from GitLab don't return confirm status

// To be deprecated
