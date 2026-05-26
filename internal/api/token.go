package api

import (
	"context"
	"net/http"

	"github.com/gofrs/uuid"

	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
	"github.com/supabase/auth/internal/tokens"
)

// Aliases for backward compatibility
type AccessTokenClaims = tokens.AccessTokenClaims
type AccessTokenResponse = tokens.AccessTokenResponse

// PasswordGrantParams are the parameters the ResourceOwnerPasswordGrant method accepts
type PasswordGrantParams struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// PKCEGrantParams are the parameters the PKCEGrant method accepts
type PKCEGrantParams struct {
	AuthCode     string `json:"auth_code"`
	CodeVerifier string `json:"code_verifier"`
}

const useCookieHeader = "x-use-cookie"
const InvalidLoginMessage = "Invalid login credentials"

// Token is the endpoint for OAuth access token requests
func (a *API) Token(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// set above

// ResourceOwnerPasswordGrant implements the password grant type flow
func (a *API) ResourceOwnerPasswordGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// directly change this in the database without
// calling user.UpdatePassword() because this
// is not a password change, just encryption
// change in the database

func (a *API) PKCE(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// There is a slight problem with this as it will pick-up the
// User-Agent and IP addresses from the server if used on the server
// side. Currently there's no mechanism to distinguish, but the server
// can be told to at least propagate the User-Agent header.

// Sanity check in case user ID was not set properly

// error type is already handled in issueRefreshToken

// Because not all providers give out a refresh token
// See corresponding OAuth2 spec: <https://www.rfc-editor.org/rfc/rfc6749.html#section-5.1>

func (a *API) generateAccessToken(r *http.Request, tx *storage.Connection, user *models.User, sessionId *uuid.UUID, authenticationMethod models.AuthenticationMethod) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (a *API) issueRefreshToken(r *http.Request, headers http.Header, conn *storage.Connection, user *models.User, authenticationMethod models.AuthenticationMethod, grantParams models.GrantParams) (*tokens.AccessTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *API) updateMFASessionAndClaims(r *http.Request, tx *storage.Connection, user *models.User, authenticationMethod models.AuthenticationMethod, grantParams models.GrantParams) (*tokens.AccessTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// issue a new refresh token on successful verification

// Incrementing the refresh token counter by 2 here is
// counter intuitive, but is important for security. It
// means that the previous refresh token (issued with
// AAL1) will no longer be able to issue AAL2 sessions.
// It forces the client to have received the refresh
// token from the MFA verification flow.

// Legacy RTs: swap to ensure current token is the latest one
