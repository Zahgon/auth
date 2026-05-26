package api

import (
	"context"
	"net/http"
	"time"

	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

const (
	smsVerification         = "sms"
	phoneChangeVerification = "phone_change"
	// includes signupVerification and magicLinkVerification
)

const (
	zeroConfirmation int = iota
	singleConfirmation
)

// Only applicable when SECURE_EMAIL_CHANGE_ENABLED
const singleConfirmationAccepted = "Confirmation link accepted. Please proceed to confirm link sent to the other email"

// VerifyParams are the parameters the Verify endpoint accepts
type VerifyParams struct {
	Type       string `json:"type"`
	Token      string `json:"token"`
	TokenHash  string `json:"token_hash"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	RedirectTo string `json:"redirect_to"`
}

func (p *VerifyParams) Validate(r *http.Request, a *API) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: deprecate the token query param from GET /verify and use token_hash instead (breaking change)

// Verify exchanges a confirmation or recovery token to a refresh token
func (a *API) Verify(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// this should have been handled by Chi

func (a *API) verifyGet(w http.ResponseWriter, r *http.Request, params *VerifyParams) error {
	_ = "STUB: not implemented"
	return nil
}

// only one OTP is confirmed at this point, so we return early and ask the user to confirm the second OTP

// Reload user model from db.
// This is important for refreshing the data in any generated columns like IsAnonymous.

func (a *API) verifyPost(w http.ResponseWriter, r *http.Request, params *VerifyParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Reload user model from db.
// This is important for refreshing the data in any generated columns like IsAnonymous.

// Record login for analytics - determine provider based on verification type
// default

func (a *API) signupVerify(r *http.Request, ctx context.Context, conn *storage.Connection, user *models.User) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sign them up with temporary password, and require application
// to present the user with a password set form

// password generation must succeed

func (a *API) recoverVerify(r *http.Request, conn *storage.Connection, user *models.User) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *API) smsVerify(r *http.Request, conn *storage.Connection, user *models.User, params *VerifyParams) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// confirming the phone change should create a new phone identity if the user doesn't have one

// Send phone changed notification email if enabled and phone was changed

// Log the error but don't fail the verification

// Send identity linked notification email if a new phone identity was created

// Log the error but don't fail the verification

func (a *API) prepErrorRedirectURL(err *HTTPError, r *http.Request, rurl string, flowType models.FlowType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Maintain separate query params for hash and query

// Additionally, may override existing error query param if set to PKCE.

// Left as hash fragment to comply with spec.
// Add Supabase Auth identifier to help clients distinguish Supabase Auth redirects

func (a *API) prepRedirectURL(message string, rurl string, flowType models.FlowType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Add Supabase Auth identifier to help clients distinguish Supabase Auth redirects

func (a *API) prepPKCERedirectURL(rurl, code string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *API) emailChangeVerify(r *http.Request, conn *storage.Connection, params *VerifyParams, user *models.User) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// one email is confirmed at this point if GOTRUE_MAILER_SECURE_EMAIL_CHANGE_ENABLED is enabled

// confirming the email change should create a new email identity if the user doesn't have one

// send an Email Changed email notification to the user's old email address

// we don't want to fail the whole request if the email can't be sent

func (a *API) verifyTokenHash(conn *storage.Connection, params *VerifyParams) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// need to find user by confirmation token or recovery token with the token hash

// verifyUserAndToken verifies the token associated to the user based on the verify type
func (a *API) verifyUserAndToken(conn *storage.Connection, params *VerifyParams, aud string) (*models.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Since the email change could be trigger via the implicit or PKCE flow,
// the query used has to also check if the token saved in the db contains the pkce_ prefix

// if the type is emailOTPVerification, we'll check both the confirmation_token and recovery_token columns

// isOtpValid checks the actual otp sent against the expected otp and ensures that it's within the valid window
func isOtpValid(actual, expected string, sentAt *time.Time, otpExp uint) bool {
	_ = "STUB: not implemented"
	return false
}

func isOtpExpired(sentAt *time.Time, otpExp uint) bool { _ = "STUB: not implemented"; return false }

// #nosec G115

// isPhoneOtpVerification checks if the verification came from a phone otp
func isPhoneOtpVerification(params *VerifyParams) bool { _ = "STUB: not implemented"; return false }

// isEmailOtpVerification checks if the verification came from an email otp
func isEmailOtpVerification(params *VerifyParams) bool { _ = "STUB: not implemented"; return false }

func isUsingTokenHash(params *VerifyParams) bool { _ = "STUB: not implemented"; return false }

// emailAddressChanged checks if the email address has changed, ensuring neither is empty
func emailAddressChanged(oldEmail, newEmail string) bool { _ = "STUB: not implemented"; return false }

// phoneNumberChanged checks if the phone number has changed, ensuring neither is empty
func phoneNumberChanged(oldPhone, newPhone string) bool { _ = "STUB: not implemented"; return false }
