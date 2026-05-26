package api

import (
	"context"
	"net/http"
)

// requireAuthentication checks incoming requests for tokens presented using the Authorization header
func (a *API) requireAuthentication(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *API) requireNotAnonymous(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *API) requireAdmin(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	// Find the administrative user
	return *new(context.Context), nil
}

// successful authentication

func (a *API) extractBearerToken(r *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *API) parseJWTClaims(bearer string, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// otherwise try to use fallback

// preserve backward compatibility for cases where the kid is not set

func (a *API) maybeLoadUserOrSession(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Also store in shared context for cross-package access (e.g., oauthserver package)
