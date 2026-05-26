package api

import (
	"context"
	"net/url"

	"github.com/gofrs/uuid"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/supabase/auth/internal/models"
)

type contextKey string

func (c contextKey) String() string { _ = "STUB: not implemented"; return "" }

const (
	externalProviderTypeKey          = contextKey("external_provider_type")
	externalProviderEmailOptionalKey = contextKey("external_provider_allow_no_email")

	tokenKey            = contextKey("jwt")
	inviteTokenKey      = contextKey("invite_token")
	signatureKey        = contextKey("signature")
	targetUserKey       = contextKey("target_user")
	factorKey           = contextKey("factor")
	sessionKey          = contextKey("session")
	externalReferrerKey = contextKey("external_referrer")
	functionHooksKey    = contextKey("function_hooks")
	adminUserKey        = contextKey("admin_user")
	oauthTokenKey       = contextKey("oauth_token") // for OAuth1.0, also known as request token
	oauthVerifierKey    = contextKey("oauth_verifier")
	ssoProviderKey      = contextKey("sso_provider")
	externalHostKey     = contextKey("external_host")
	oauthClientStateKey = contextKey("oauth_client_state_id")
	flowStateContextKey = contextKey("flow_state")
)

// withToken adds the JWT token to the context.
func withToken(ctx context.Context, token *jwt.Token) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getToken reads the JWT token from the context.
func getToken(ctx context.Context) *jwt.Token { _ = "STUB: not implemented"; return nil }

func getClaims(ctx context.Context) *AccessTokenClaims { _ = "STUB: not implemented"; return nil }

// withUser adds the user to the context.
func withUser(ctx context.Context, u *models.User) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// withTargetUser adds the target user for linking to the context.
func withTargetUser(ctx context.Context, u *models.User) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// with Factor adds the factor id to the context.
func withFactor(ctx context.Context, f *models.Factor) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getUser reads the user from the context.
func getUser(ctx context.Context) *models.User { _ = "STUB: not implemented"; return nil }

// getTargetUser reads the user from the context.
func getTargetUser(ctx context.Context) *models.User { _ = "STUB: not implemented"; return nil }

// getFactor reads the factor id from the context
func getFactor(ctx context.Context) *models.Factor { _ = "STUB: not implemented"; return nil }

// withSession adds the session to the context.
func withSession(ctx context.Context, s *models.Session) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getSession reads the session from the context.
func getSession(ctx context.Context) *models.Session { _ = "STUB: not implemented"; return nil }

// withSignature adds the provided request ID to the context.
func withSignature(ctx context.Context, id string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func withInviteToken(ctx context.Context, token string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func withOAuthClientStateID(ctx context.Context, oauthClientStateID uuid.UUID) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getOAuthClientStateID(ctx context.Context) uuid.UUID {
	_ = "STUB: not implemented"
	return *new(uuid.UUID)
}

// withFlowState stores the entire FlowState object in the context
func withFlowState(ctx context.Context, flowState *models.FlowState) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getFlowState retrieves the FlowState object from the context
func getFlowState(ctx context.Context) *models.FlowState { _ = "STUB: not implemented"; return nil }

func getInviteToken(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// withExternalProviderType adds the provided request ID to the context.
func withExternalProviderType(ctx context.Context, id string, emailOptional bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getExternalProviderType returns the provider type and whether user data without email address should be allowed.
func getExternalProviderType(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func withExternalReferrer(ctx context.Context, token string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getExternalReferrer(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// withAdminUser adds the admin user to the context.
func withAdminUser(ctx context.Context, u *models.User) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getAdminUser reads the admin user from the context.
func getAdminUser(ctx context.Context) *models.User { _ = "STUB: not implemented"; return nil }

// withRequestToken adds the request token to the context
func withRequestToken(ctx context.Context, token string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getRequestToken(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func withOAuthVerifier(ctx context.Context, token string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getOAuthVerifier(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func withSSOProvider(ctx context.Context, provider *models.SSOProvider) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getSSOProvider(ctx context.Context) *models.SSOProvider { _ = "STUB: not implemented"; return nil }

func withExternalHost(ctx context.Context, u *url.URL) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getExternalHost(ctx context.Context) *url.URL { _ = "STUB: not implemented"; return nil }
