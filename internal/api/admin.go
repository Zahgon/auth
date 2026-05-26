package api

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/models"
)

type AdminUserParams struct {
	Id           string                 `json:"id"`
	Aud          string                 `json:"aud"`
	Role         string                 `json:"role"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	Password     *string                `json:"password"`
	PasswordHash string                 `json:"password_hash"`
	EmailConfirm bool                   `json:"email_confirm"`
	PhoneConfirm bool                   `json:"phone_confirm"`
	UserMetaData map[string]interface{} `json:"user_metadata"`
	AppMetaData  map[string]interface{} `json:"app_metadata"`
	BanDuration  string                 `json:"ban_duration"`
}

type adminUserDeleteParams struct {
	ShouldSoftDelete bool `json:"should_soft_delete"`
}

type adminUserUpdateFactorParams struct {
	FriendlyName string `json:"friendly_name"`
	Phone        string `json:"phone"`
}

type AdminListUsersResponse struct {
	Users []*models.User `json:"users"`
	Aud   string         `json:"aud"`
}

func (a *API) loadUser(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Use only after requireAuthentication, so that there is a valid user
func (a *API) loadFactor(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *API) getAdminParams(r *http.Request) (*AdminUserParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// adminUsers responds with a list of all users in a given audience
func (a *API) adminUsers(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// adminUserGet returns information about a single user
func (a *API) adminUserGet(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// adminUserUpdate updates a single user object
func (a *API) adminUserUpdate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// if the user doesn't have an existing email
// then updating the user's email should create a new email identity

// update the existing email identity

// if the user doesn't have an existing phone
// then updating the user's phone should create a new phone identity

// update the existing phone identity

// adminUserCreate creates a new user based on the provided data
func (a *API) adminUserCreate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Deprecate "provider" field
// default to the first provider in the providers slice

// complying with the user.AppMetaData["provider"] field as above

// adminUserDelete deletes a user
func (a *API) adminUserDelete(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// ShouldSoftDelete defaults to false

// we only want to parse the body if it's not empty
// retrieveRequestParams will handle any errors with stream

// user has been soft deleted already

// hard delete all associated factors

// hard delete all associated sessions

func (a *API) adminUserDeleteFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) adminUserGetFactors(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// adminUserUpdate updates a single factor object
func (a *API) adminUserUpdateFactor(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
