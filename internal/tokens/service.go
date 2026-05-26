package tokens

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xeipuuv/gojsonschema"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

const retryLoopDuration = 5.0

// AMRClaim supports unmarshalling AMR as either strings or AMREntry objects.
type AMRClaim []models.AMREntry

// UnmarshalJSON accepts either an array of strings or AMREntry objects.
func (a *AMRClaim) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Handle null explicitly - null cannot be unmarshaled into a slice
	return nil
}

// AccessTokenClaims is a struct thats used for JWT claims
type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Email                         string                 `json:"email"`
	Phone                         string                 `json:"phone"`
	AppMetaData                   map[string]interface{} `json:"app_metadata"`
	UserMetaData                  map[string]interface{} `json:"user_metadata"`
	Role                          string                 `json:"role"`
	AuthenticatorAssuranceLevel   string                 `json:"aal,omitempty"`
	AuthenticationMethodReference AMRClaim               `json:"amr,omitempty"`
	SessionId                     string                 `json:"session_id,omitempty"`
	IsAnonymous                   bool                   `json:"is_anonymous"`
	ClientID                      string                 `json:"client_id,omitempty"`
	Scope                         string                 `json:"scope,omitempty"`
}

// IDTokenClaims represents OpenID Connect ID Token claims
type IDTokenClaims struct {
	jwt.RegisteredClaims
	Nonce               string `json:"nonce,omitempty"`
	AuthTime            int64  `json:"auth_time"`
	Email               string `json:"email,omitempty"`
	EmailVerified       bool   `json:"email_verified"` // not omitempty because it's required by OIDC spec
	PhoneNumber         string `json:"phone_number,omitempty"`
	PhoneNumberVerified bool   `json:"phone_number_verified"` // not omitempty because it's required by OIDC spec
	Name                string `json:"name,omitempty"`
	Picture             string `json:"picture,omitempty"`
	UpdatedAt           int64  `json:"updated_at,omitempty"`
	PreferredUsername   string `json:"preferred_username,omitempty"`
	ClientID            string `json:"client_id,omitempty"`
}

// AccessTokenResponse represents an OAuth2 success response
type AccessTokenResponse struct {
	Token                string       `json:"access_token"`
	TokenType            string       `json:"token_type"` // Bearer
	ExpiresIn            int          `json:"expires_in"`
	ExpiresAt            int64        `json:"expires_at"`
	RefreshToken         string       `json:"refresh_token"`
	User                 *models.User `json:"user"`
	ProviderAccessToken  string       `json:"provider_token,omitempty"`
	ProviderRefreshToken string       `json:"provider_refresh_token,omitempty"`
	WeakPassword         interface{}  `json:"weak_password,omitempty"`
	IDToken              string       `json:"id_token,omitempty"` // OIDC ID Token
}

// GenerateAccessTokenParams contains parameters for generating access tokens
type GenerateAccessTokenParams struct {
	User                 *models.User
	SessionID            *uuid.UUID
	AuthenticationMethod models.AuthenticationMethod
	ClientID             *uuid.UUID // OAuth2 server client ID if applicable
}

// GenerateIDTokenParams contains parameters for generating OIDC ID tokens
type GenerateIDTokenParams struct {
	User     *models.User
	ClientID uuid.UUID  // OAuth2 client ID (required for ID tokens)
	Nonce    string     // OIDC nonce from authorization request (optional)
	AuthTime *time.Time // Time when authentication occurred (optional, uses user.LastSignInAt if not provided)
	Scopes   []string   // OAuth scopes granted (used to filter claims)
}

// RefreshTokenGrantParams contains parameters for refresh token grant
type RefreshTokenGrantParams struct {
	RefreshToken string
	ClientID     *uuid.UUID // OAuth2 server client ID if applicable
}

// AsRedirectURL encodes the AccessTokenResponse as a redirect URL that
// includes the access token response data in a URL fragment.
func (r *AccessTokenResponse) AsRedirectURL(redirectURL string, extraParams url.Values) string {
	_ = "STUB: not implemented"
	return ""
}

