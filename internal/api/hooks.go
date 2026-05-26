package api

import (
	"net/http"

	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

func (a *API) triggerAfterUserCreated(
	r *http.Request,
	conn *storage.Connection,
	user *models.User,
) error {
	_ = "STUB: not implemented"
	return nil
}

// We still check tx because we want to make sure we aren't calling this
// trigger in code paths that haven't actually created the user yet.

func (a *API) triggerBeforeUserCreated(
	r *http.Request,
	db *storage.Connection,
	user *models.User,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) triggerBeforeUserCreatedExternal(
	r *http.Request,
	db *storage.Connection,
	userData *provider.UserProvidedData,
	providerType string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func checkTX(conn *storage.Connection) error { _ = "STUB: not implemented"; return nil }
