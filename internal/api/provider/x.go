package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// X (formerly Twitter) API v2 OAuth 2.0 endpoints
// See: https://developer.x.com/en/docs/authentication/oauth-2-0/authorization-code
const (
	defaultXAuthBase = "x.com"
	defaultXAPIBase  = "api.x.com"
)

type xProvider struct {
	*oauth2.Config
	APIHost string
}

// xUser represents the user object from X API v2
// See: https://developer.x.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
type xUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ConfirmedEmail  string `json:"confirmed_email"`
	ProfileImageURL string `json:"profile_image_url"`
	URL             string `json:"url"`
	CreatedAt       string `json:"created_at"`
}

// xUserResponse is the wrapper for the X API v2 response
type xUserResponse struct {
	Data xUser `json:"data"`
}

// NewXProvider creates an X (formerly Twitter) v2 OAuth 2.0 provider.
// This uses OAuth 2.0 with PKCE instead of OAuth 1.0a.
// See: https://developer.x.com/en/docs/authentication/oauth-2-0/authorization-code
func NewXProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// Default scopes for user authentication
// users.email: Access to the user's email address (confirmed_email field)
// users.read: Read user profile information
// tweet.read: Required scope for OAuth 2.0 user context even if not accessing tweets
// offline.access: Get refresh tokens for long-lived access
// See: https://developer.x.com/en/docs/authentication/oauth-2-0/authorization-code
// and: https://docs.x.com/fundamentals/authentication/guides/v2-authentication-mapping

func (x xProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x xProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (x xProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil,

		// See: https://developer.x.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
		nil
}

// Custom claims for X specific data

// To be deprecated

// X returns only confirmed emails
