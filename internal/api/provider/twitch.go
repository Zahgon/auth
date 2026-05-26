package provider

import (
	"context"
	"time"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

// Twitch

const (
	defaultTwitchAuthBase = "id.twitch.tv"
	defaultTwitchAPIBase  = "api.twitch.tv"
)

type twitchProvider struct {
	*oauth2.Config
	APIHost string
}

type twitchUsers struct {
	Data []struct {
		ID              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageURL string    `json:"profile_image_url"`
		OfflineImageURL string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		Email           string    `json:"email"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"data"`
}

// NewTwitchProvider creates a Twitch account provider.
func NewTwitchProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (t twitchProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t twitchProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (t twitchProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"

	// Perform http request, because we neeed to set the Client-Id header
	return nil, nil
}

// set headers

// To be deprecated
