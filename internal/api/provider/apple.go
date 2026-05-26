package provider

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const DefaultAppleIssuer = "https://appleid.apple.com"
const OtherAppleIssuer = "https://account.apple.com"

func IsAppleIssuer(issuer string) bool { _ = "STUB: not implemented"; return false }

func DetectAppleIDTokenIssuer(ctx context.Context, idToken string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AppleProvider stores the custom config for apple provider
type AppleProvider struct {
	*oauth2.Config
	oidc *oidc.Provider
}

type IsPrivateEmail bool

// Apple returns an is_private_email field that could be a string or boolean value so we need to implement a custom unmarshaler
// https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/authenticating_users_with_sign_in_with_apple
func (b *IsPrivateEmail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// ignore the error and try to unmarshal as a string

type appleName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type appleUser struct {
	Name  appleName `json:"name"`
	Email string    `json:"email"`
}

// NewAppleProvider creates a Apple account provider.
func NewAppleProvider(ctx context.Context, ext conf.OAuthProviderConfiguration, cache *OIDCProviderCache) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// GetOAuthToken returns the apple provider access token
func (p AppleProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p AppleProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p AppleProvider) AuthCodeURL(state string, args ...oauth2.AuthCodeOption) string {
	_ = "STUB: not implemented"
	return ""
}

// GetUserData returns the user data fetched from the apple provider
func (p AppleProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apple returns user data only the first time

// ParseUser parses the apple user's info
func (p AppleProvider) ParseUser(data string, userData *UserProvidedData) error {
	_ = "STUB: not implemented"
	return nil
}
