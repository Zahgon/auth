package provider

import (
	"context"

	"github.com/mrjones/oauth"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultTwitterAPIBase = "api.twitter.com"
	requestURL            = "/oauth/request_token"
	authenticateURL       = "/oauth/authenticate"
	tokenURL              = "/oauth/access_token" //#nosec G101 -- Not a secret value.
	endpointProfile       = "/1.1/account/verify_credentials.json"
)

// TwitterProvider stores the custom config for twitter provider
type TwitterProvider struct {
	ClientKey     string
	Secret        string
	CallbackURL   string
	AuthURL       string
	RequestToken  *oauth.RequestToken
	OauthVerifier string
	Consumer      *oauth.Consumer
	UserInfoURL   string
}

type twitterUser struct {
	UserName  string `json:"screen_name"`
	Name      string `json:"name"`
	AvatarURL string `json:"profile_image_url_https"`
	Email     string `json:"email"`
	ID        string `json:"id_str"`
}

// NewTwitterProvider creates a Twitter account provider.
func NewTwitterProvider(ext conf.OAuthProviderConfiguration, scopes string) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// GetOAuthToken is a stub method for OAuthProvider interface, unused in OAuth1.0 protocol
func (t TwitterProvider) GetOAuthToken(_ context.Context, _ string, _ ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TwitterProvider) RequiresPKCE() bool {
	_ = "STUB: not implemented"

	// GetUserData is a stub method for OAuthProvider interface, unused in OAuth1.0 protocol
	return false
}

func (t TwitterProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchUserData retrieves the user's data from the twitter provider
func (t TwitterProvider) FetchUserData(ctx context.Context, tok *oauth.AccessToken) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To be deprecated

// AuthCodeURL fetches the request token from the twitter provider
func (t *TwitterProvider) AuthCodeURL(state string, args ...oauth2.AuthCodeOption) string {
	_ = "STUB: not implemented"
	// we do nothing with the state here as the state is passed in the requestURL step
	return ""
}

func newConsumer(provider *TwitterProvider, authHost string) *oauth.Consumer {
	_ = "STUB: not implemented"
	return nil
}

// Marshal encodes the twitter request token
func (t TwitterProvider) Marshal() string { _ = "STUB: not implemented"; return "" }

// Unmarshal decodes the twitter request token
func (t TwitterProvider) Unmarshal(data string) (*oauth.RequestToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
