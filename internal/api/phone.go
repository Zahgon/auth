package api

import (
	"net/http"
	"regexp"
	"text/template"

	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

var e164Format = regexp.MustCompile("^[1-9][0-9]{1,14}$")

const (
	phoneConfirmationOtp     = "confirmation"
	phoneReauthenticationOtp = "reauthentication"
)

func validatePhone(phone string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// validateE164Format checks if phone number follows the E.164 format
func validateE164Format(phone string) bool { _ = "STUB: not implemented"; return false }

// formatPhoneNumber removes "+" and whitespaces in a phone number
func formatPhoneNumber(phone string) string { _ = "STUB: not implemented"; return "" }

// sendPhoneConfirmation sends an otp to the user's phone number
func (a *API) sendPhoneConfirmation(r *http.Request, tx *storage.Connection, user *models.User, phone, otpType string, channel string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// intentionally keeping this before the test OTP, so that the behavior
// of regular and test OTPs is similar

// not using test OTPs

// TODO(km): Deprecate this behaviour - rate limits should still be applied to autoconfirm

// apply rate limiting before the sms is sent out

func generateSMSFromTemplate(SMSTemplate *template.Template, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
