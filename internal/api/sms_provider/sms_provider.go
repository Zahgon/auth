package sms_provider

import (
	"log"
	"os"
	"time"

	"github.com/supabase/auth/internal/conf"
)

// overrides the SmsProvider set to always return the mock provider
var MockProvider SmsProvider = nil

var defaultTimeout time.Duration = time.Second * 10

const SMSProvider = "sms"
const WhatsappProvider = "whatsapp"

func init() {
	timeoutStr := os.Getenv("GOTRUE_INTERNAL_HTTP_TIMEOUT")
	if timeoutStr != "" {
		if timeout, err := time.ParseDuration(timeoutStr); err != nil {
			log.Fatalf("error loading GOTRUE_INTERNAL_HTTP_TIMEOUT: %v", err.Error()) // #nosec G706
		} else if timeout != 0 {
			defaultTimeout = timeout
		}
	}
}

type SmsProvider interface {
	SendMessage(phone, message, channel, otp string) (string, error)
	VerifyOTP(phone, token string) error
}

func GetSmsProvider(config conf.GlobalConfiguration) (SmsProvider, error) {
	_ = "STUB: not implemented"
	return *new(SmsProvider), nil
}

func IsValidMessageChannel(channel string, config *conf.GlobalConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// channel doesn't matter if SMS hook is enabled
