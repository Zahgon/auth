package provider

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

type googleUser struct {
	ID            string `json:"id"`
	Subject       string `json:"sub"`
	Issuer        string `json:"iss"`
	Name          string `json:"name"`
	AvatarURL     string `json:"picture"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	EmailVerified bool   `json:"email_verified"`
	HostedDomain  string `json:"hd"`
}

func (u googleUser) IsEmailVerified() bool { _ = "STUB: not implemented"; return false }

const IssuerGoogle = "https://accounts.google.com"

var internalIssuerGoogle = IssuerGoogle

type googleProvider struct {
	*oauth2.Config

	oidc *oidc.Provider
}

// NewGoogleProvider creates a Google OAuth2 identity provider.
func NewGoogleProvider(ctx context.Context, ext conf.OAuthProviderConfiguration, scopes string, cache *OIDCProviderCache) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g googleProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g googleProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

const UserInfoEndpointGoogle = "https://www.googleapis.com/userinfo/v2/me"

var internalUserInfoEndpointGoogle = UserInfoEndpointGoogle

func (g googleProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This whole section offers legacy support in case the Google OAuth2
// flow does not return an ID Token for the user, which appears to
// always be the case.

// To be deprecated

// ResetGoogleProvider should only be used in tests!
func ResetGoogleProvider() { _ = "STUB: not implemented"; return }

// OverrideGoogleProvider should only be used in tests!
func OverrideGoogleProvider(issuer, userInfo string) { _ = "STUB: not implemented"; return }
