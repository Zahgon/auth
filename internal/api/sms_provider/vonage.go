package sms_provider

import (
	"github.com/supabase/auth/internal/conf"
)

const (
	defaultVonageApiBase = "https://rest.nexmo.com"
)

type VonageProvider struct {
	Config  *conf.VonageProviderConfiguration
	APIPath string
}

type VonageResponseMessage struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	ErrorText string `json:"error-text"`
}

type VonageResponse struct {
	Messages []VonageResponseMessage `json:"messages"`
}

// Creates a SmsProvider with the Vonage Config
func NewVonageProvider(config conf.VonageProviderConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func (t *VonageProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Send an SMS containing the OTP with Vonage's API
func (t *VonageProvider) SendSms(phone string, message string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// A status of zero indicates success; a non-zero value means something went wrong.

func (t *VonageProvider) VerifyOTP(phone, code string) error { _ = "STUB: not implemented"; return nil }
