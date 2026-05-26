package api

import (
	"net/http"
)

// ResendConfirmationParams holds the parameters for a resend request
type ResendConfirmationParams struct {
	Type                string `json:"type"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
}

func (p *ResendConfirmationParams) Validate(a *API) error { _ = "STUB: not implemented"; return nil }

// type does not match one of the above

// both email and phone are empty

// Recover sends a recovery email
func (a *API) Resend(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// if the user's email is confirmed already, we don't need to send a confirmation email again

// if the user's phone is confirmed already, we don't need to send a confirmation sms again

// do not resend if user doesn't have a new email address

// do not resend if user doesn't have a new phone number