// Add Supabase Auth identifier to help clients distinguish Supabase Auth redirects

// HookManager interface for access token hooks
type HookManager interface {
	InvokeHook(tx *storage.Connection, r *http.Request, input any, output any) error
}

// Service handles token operations
type Service struct {
	config      *conf.GlobalConfiguration
	hookManager HookManager
	now         func() time.Time
}

// NewService creates a new token service
func NewService(config *conf.GlobalConfiguration, hookManager HookManager) *Service {
	_ = "STUB: not implemented"
	return nil
}

// Default to system time

// SetTimeFunc allows overriding the time function (only for testing!!)
func (s *Service) SetTimeFunc(timeFunc func() time.Time) { _ = "STUB: not implemented"; return }

// RefreshTokenGrant implements the refresh_token grant type flow
func (s *Service) RefreshTokenGrant(ctx context.Context, db *storage.Connection, r *http.Request, responseHeaders http.Header, params RefreshTokenGrantParams) (*AccessTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A 5 second retry loop is used to make sure that refresh token
// requests do not waste database connections waiting for each other.
// Instead of waiting at the database level, they're waiting at the API
// level instead and retry to refresh the locked row every 10-30
// milliseconds.

// a refresh token won't have a session if it's created prior to the sessions table introduced

// OAuth client validation will be done inside the transaction

// do nothing

// Basic checks above passed, now we need to serialize access
// to the session in a transaction so that there's no
// concurrent modification. In the event that the refresh
// token's row or session is locked, the transaction is closed
// and the whole process will be retried a bit later so that
// the connection pool does not get exhausted.

/* forUpdate */

// because forUpdate was set, and the
// previous check outside the
// transaction found a refresh token
// and session, but now we're getting a
// IsNotFoundError, this means that the
// refresh token row and session are
// probably locked so we need to retry
// in a few milliseconds.

// Validate OAuth client consistency between session and current request

// Session has an OAuth client, current request must have matching client

// Session has no OAuth client, current request should not have one either

/* forUpdate */

// because forUpdate was set, and the
// previous check outside the
// transaction found a user and
// session, but now we're getting a
// IsNotFoundError, this means that the
// user is locked and we need to retry
// in a few milliseconds

// go through all sessions of the user and
// check if the current session is the user's
// most recently refreshed valid session

// current session, skip it

// session is not valid so it
// can't be regarded as active
// on the user

// if tags are specified,
// ignore sessions with a
// mismatching tag

// since token is not the refresh token
// of s, we can't use it's UpdatedAt
// time to compare!

// session is not the most
// recently active one

// this session is the user's active session

// refresh token row and session are locked at this
// point, cannot be concurrently refreshed

// Token was revoked, but it's the
// parent of the currently active one.
// This indicates that the client was
// not able to store the result when it
// refreshed token. This case is
// allowed, provided we return back the
// active refresh token instead of
// creating a new one.

// For a revoked refresh token to be reused, it
// has to fall within the reuse interval.

// not OK to reuse this token

// Revoke all tokens in token family

// convert the session ID to a number in the range [0, 100) and check whether it should be upgraded
// we don't want a % of refresh token requests, but a % of sessions here!

// #nosec
// #nosec

// got v1 refresh token that should be upgraded to v2
// so discard the previously generated v1 token, revoke it and issue a v2 token instead

// session already set up, increment the counter by 1

// refresh token was not issued by this server

// normal refresh token use

// refresh token is being reused

// This is caused when the client has
// failed to receive or save the
// response from the last refresh token
// requests. This occurs more
// frequently than you can imagine, so
// it's an allowed reuse.

// Concurrent refreshes occur when the
// client sends off multiple refresh
// token requests at once or close by.
// Often this happens when your browser
// remembers multiple tabs of the app,
// which were paused by it or by the OS
// (such as you quitting the browser)
// and then opening it back up. If the
// app uses SSR it is likely that the N
// open tabs will immediately send a
// request to the app's hosting server,
// which will attempt to concurrently
// refresh the session at once using
// refresh token.

// When reuse is allowed, we do
// not increment the counter.
// This allows all of the
// concurrent clients to
// synchronize their state
// within the refresh token
// reuse interval to the
// currently active refresh
// token.

// Reuse is not allowed, in
// which case the whole session
// must go preventing any
// client with any refresh and
// access token for this
// session from being used.

// Reuse is not allowed, but no
// refresh token rotation
// enabled. So only fail this
// request.

// refresh token and session row were likely locked, so
// we need to wait a moment before retrying the whole
// process anew
// #nosec

// GenerateAccessToken generates an access token using shared logic
func (s *Service) GenerateAccessToken(r *http.Request, tx *storage.Connection, params GenerateAccessTokenParams) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// if user has mfa enabled and the session has not yet been upgraded
// and Limit duration of AAL1 sessions is enabled
// expiresAt should be set to the maximum duration for low aal sessions
// don't allow sessions.AllowLowAAL to exceed config.JWT.Exp

// Get scopes from session if this is an OAuth session

// GenerateIDToken generates an OpenID Connect ID Token
// IDToken is generated only when the signing key is an asymmetric one.
// HS256 is not supported for ID token signing.
// Claims are filtered based on the granted scopes per OIDC spec:
// - openid: sub (always included when openid scope is present)
// - email: email, email_verified
// - profile: name, picture, updated_at, preferred_username
// - phone: phone_number, phone_number_verified
func (s *Service) GenerateIDToken(params GenerateIDTokenParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Determine auth_time (when authentication occurred)

// Fallback to current time

// ID tokens typically expire in 1 hour (3600 seconds) per OIDC spec

// Build ID token claims per OIDC spec
// Base claims (always included)

// Add nonce if provided

// Add scope-specific claims
// Check if scope was granted before adding claims

// Email scope: email, email_verified

// Phone scope: phone_number, phone_number_verified

// Profile scope: name, picture, updated_at, preferred_username

// Extract name from user metadata if available

// Fallback to email as name if no name is set

// Extract picture URL from user metadata if available

// Add preferred_username if available

// Add updated_at timestamp

// Sign the ID token with the same key as access tokens

// IssueRefreshToken creates a new refresh token and access token
func (s *Service) IssueRefreshToken(r *http.Request, responseHeaders http.Header, conn *storage.Connection, user *models.User, authenticationMethod models.AuthenticationMethod, grantParams models.GrantParams) (*AccessTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Account for Hook Error

// SignJWT signs a JWT token with the configured signing key
func SignJWT(config *conf.JWTConfiguration, claims jwt.Claims) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// this serializes the aud claim to a string

var schemaLoader = gojsonschema.NewStringLoader(MinimumViableTokenSchema)

func validateTokenClaims(outputClaims map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec
const MinimumViableTokenSchema = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "aud": {
      "type": ["string", "array"]
    },
    "exp": {
      "type": "integer"
    },
    "jti": {
      "type": "string"
    },
    "iat": {
      "type": "integer"
    },
    "iss": {
      "type": "string"
    },
    "nbf": {
      "type": "integer"
    },
    "sub": {
      "type": "string"
    },
    "email": {
      "type": "string"
    },
    "phone": {
      "type": "string"
    },
    "app_metadata": {
      "type": "object",
      "additionalProperties": true
    },
    "user_metadata": {
      "type": "object",
      "additionalProperties": true
    },
    "role": {
      "type": "string"
    },
    "aal": {
      "type": "string"
    },
    "amr": {
      "type": "array",
      "items": {
        "anyOf": [
          {"type": "string"},
          {"type": "object"}
        ]
      }
    },
    "session_id": {
      "type": "string"
    },
    "client_id": {
      "type": "string"
    }
  },
  "required": ["aud", "exp", "iat", "sub", "email", "phone", "role", "aal", "session_id", "is_anonymous"]
}`
