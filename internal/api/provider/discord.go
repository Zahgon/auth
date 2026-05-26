package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultDiscordAPIBase = "discord.com"
)

type discordProvider struct {
	*oauth2.Config
	APIPath string
}

type discordUser struct {
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email"`
	ID            string `json:"id"`
	Name          string `json:"username"`
	GlobalName    string `json:"global_name"`
	Verified      bool   `json:"verified"`
}

// NewDiscordProvider creates a Discord account provider.
func NewDiscordProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g discordProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g discordProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g discordProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://discord.com/developers/docs/reference#image-formatting-cdn-endpoints:
// In the case of the Default User Avatar endpoint, the value for
// user_discriminator in the path should be the user's discriminator modulo 5

// https://discord.com/developers/docs/reference#image-formatting:
// "In the case of endpoints that support GIFs, the hash will begin with a_
// if it is available in GIF format."

// To be deprecated
