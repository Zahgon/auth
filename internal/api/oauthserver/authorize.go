package oauthserver

import (
	"net/http"

	"github.com/supabase/auth/internal/models"
)

// AuthorizeParams represents the parameters for an OAuth authorization request
type AuthorizeParams struct {
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
	State        string `json:"state"`

	// Resource Resource Indicator per RFC8707
	Resource            string `json:"resource"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
	Nonce               string `json:"nonce"` // OIDC nonce parameter
}

// AuthorizationDetailsResponse represents the response for getting authorization details
type AuthorizationDetailsResponse struct {
	AuthorizationID string                `json:"authorization_id"`
	RedirectURI     string                `json:"redirect_uri,omitempty"`
	Client          ClientDetailsResponse `json:"client,omitempty"`
	User            UserDetailsResponse   `json:"user,omitempty"`
	Scope           string                `json:"scope,omitempty"`
}

// ClientDetailsResponse represents client details in authorization response
type ClientDetailsResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	URI     string `json:"uri,omitempty"`
	LogoURI string `json:"logo_uri,omitempty"`
}

// UserDetailsResponse represents user details in authorization response
type UserDetailsResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// ConsentRequest represents a consent decision request
type ConsentRequest struct {
	Action OAuthServerConsentAction `json:"action"`
}

// ConsentResponse represents the response after processing consent
type ConsentResponse struct {
	RedirectURL string `json:"redirect_url,omitempty"`
}

type OAuthServerConsentAction string

const (
	OAuthServerConsentActionApprove OAuthServerConsentAction = "approve"
	OAuthServerConsentActionDeny    OAuthServerConsentAction = "deny"
)

// OAuth2 error codes per RFC 6749
const (
	oAuth2ErrorInvalidRequest = "invalid_request"
	oAuth2ErrorServerError    = "server_error"
	oAuth2ErrorAccessDenied   = "access_denied"
)

// OAuthServerAuthorize handles GET /oauth/authorize
func (s *Server) OAuthServerAuthorize(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// validate basic required parameters (client_id, redirect_uri)
// this errors wont be redirected, just returned in the json

// Parse client_id as UUID

// validate redirect_uri matches client's registered URIs

// Invalid redirect_uri should NOT redirect per OAuth2 spec since we can't trust it

// From this point on, we have valid client + redirect_uri + all params, so we can redirect errors
// validate all other parameters - now we can redirect errors

// Store authorization request in database (without user initially)

// Error creating authorization - redirect with server_error

// Redirect to authorization path with authorization_id

// OAuth authorization path not configured - redirect with server_error

// OAuthServerGetAuthorization handles GET /oauth/authorizations/{authorization_id}
func (s *Server) OAuthServerGetAuthorization(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate request origin - the request must come from the site URL as we redirected there at the first place

// Get authenticated user

// Lookup, user association, consent check, and optional auto-approve
// run under a FOR UPDATE SKIP LOCKED row lock so two concurrent callers
// cannot both claim the same pending authorization and each receive a
// valid authorization code.

// Commit the MarkExpired update but still return not-found.

// OAuthServerConsent handles POST /oauth/authorizations/{authorization_id}/consent
func (s *Server) OAuthServerConsent(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate request origin - the request must come from the site URL as we redirected there at the first place

// Get authenticated user

// Row is locked FOR UPDATE SKIP LOCKED so concurrent approve/deny
// requests for the same authorization are serialised and the ownership
// check below can't race a SetUser from another request.

// Commit the MarkExpired update but still return not-found.

// Approve authorization

// Store consent for future use

// Build success redirect URL

// Deny authorization

// Build error redirect URL
// Errors are being returned to the client in the redirect url per OAuth2 spec

// Return redirect URL to frontend

// Helper functions

// validateRequestOrigin checks if the request is coming from an authorized origin
func (s *Server) validateRequestOrigin(r *http.Request) error {
	_ = "STUB: not implemented"
	// Check Origin header
	// browsers add this header by default, we can at least prevent some basic cross-origin attacks
	return nil
}

// Empty Origin header is ok (e.g., for backend-originated requests or mobile apps)

// validateBasicAuthorizeParams validates only client_id and redirect_uri (needed before we can redirect errors)
func (s *Server) validateBasicAuthorizeParams(params *AuthorizeParams) (*AuthorizeParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateRemainingAuthorizeParams validates all other parameters (can redirect errors since we have valid client + redirect_uri)
func (s *Server) validateRemainingAuthorizeParams(params *AuthorizeParams) error {
	_ = "STUB: not implemented"
	return nil
}

// OAuth 2.1 only supports "code" response type

// Validate scopes

// Resource parameter validation (per RFC 8707)

// PKCE validation

func (s *Server) validatePKCEParams(codeChallengeMethod, codeChallenge string) error {
	_ = "STUB: not implemented"
	// PKCE is mandatory for the authorization code flow OAuth2.1
	// Both code_challenge and code_challenge_method must be provided together
	return nil
}

// Validate code challenge method (case-insensitive)

// Validate code challenge format and length (per OAuth2 spec)

// validateResourceParam validates the resource parameter per RFC 8707
func (s *Server) validateResourceParam(resource string) error {
	_ = "STUB: not implemented"
	// Resource parameter is optional
	return nil
}

// Parse URL to validate it's an absolute URI

// Must be an absolute URI (have scheme)

// Must not include a fragment component

// Should not include a query component

// validateScopes validates the requested scopes
func (s *Server) validateScopes(scopeString string) error { _ = "STUB: not implemented"; return nil }

// Validate each scope against the centrally defined supported scopes

func (s *Server) isValidRedirectURI(client *models.OAuthServerClient, redirectURI string) bool {
	_ = "STUB: not implemented"
	return false
}

// exact string matching per OAuth2 spec

func (s *Server) consentCoversScopes(consent *models.OAuthServerConsent, requestedScope string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Server) buildSuccessRedirectURL(authorization *models.OAuthServerAuthorization) string {
	_ = "STUB: not implemented"
	return ""
}

// buildErrorRedirectURL builds an error redirect URL with the given parameters
func (s *Server) buildErrorRedirectURL(redirectURI, errorCode, errorDescription, state string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildAuthorizationURL safely joins a base URL with a path, handling slashes correctly
func (s *Server) buildAuthorizationURL(baseURL, pathToJoin string) string {
	_ = "STUB: not implemented"
	// Trim trailing slash from baseURL
	return ""
}

// Ensure pathToJoin starts with a slash
