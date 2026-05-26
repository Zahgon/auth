package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// Keycloak
type keycloakProvider struct {
	*oauth2.Config
	Host string
}

type keycloakUser struct {
	Name          string                 `json:"name"`
	Sub           string                 `json:"sub"`
	Email         string                 `json:"email"`
	EmailVerified bool                   `json:"email_verified"`
	RawClaims     map[string]interface{} `json:"-"`
}

func (u *keycloakUser) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Extract known fields

// NewKeycloakProvider creates a Keycloak account provider.
func NewKeycloakProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g keycloakProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g keycloakProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g keycloakProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated
