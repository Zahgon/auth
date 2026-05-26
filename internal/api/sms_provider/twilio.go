package sms_provider

import (
	"regexp"

	"github.com/supabase/auth/internal/conf"
)

const (
	defaultTwilioApiBase = "https://api.twilio.com"
	apiVersion           = "2010-04-01"
)

type TwilioProvider struct {
	Config  *conf.TwilioProviderConfiguration
	APIPath string
}

var isPhoneNumber = regexp.MustCompile("^[1-9][0-9]{1,14}$")

// formatPhoneNumber removes "+" and whitespaces in a phone number
func formatPhoneNumber(phone string) string { _ = "STUB: not implemented"; return "" }

type SmsStatus struct {
	To           string `json:"to"`
	From         string `json:"from"`
	MessageSID   string `json:"sid"`
	Status       string `json:"status"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         string `json:"body"`
}

type twilioErrResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (t twilioErrResponse) Error() string { _ = "STUB: not implemented"; return "" }

// Creates a SmsProvider with the Twilio Config
func NewTwilioProvider(config conf.TwilioProviderConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func (t *TwilioProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Send an SMS containing the OTP with Twilio's API
func (t *TwilioProvider) SendSms(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// twilio api requires "+" extension to be included

// Programmable Messaging (WhatsApp) takes in different set of inputs

// twilio api requires "+" extension to be included

// For backward compatibility with old API.

// Used to substitute OTP. See https://www.twilio.com/docs/content/whatsappauthentication for more details

// validate sms status

func (t *TwilioProvider) VerifyOTP(phone, code string) error { _ = "STUB: not implemented"; return nil }
