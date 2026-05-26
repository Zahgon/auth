package provider

import (
	"context"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
)

type ParseIDTokenOptions struct {
	SkipAccessTokenCheck bool
	AccessToken          string
}

// OverrideVerifiers can be used to set a custom verifier for an OIDC provider
// (identified by the provider's Endpoint().AuthURL string). Should only be
// used in tests.
var OverrideVerifiers = make(map[string]func(context.Context, *oidc.Config) *oidc.IDTokenVerifier)

// OverrideClock can be used to set a custom clock function to be used when
// parsing ID tokens. Should only be used in tests.
var OverrideClock func() time.Time

func ParseIDToken(ctx context.Context, provider *oidc.Provider, config *oidc.Config, idToken string, options ParseIDTokenOptions) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// aud claim check to be performed by other flows

// Handle only Facebook Limited Login JWT, NOT Facebook Access Token

func parseGoogleIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// To be deprecated

// FacebookIDTokenClaims represents the claims in a Facebook Limited Login ID token
type FacebookIDTokenClaims struct {
	jwt.RegisteredClaims
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Nonce   string `json:"nonce,omitempty"`
}

func parseFacebookIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Facebook Limited Login emails are always verified

type AppleIDTokenClaims struct {
	jwt.RegisteredClaims

	Email string `json:"email"`

	AuthTime       *float64        `json:"auth_time"`
	IsPrivateEmail *IsPrivateEmail `json:"is_private_email"`

	TransferSub string `json:"transfer_sub"`
}

func parseAppleIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type LinkedinIDTokenClaims struct {
	jwt.RegisteredClaims

	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Locale        string `json:"locale"`
	Picture       string `json:"picture"`
}

func parseLinkedinIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type AzureIDTokenClaims struct {
	jwt.RegisteredClaims

	Email                              string `json:"email"`
	Name                               string `json:"name"`
	PreferredUsername                  string `json:"preferred_username"`
	XMicrosoftEmailDomainOwnerVerified any    `json:"xms_edov"`
}

func (c *AzureIDTokenClaims) IsEmailVerified() bool { _ = "STUB: not implemented"; return false }

// If xms_edov is not set, and an email is present or xms_edov is true,
// only then is the email regarded as verified.
// https://learn.microsoft.com/en-us/azure/active-directory/develop/migrate-off-email-claim-authorization#using-the-xms_edov-optional-claim-to-determine-email-verification-status-and-migrate-users

// An email is provided, but xms_edov is not -- probably not
// configured, so we must assume the email is verified as Azure
// will only send out a potentially unverified email address in
// single-tenanat apps.

// Azure can't be trusted with how they encode the xms_edov
// claim. Sometimes it's "xms_edov": "1", sometimes "xms_edov": true.

// removeAzureClaimsFromCustomClaims contains the list of claims to be removed
// from the CustomClaims map. See:
// https://learn.microsoft.com/en-us/azure/active-directory/develop/id-token-claims-reference
var removeAzureClaimsFromCustomClaims = []string{
	"aud",
	"iss",
	"iat",
	"nbf",
	"exp",
	"c_hash",
	"at_hash",
	"aio",
	"nonce",
	"rh",
	"uti",
	"jti",
	"ver",
	"sub",
	"name",
	"preferred_username",
}

func parseAzureIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type KakaoIDTokenClaims struct {
	jwt.RegisteredClaims

	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Picture  string `json:"picture"`
}

func parseKakaoIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type VercelMarketplaceIDTokenClaims struct {
	jwt.RegisteredClaims

	GlobalUserID  string `json:"global_user_id"`
	UserEmail     string `json:"user_email"`
	UserName      string `json:"user_name"`
	UserAvatarUrl string `json:"user_avatar_url"`
}

func parseVercelMarketplaceIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseGenericIDToken(token *oidc.IDToken) (*oidc.IDToken, *UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
