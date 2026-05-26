package sms_provider

import (
	"github.com/supabase/auth/internal/conf"
)

const (
	defaultTextLocalApiBase    = "https://api.textlocal.in"
	textLocalTemplateErrorCode = 80
)

type TextlocalProvider struct {
	Config  *conf.TextlocalProviderConfiguration
	APIPath string
}

type TextlocalError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TextlocalResponse struct {
	Status   string             `json:"status"`
	Errors   []TextlocalError   `json:"errors"`
	Messages []TextlocalMessage `json:"messages"`
}

type TextlocalMessage struct {
	MessageID string `json:"id"`
}

// Creates a SmsProvider with the Textlocal Config
func NewTextlocalProvider(config conf.TextlocalProviderConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func (t *TextlocalProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Send an SMS containing the OTP with Textlocal's API
func (t *TextlocalProvider) SendSms(phone string, message string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TextlocalProvider) VerifyOTP(phone, code string) error {
	_ = "STUB: not implemented"
	return nil
}
