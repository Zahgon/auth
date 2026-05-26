package shared

import (
	"context"

	"github.com/supabase/auth/internal/models"
)

// ContextKey is the type for context keys to avoid collisions
type ContextKey string

func (c ContextKey) String() string { _ = "STUB: not implemented"; return "" }

// Context keys used across packages
const (
	UserKey              ContextKey = "user"
	SessionKey           ContextKey = "session"
	OAuthServerClientKey ContextKey = "oauth_server_client"
)

// GetUser reads the user from the context - shared implementation
func GetUser(ctx context.Context) *models.User { _ = "STUB: not implemented"; return nil }

// WithUser adds the user to the context - shared implementation
func WithUser(ctx context.Context, u *models.User) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetSession reads the session from the context - shared implementation
func GetSession(ctx context.Context) *models.Session { _ = "STUB: not implemented"; return nil }

// WithSession adds the session to the context - shared implementation
func WithSession(ctx context.Context, s *models.Session) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithOAuthServerClient adds an OAuth server client to the context
func WithOAuthServerClient(ctx context.Context, client *models.OAuthServerClient) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetOAuthServerClient retrieves an OAuth server client from the context
func GetOAuthServerClient(ctx context.Context) *models.OAuthServerClient {
	_ = "STUB: not implemented"
	return nil
}
