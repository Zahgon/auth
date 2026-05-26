package api

import (
	"context"
	"net/http"
)

// UserUpdateParams parameters for updating a user
type UserUpdateParams struct {
	Email               string                 `json:"email"`
	Password            *string                `json:"password"`
	CurrentPassword     *string                `json:"current_password,omitempty"`
	Nonce               string                 `json:"nonce"`
	Data                map[string]interface{} `json:"data"`
	AppData             map[string]interface{} `json:"app_metadata,omitempty"`
	Phone               string                 `json:"phone"`
	Channel             string                 `json:"channel"`
	CodeChallenge       string                 `json:"code_challenge"`
	CodeChallengeMethod string                 `json:"code_challenge_method"`
}

func (a *API) validateUserUpdateParams(ctx context.Context, p *UserUpdateParams) error {
	_ = "STUB: not implemented"
	return nil
}

// UserGet returns a user
func (a *API) UserGet(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// UserUpdate updates fields on a user
func (a *API) UserUpdate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Check if a user is SSO via rows in identities table, not via this flag.

// we require reauthentication if the user hasn't signed in recently in the current session

// current password required when updating password

// ensure user is not in a password recovery flow

// send a Password Changed email notification to the user to inform them that their password has been changed

// we don't want to fail the whole request if the email can't be sent

// anonymous users can add an email with automatic confirmation, which is similar to signing up
// permanent users always need to verify their email address when changing it
