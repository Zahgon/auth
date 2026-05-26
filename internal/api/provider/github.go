package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// Github

const (
	defaultGitHubAuthBase = "github.com"
	defaultGitHubAPIBase  = "api.github.com"
)

type githubProvider struct {
	*oauth2.Config
	APIHost string
}

type githubUser struct {
	ID        int    `json:"id"`
	UserName  string `json:"login"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type githubUserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

// NewGithubProvider creates a Github account provider.
func NewGithubProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g githubProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g githubProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g githubProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated
