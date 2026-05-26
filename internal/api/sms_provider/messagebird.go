package sms_provider

import (
	"github.com/supabase/auth/internal/conf"
)

const (
	defaultMessagebirdApiBase = "https://rest.messagebird.com"
)

type MessagebirdProvider struct {
	Config  *conf.MessagebirdProviderConfiguration
	APIPath string
}

type MessagebirdResponseRecipients struct {
	TotalSentCount int `json:"totalSentCount"`
}

type MessagebirdResponse struct {
	ID         string                        `json:"id"`
	Recipients MessagebirdResponseRecipients `json:"recipients"`
}

type MessagebirdError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter"`
}

type MessagebirdErrResponse struct {
	Errors []MessagebirdError `json:"errors"`
}

func (t MessagebirdErrResponse) Error() string { _ = "STUB: not implemented"; return "" }

// Creates a SmsProvider with the Messagebird Config
func NewMessagebirdProvider(config conf.MessagebirdProviderConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func (t *MessagebirdProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Send an SMS containing the OTP with Messagebird's API
func (t *MessagebirdProvider) SendSms(phone string, message string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validate sms status

func (t *MessagebirdProvider) VerifyOTP(phone, code string) error {
	_ = "STUB: not implemented"
	return nil
}
