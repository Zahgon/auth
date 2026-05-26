package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const defaultSlackOIDCApiBase = "slack.com"

type slackOIDCProvider struct {
	*oauth2.Config
	APIPath string
}

type slackOIDCUser struct {
	ID            string `json:"https://slack.com/user_id"`
	TeamID        string `json:"https://slack.com/team_id"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	AvatarURL     string `json:"picture"`
}

// NewSlackOIDCProvider creates a Slack account provider with Sign in with Slack.
func NewSlackOIDCProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// these are required scopes for slack's OIDC flow
// see https://api.slack.com/authentication/sign-in-with-slack#implementation

func (g slackOIDCProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g slackOIDCProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g slackOIDCProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// email_verified is returned as part of the response
// see: https://api.slack.com/authentication/sign-in-with-slack#response

// To be deprecated
