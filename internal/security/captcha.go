package security

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/conf"
)

type VerificationResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
	Hostname   string   `json:"hostname"`
}

// CaptchaVerifier abstracts CAPTCHA verification for different providers (hCaptcha, Cloudflare Turnstile, etc.)
// and allows for mocking in tests.
type CaptchaVerifier interface {
	Verify(ctx context.Context, token, clientIP string) (*VerificationResponse, error)
}

// HTTPCaptchaVerifier is the default implementation that calls out to hCaptcha / Turnstile.
type HTTPCaptchaVerifier struct {
	client   *http.Client
	secret   string
	provider string
}

func NewCaptchaVerifier(cfg *conf.CaptchaConfiguration) *HTTPCaptchaVerifier {
	_ = "STUB: not implemented"
	return nil
}

func (v *HTTPCaptchaVerifier) Verify(ctx context.Context, token, clientIP string) (*VerificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *HTTPCaptchaVerifier) verifyCaptchaCode(ctx context.Context, token, clientIP, captchaURL string) (*VerificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO (darora): pipe through sitekey

func getCaptchaURL(captchaProvider string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
