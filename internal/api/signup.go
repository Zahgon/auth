package api

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

// SignupParams are the parameters the Signup endpoint accepts
type SignupParams struct {
	Email               string                 `json:"email"`
	Phone               string                 `json:"phone"`
	Password            string                 `json:"password"`
	Data                map[string]interface{} `json:"data"`
	Provider            string                 `json:"-"`
	Aud                 string                 `json:"-"`
	Channel             string                 `json:"channel"`
	CodeChallengeMethod string                 `json:"code_challenge_method"`
	CodeChallenge       string                 `json:"code_challenge"`
}

func (a *API) validateSignupParams(ctx context.Context, p *SignupParams) error {
	_ = "STUB: not implemented"
	return nil
}

// PKCE not needed as phone signups already return access token in body

func (p *SignupParams) ConfigureDefaults() { _ = "STUB: not implemented"; return }

// For backwards compatibility, we default to SMS if params Channel is not specified

func (params *SignupParams) ToUserModel(isSSOUser bool) (user *models.User, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handles external provider case

// TODO: Deprecate "provider" field

// Signup is the endpoint for registering a new user
func (a *API) Signup(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// always call this outside of a database transaction as this method
// can be computationally hard and block due to password hashing
/* <- isSSOUser */

// do not update the user because we can't be sure of their claimed identity

// handles case where Mailer.Autoconfirm is true or Phone.Autoconfirm is true

// add extra context to indicate this is immediate login right after signup

// Remove sensitive fields

// Trigger the after user created hook.

// sanitizeUser removes all user sensitive information from the user object
// Should be used whenever we want to prevent information about whether a user is registered or not from leaking
func sanitizeUser(u *models.User, params *SignupParams) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sanitize app_metadata

// sanitize param fields

func (a *API) signupNewUser(conn *storage.Connection, user *models.User) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// there may be triggers or generated column values in the database that will modify the
// user data as it is being inserted. thus we load the user object
// again to fetch those changes.
