package api

import (
	"net/http"

	"github.com/supabase/auth/internal/conf"
)

// OtpParams contains the request body params for the otp endpoint
type OtpParams struct {
	Email               string                 `json:"email"`
	Phone               string                 `json:"phone"`
	CreateUser          bool                   `json:"create_user"`
	Data                map[string]interface{} `json:"data"`
	Channel             string                 `json:"channel"`
	CodeChallengeMethod string                 `json:"code_challenge_method"`
	CodeChallenge       string                 `json:"code_challenge"`
}

// SmsParams contains the request body params for sms otp
type SmsParams struct {
	Phone               string                 `json:"phone"`
	Channel             string                 `json:"channel"`
	Data                map[string]interface{} `json:"data"`
	CodeChallengeMethod string                 `json:"code_challenge_method"`
	CodeChallenge       string                 `json:"code_challenge"`
}

func (p *OtpParams) Validate() error { _ = "STUB: not implemented"; return nil }

func (p *SmsParams) Validate(config *conf.GlobalConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

// Otp returns the MagicLink or SmsOtp handler based on the request body params
func (a *API) Otp(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type SmsOtpResponse struct {
	MessageID string `json:"message_id,omitempty"`
}

// SmsOtp sends the user an otp via sms
func (a *API) SmsOtp(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// For backwards compatibility, we default to SMS if params Channel is not specified

// User either doesn't exist or hasn't completed the signup process.
// Sign them up with temporary password.

// SignupParams must be marshallable

// signups are autoconfirmed, send otp after signup

// SignupParams must be marshallable

func (a *API) shouldCreateUser(r *http.Request, params *OtpParams) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
