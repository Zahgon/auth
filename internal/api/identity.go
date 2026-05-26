package api

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

func (a *API) DeleteIdentity(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Send identity unlinked notification email if enabled and user has an email

// Log the error but don't fail the unlinking

func (a *API) LinkIdentity(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G710

func (a *API) linkIdentityToUser(r *http.Request, ctx context.Context, tx *storage.Connection, userData *provider.UserProvidedData, providerType string) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send identity linked notification email if enabled and user has an email

// Log the error but don't fail the linking
