package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

// ExternalProviderRedirect redirects the request to the oauth provider
func (a *API) ExternalProviderRedirect(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G710

// GetExternalProviderRedirectURL returns the URL to start the oauth flow with the corresponding oauth provider
func (a *API) GetExternalProviderRedirectURL(w http.ResponseWriter, r *http.Request, linkingTargetUser *models.User) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// See https://workos.com/docs/reference/sso/authorize/get

// Handle OAuthClientState for providers that require PKCE on their end

// Build flow state params with all context

// this means that the user is performing manual linking

// Always create flow state for all flows (both PKCE and implicit)
// The flow state ID is used as the state parameter instead of JWT

// Use the flow state ID as the state parameter (UUID format)

// ExternalProviderCallback handles the callback endpoint in the external oauth provider flow
func (a *API) ExternalProviderCallback(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) handleOAuthCallback(r *http.Request) (*OAuthProviderData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// future OAuth1.0 providers will use this method

func (a *API) internalExternalProviderCallback(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// PKCE flow: update flow state with user ID and tokens
// Re-fetch with FOR UPDATE lock inside the transaction to prevent concurrent claims

// Implicit flow: issue tokens directly

// Record login for analytics - only when token is issued (not during pkce authorize)

// PKCE flow: redirect with auth code

// Because not all providers give out a refresh token
// See corresponding OAuth2 spec: <https://www.rfc-editor.org/rfc/rfc6749.html#section-5.1>

// #nosec G710

func (a *API) createAccountFromExternalIdentity(tx *storage.Connection, r *http.Request, userData *provider.UserProvidedData, providerType string, emailOptional bool) (models.AccountLinkingDecision, *models.User, error) {
	_ = "STUB: not implemented"
	return *new(models.AccountLinkingDecision), nil, nil
}

// This is a little bit of a hack. Let me explain: When
// is_sso_user == true, it allows there to be different user
// rows with the same email address. Initially it was added to
// support SSO accounts, but at this point renaming the column
// or adding a new one requires re-indexing the table which is
// expensive and introduces a potentially unnecessary API
// surface change. It therefore set to true for other linking
// domains, not just SSO ones. This enables different linking
// domains to co-exist, such as when using
// GOTRUE_EXPERIMENTAL_PROVIDERS_WITH_OWN_LINKING_DOMAIN="provider_a,provider_b".

// because params above sets no password, this method is not
// computationally hard so it can be used within a database
// transaction

// The user may have other unconfirmed email + password
// combination, phone or oauth identities. These identities
// need to be removed when a new oauth identity is being added
// to prevent pre-account takeover attacks from happening.

// fall through to auto-confirm and issue token

func (a *API) processInvite(r *http.Request, tx *storage.Connection, userData *provider.UserProvidedData, inviteToken, providerType string) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// an account with a previously unconfirmed email + password
// combination or phone may exist. so now that there is an
// OAuth identity bound to this user, and since they have not
// confirmed their email or phone, they are unaware that a
// potentially malicious door exists into their account; thus
// the password and phone needs to be removed.

// confirm because they were able to respond to invite email

func (a *API) loadExternalState(ctx context.Context, r *http.Request, db *storage.Connection) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// loadExternalStateFromUUID loads OAuth state from a flow_state record (new UUID format)
func (a *API) loadExternalStateFromUUID(ctx context.Context, db *storage.Connection, stateID uuid.UUID) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Check expiration

// UserID is nil at creation and set during callback, so non-nil means already consumed.

// Store the entire flow state in context for later use

// Provider returns a Provider interface for the given name.
func (a *API) Provider(ctx context.Context, name string, scopes string) (provider.Provider, conf.OAuthProviderConfiguration, error) {
	_ = "STUB: not implemented"
	return *new(provider.Provider), *new(conf.OAuthProviderConfiguration), nil
}

// Check if this is a custom provider (format: custom:identifier)

// loadCustomProvider loads a custom OAuth or OIDC provider from the database
// identifier should be the full provider name with 'custom:' prefix (e.g., 'custom:github-enterprise')
func (a *API) loadCustomProvider(ctx context.Context, db *storage.Connection, identifier string, scopes string) (provider.Provider, conf.OAuthProviderConfiguration, error) {
	_ = "STUB: not implemented"
	return *new(provider.Provider), *new(conf.OAuthProviderConfiguration), nil
}

// Build the redirect URL

// Parse scopes (space-separated per RFC 6749)

// Find the custom provider by identifier (which now includes 'custom:' prefix)

// Check if provider is enabled

// Use provider scopes if not overridden

// Decrypt client secret for runtime use

// Handle based on provider type

// OAuth2 provider

// Create custom OAuth provider instance

// Build provider configuration

// OIDC provider

// Create custom OIDC provider instance
// oidc.NewProvider() will automatically fetch discovery document

// Build provider configuration

func redirectErrors(handler apiHandler, w http.ResponseWriter, r *http.Request, u *url.URL) {
	_ = "STUB: not implemented"
	return
}

// TODO: deprecate returning error details in the query fragment

// Add Supabase Auth identifier to help clients distinguish Supabase Auth redirects

// #nosec G710

func getErrorQueryString(err error, errorID string, log logrus.FieldLogger, q url.Values) *url.Values {
	_ = "STUB: not implemented"
	return nil
}

// this will get us the stack trace too

// Provide better error messages for certain user-triggered Postgres errors.

func (a *API) getExternalRedirectURL(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func (a *API) createNewIdentity(tx *storage.Connection, user *models.User, providerType string, identityData map[string]interface{}) (*models.Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
