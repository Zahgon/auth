package sms_provider

import (
	"github.com/supabase/auth/internal/conf"
)

const (
	verifyServiceApiBase = "https://verify.twilio.com/v2/Services/"
)

type TwilioVerifyProvider struct {
	Config  *conf.TwilioVerifyProviderConfiguration
	APIPath string
}

type VerificationResponse struct {
	To              string `json:"to"`
	Status          string `json:"status"`
	Channel         string `json:"channel"`
	Valid           bool   `json:"valid"`
	VerificationSID string `json:"sid"`
	ErrorCode       string `json:"error_code"`
	ErrorMessage    string `json:"error_message"`
}

// See: https://www.twilio.com/docs/verify/api/verification-check
type VerificationCheckResponse struct {
	To           string `json:"to"`
	Status       string `json:"status"`
	Channel      string `json:"channel"`
	Valid        bool   `json:"valid"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// Creates a SmsProvider with the Twilio Config
func NewTwilioVerifyProvider(config conf.TwilioVerifyProviderConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func (t *TwilioVerifyProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Send an SMS containing the OTP with Twilio's API
func (t *TwilioVerifyProvider) SendSms(phone, message, channel string) (string, error) {
	_ = "STUB: not implemented"
	// Unlike Programmable Messaging, Verify does not require a prefix for channel
	return "", nil
}

func (t *TwilioVerifyProvider) VerifyOTP(phone, code string) error {
	_ = "STUB: not implemented"
	return nil
}

// twilio api requires "+" extension to be included
