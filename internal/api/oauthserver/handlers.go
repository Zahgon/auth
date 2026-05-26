package oauthserver

import (
	"context"
	"net/http"
	"time"

	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/tokens"
)

// OAuth 2.1 Grant Types
const (
	GrantTypeAuthorizationCode = "authorization_code"
	GrantTypeRefreshToken      = "refresh_token"
)

// OAuthServerClientResponse represents the response format for OAuth client operations
type OAuthServerClientResponse struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"` // only returned on registration
	ClientType   string `json:"client_type"`

	RedirectURIs            []string `json:"redirect_uris,omitempty"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method,omitempty"`
	GrantTypes              []string `json:"grant_types,omitempty"`
	ResponseTypes           []string `json:"response_types,omitempty"`
	ClientName              string   `json:"client_name,omitempty"`
	ClientURI               string   `json:"client_uri,omitempty"`
	LogoURI                 string   `json:"logo_uri,omitempty"`

	// Metadata fields
	RegistrationType string    `json:"registration_type,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// OAuthServerClientListResponse represents the response for listing OAuth clients
type OAuthServerClientListResponse struct {
	Clients []OAuthServerClientResponse `json:"clients,omitempty"`
}

// oauthServerClientToResponse converts a model to response format
func oauthServerClientToResponse(client *models.OAuthServerClient) *OAuthServerClientResponse {
	_ = "STUB: not implemented"
	return nil
}

// OAuth 2.1 DCR fields

// Always "code" in OAuth 2.1

// Metadata fields

// LoadOAuthServerClient is middleware that loads an OAuth server client from the URL parameter
func (s *Server) LoadOAuthServerClient(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Parse client_id as UUID

// AdminOAuthServerClientRegister handles POST /admin/oauth/clients (manual registration by admins)
func (s *Server) AdminOAuthServerClientRegister(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Force registration type to manual for admin endpoint

// OAuthServerClientDynamicRegister handles POST /oauth/register (OAuth 2.1 Dynamic Client Registration)
func (s *Server) OAuthServerClientDynamicRegister(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"

	// Check if dynamic registration is enabled
	return nil
}

// OAuthServerClientGet handles GET /admin/oauth/clients/{client_id}
func (s *Server) OAuthServerClientGet(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// OAuthServerClientUpdate handles PUT /admin/oauth/clients/{client_id}
func (s *Server) OAuthServerClientUpdate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Return early if no fields are provided for update

// OAuthServerClientDelete handles DELETE /admin/oauth/clients/{client_id}
func (s *Server) OAuthServerClientDelete(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// OAuthServerClientRegenerateSecret handles POST /admin/oauth/clients/{client_id}/regenerate_secret
func (s *Server) OAuthServerClientRegenerateSecret(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Only confidential clients can have their secrets regenerated

// OAuthServerClientList handles GET /admin/oauth/clients
func (s *Server) OAuthServerClientList(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(cemal) :: Add pagination, check the `/admin/users` endpoint for reference

// OAuthTokenParams represents the parameters for the OAuth token endpoint
type OAuthTokenParams struct {
	GrantType    string `json:"grant_type" form:"grant_type"`
	Code         string `json:"code" form:"code"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri"`
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	CodeVerifier string `json:"code_verifier" form:"code_verifier"`
	Resource     string `json:"resource" form:"resource"`
}

// OAuthToken handles POST /oauth/token
func (s *Server) OAuthToken(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle both JSON and form-encoded requests

// Parse form data

// Validate grant_type

// Validate that the authenticated client is allowed to use the requested grant type

// handleAuthorizationCodeGrant handles the authorization_code grant type
func (s *Server) handleAuthorizationCodeGrant(ctx context.Context, w http.ResponseWriter, r *http.Request, params *OAuthTokenParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Get authenticated client from middleware

// Exchange authorization code for tokens

// Find the OAuth authorization for this authorization code

// Check if the authorization has expired

// Validate that the authorization code was issued for this client

// Validate that (if exists) the resource parameter matches the authorization code resource

// Validate redirect_uri if provided - must match the one used in authorization

// Validate PKCE if used in the authorization

// Get the user for the authorization code

// Exchange the authorization code for tokens

// Store scopes from authorization in session

// Create audit log entry for OAuth token exchange

// Issue the refresh token and access token

// Mark authorization as used - authorization codes are single use
// We could either delete it or mark it as consumed

// Check if we need to generate OIDC ID Token (only if 'openid' scope is present)

// Convert to OAuth-compliant response format (exclude user info for OAuth clients)

// Include ID token if generated (OIDC)

// handleRefreshTokenGrant handles the refresh_token grant type
func (s *Server) handleRefreshTokenGrant(ctx context.Context, w http.ResponseWriter, r *http.Request, params *OAuthTokenParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the token service to handle refresh token grant

// Get OAuth client from context if present

// Convert to OAuth-compliant response format (exclude user info for OAuth clients)

// getTokenService retrieves the token service from the server
func (s *Server) getTokenService() *tokens.Service { _ = "STUB: not implemented"; return nil }

// UserOAuthGrantResponse represents an OAuth grant that a user has authorized
type UserOAuthGrantResponse struct {
	Client    ClientDetailsResponse `json:"client"`
	Scopes    []string              `json:"scopes"`
	GrantedAt time.Time             `json:"granted_at"`
}

// UserListOAuthGrants handles GET /user/oauth/grants
// Lists all OAuth grants that the authenticated user has authorized (active consents)
func (s *Server) UserListOAuthGrants(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Get all active (non-revoked) consents for this user

// Build response with client information

// Fetch client details

// Skip clients that no longer exist or are deleted

// UserRevokeOAuthGrant handles DELETE /user/oauth/grants?client_id=...
// Revokes the user's OAuth grant for a specific client
func (s *Server) UserRevokeOAuthGrant(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse client_id as UUID

// Find the active consent for this user and client

// Revoke the consent in a transaction

// Delete all sessions associated with this OAuth client for this user
// This will invalidate all refresh tokens for those sessions

// Create audit log entry

// OAuthUserInfo handles GET /oauth/userinfo (OIDC UserInfo endpoint)
// Per OIDC Core Section 5.3
//
// Returns user information filtered by the scopes granted in the access token:
// - openid: sub (user ID) - always included as base claim
// - email: email, email_confirmed_at, new_email
// - profile: name, picture, preferred_username, updated_at, user_metadata
// - phone: phone, phone_confirmed_at, new_phone
func (s *Server) OAuthUserInfo(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"

	// Get authenticated user from context (set by requireAuthentication middleware)
	return nil
}

// Get the session to retrieve scopes
// The access token contains session_id claim, and requireAuthentication middleware
// loads the session into context

// Build base userInfo response with 'sub' (required by OIDC spec)

// If no session in context, this is likely a non-OAuth token
// Return minimal user info (just sub claim)

// Get scopes from session

// Add scope-specific claims

// Email scope claims

// Profile scope claims

// Extract name from user metadata

// Extract picture

// Extract preferred_username

// Add updated_at

// Include user_metadata with profile scope

// Phone scope claims
