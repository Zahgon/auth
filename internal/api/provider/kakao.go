package provider

import (
	"context"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultKakaoAuthBase = "kauth.kakao.com"
	defaultKakaoAPIBase  = "kapi.kakao.com"
	IssuerKakao          = "https://kauth.kakao.com"
)

type kakaoProvider struct {
	*oauth2.Config
	APIHost string
}

type kakaoUser struct {
	ID      int `json:"id"`
	Account struct {
		Profile struct {
			Name            string `json:"nickname"`
			ProfileImageURL string `json:"profile_image_url"`
		} `json:"profile"`
		Email         string `json:"email"`
		EmailValid    bool   `json:"is_email_valid"`
		EmailVerified bool   `json:"is_email_verified"`
	} `json:"kakao_account"`
}

func (p kakaoProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p kakaoProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func (p kakaoProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated

func NewKakaoProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}
