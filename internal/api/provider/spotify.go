package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultSpotifyAPIBase  = "api.spotify.com/v1"   // Used to get user data
	defaultSpotifyAuthBase = "accounts.spotify.com" // Used for OAuth flow
)

type spotifyProvider struct {
	*oauth2.Config
	APIPath string
}

type spotifyUser struct {
	DisplayName string             `json:"display_name"`
	Avatars     []spotifyUserImage `json:"images"`
	Email       string             `json:"email"`
	ID          string             `json:"id"`
}

type spotifyUserImage struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// NewSpotifyProvider creates a Spotify account provider.
func NewSpotifyProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g spotifyProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g spotifyProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g spotifyProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Spotify dosen't provide data on whether the user's email is verified.
// https://developer.spotify.com/documentation/web-api/reference/get-current-users-profile

// Spotify returns a list of avatars, we want to use the largest one

// To be deprecated
