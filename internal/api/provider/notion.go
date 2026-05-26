package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultNotionApiBase = "api.notion.com"
	notionApiVersion     = "2021-08-16"
)

type notionProvider struct {
	*oauth2.Config
	APIPath string
}

type notionUser struct {
	Bot struct {
		Owner struct {
			User struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				AvatarURL string `json:"avatar_url"`
				Person    struct {
					Email string `json:"email"`
				} `json:"person"`
			} `json:"user"`
		} `json:"owner"`
	} `json:"bot"`
}

// NewNotionProvider creates a Notion account provider.
func NewNotionProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g notionProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g notionProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (g notionProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"

	// Perform http request, because we need to set the Notion-Version header
	return nil, nil
}

// set headers

// Notion dosen't provide data on if email is verified.

// To be deprecated
