package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultWorkOSAPIBase = "api.workos.com"
)

type workosProvider struct {
	*oauth2.Config
	APIPath string
}

// See https://workos.com/docs/reference/sso/profile.
type workosUser struct {
	ID             string                 `mapstructure:"id"`
	ConnectionID   string                 `mapstructure:"connection_id"`
	OrganizationID string                 `mapstructure:"organization_id"`
	ConnectionType string                 `mapstructure:"connection_type"`
	Email          string                 `mapstructure:"email"`
	FirstName      string                 `mapstructure:"first_name"`
	LastName       string                 `mapstructure:"last_name"`
	Object         string                 `mapstructure:"object"`
	IdpID          string                 `mapstructure:"idp_id"`
	RawAttributes  map[string]interface{} `mapstructure:"raw_attributes"`
}

// NewWorkOSProvider creates a WorkOS account provider.
func NewWorkOSProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g workosProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g workosProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g workosProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WorkOS API returns the user's profile data along with the OAuth2 token, so
// we can just convert from `map[string]interface{}` to `workosUser` without
// an additional network request.

// To be deprecated
