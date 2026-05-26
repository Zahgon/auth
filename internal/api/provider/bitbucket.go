package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultBitbucketAuthBase = "bitbucket.org"
	defaultBitbucketAPIBase  = "api.bitbucket.org"
)

type bitbucketProvider struct {
	*oauth2.Config
	APIPath string
}

type bitbucketUser struct {
	Name   string `json:"display_name"`
	ID     string `json:"uuid"`
	Avatar struct {
		Href string `json:"href"`
	} `json:"avatar"`
}

type bitbucketEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"is_primary"`
	Verified bool   `json:"is_confirmed"`
}

type bitbucketEmails struct {
	Values []bitbucketEmail `json:"values"`
}

// NewBitbucketProvider creates a Bitbucket account provider.
func NewBitbucketProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g bitbucketProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g bitbucketProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g bitbucketProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated
