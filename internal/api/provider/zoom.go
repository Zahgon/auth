package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultZoomAuthBase = "zoom.us"
	defaultZoomAPIBase  = "api.zoom.us"
)

type zoomProvider struct {
	*oauth2.Config
	APIPath string
}

type zoomUser struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	EmailVerified int    `json:"verified"`
	LoginType     string `json:"login_type"`
	AvatarURL     string `json:"pic_url"`
}

// NewZoomProvider creates a Zoom account provider.
func NewZoomProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g zoomProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g zoomProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g zoomProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A login_type of "100" refers to email-based logins, not oauth.
// A user is verified (type 1) only if they received an email when their profile was created and confirmed the link.
// A zoom user will only be sent an email confirmation link if they signed up using their zoom work email and not oauth.
// See: https://devforum.zoom.us/t/how-to-determine-if-a-zoom-user-actually-owns-their-email-address/44430

// To be deprecated
