package api

import (
	"github.com/supabase/auth/internal/api/apilimiter"
	"github.com/supabase/auth/internal/mailer"
	"github.com/supabase/auth/internal/security"
	"github.com/supabase/auth/internal/tokens"
)

type Option interface {
	apply(*API)
}

type optionFunc func(*API)

func (f optionFunc) apply(a *API) { _ = "STUB: not implemented"; return }

func WithMailer(m mailer.Mailer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTokenService(service *tokens.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCaptchaVerifier(v security.CaptchaVerifier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLimiter(v *apilimiter.Limiter) Option { _ = "STUB: not implemented"; return *new(Option) }
