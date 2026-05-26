package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const defaultSlackApiBase = "slack.com"

type slackProvider struct {
	*oauth2.Config
	APIPath string
}

type slackUser struct {
	ID        string `json:"https://slack.com/user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"picture"`
	TeamID    string `json:"https://slack.com/team_id"`
}

// NewSlackProvider creates a Slack account provider with Legacy Slack OAuth.
func NewSlackProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g slackProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g slackProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g slackProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Slack doesn't provide data on if email is verified.

// To be deprecated
