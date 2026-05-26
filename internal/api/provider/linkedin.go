package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultLinkedinAPIBase = "api.linkedin.com"
)

type linkedinProvider struct {
	*oauth2.Config
	APIPath      string
	UserInfoURL  string
	UserEmailUrl string
}

// See https://docs.microsoft.com/en-us/linkedin/consumer/integrations/self-serve/sign-in-with-linkedin?context=linkedin/consumer/context
// for retrieving a member's profile. This requires the r_liteprofile scope.
type linkedinUser struct {
	ID        string       `json:"id"`
	FirstName linkedinName `json:"firstName"`
	LastName  linkedinName `json:"lastName"`
	AvatarURL struct {
		DisplayImage struct {
			Elements []struct {
				Identifiers []struct {
					Identifier string `json:"identifier"`
				} `json:"identifiers"`
			} `json:"elements"`
		} `json:"displayImage~"`
	} `json:"profilePicture"`
}

func (u *linkedinUser) getAvatarUrl() string { _ = "STUB: not implemented"; return "" }

type linkedinName struct {
	Localized       interface{}    `json:"localized"`
	PreferredLocale linkedinLocale `json:"preferredLocale"`
}

type linkedinLocale struct {
	Country  string `json:"country"`
	Language string `json:"language"`
}

// See https://docs.microsoft.com/en-us/linkedin/consumer/integrations/self-serve/sign-in-with-linkedin?context=linkedin/consumer/context#retrieving-member-email-address
// for retrieving a member email address. This requires the r_email_address scope.
type linkedinElements struct {
	Elements []struct {
		Handle      string `json:"handle"`
		HandleTilde struct {
			EmailAddress string `json:"emailAddress"`
		} `json:"handle~"`
	} `json:"elements"`
}

// NewLinkedinProvider creates a Linkedin account provider.
func NewLinkedinProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

func (g linkedinProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g linkedinProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func GetName(name linkedinName) string { _ = "STUB: not implemented"; return "" }

func (g linkedinProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: Use primary contact api for handling phone numbers

// linkedin only returns the primary email which is verified for the r_emailaddress scope.

// To be deprecated
