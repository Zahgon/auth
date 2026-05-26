package provider

import (
	"context"
	"regexp"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/oauth2"
)

const IssuerAzureCommon = "https://login.microsoftonline.com/common/v2.0"
const IssuerAzureOrganizations = "https://login.microsoftonline.com/organizations/v2.0"

// IssuerAzureMicrosoft is the OIDC issuer for microsoft.com accounts:
// https://learn.microsoft.com/en-us/azure/active-directory/develop/id-token-claims-reference#payload-claims
const IssuerAzureMicrosoft = "https://login.microsoftonline.com/9188040d-6c67-4c5b-b112-36a304b66dad/v2.0"

const (
	defaultAzureAuthBase = "login.microsoftonline.com/common"
)

type azureProvider struct {
	*oauth2.Config

	// ExpectedIssuer contains the OIDC issuer that should be expected when
	// the authorize flow completes. For example, when using the "common"
	// endpoint the authorization flow will end with an ID token that
	// contains any issuer. In this case, ExpectedIssuer is an empty
	// string, because any issuer is allowed. But if a developer sets up a
	// tenant-specific authorization endpoint, then we must ensure that the
	// ID token received is issued by that specific issuer, and so
	// ExpectedIssuer contains the issuer URL of that tenant.
	ExpectedIssuer string

	cache *OIDCProviderCache
}

var azureIssuerRegexp = regexp.MustCompile("^https://login[.]microsoftonline[.]com/([^/]+)/v2[.]0/?$")
var azureCIAMIssuerRegexp = regexp.MustCompile("^https://[a-z0-9-]+[.]ciamlogin[.]com/([^/]+)/v2[.]0/?$")

func IsAzureIssuer(issuer string) bool { _ = "STUB: not implemented"; return false }

func IsAzureCIAMIssuer(issuer string) bool { _ = "STUB: not implemented"; return false }

// NewAzureProvider creates a Azure account provider.
func NewAzureProvider(ext conf.OAuthProviderConfiguration, scopes string, cache *OIDCProviderCache) (OAuthProvider, error) {
	_ = "STUB: not implemented"
	return *new(OAuthProvider), nil
}

// in tests, the URL is a local server which should not
// be the expected issuer
// also, IssuerAzure (common) never actually issues any
// ID tokens so it needs to be ignored

func (g azureProvider) GetOAuthToken(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g azureProvider) RequiresPKCE() bool { _ = "STUB: not implemented"; return false }

func DetectAzureIDTokenIssuer(ctx context.Context, idToken string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g azureProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allow basic Azure issuers, except when the expected issuer
// is configured to be the Azure CIAM issuer, allow CIAM
// issuers to pass.

// Since ExpectedIssuer was set, then the developer had
// setup GoTrue to use the tenant-specific
// authorization endpoint, which in-turn means that
// only those tenant's ID tokens will be accepted.

// Only ID tokens supported, UserInfo endpoint has a history of being less secure.
