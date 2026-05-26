package api

import (
	"context"
)

// BCrypt hashed passwords have a 72 character limit
const MaxPasswordLength = 72

// WeakPasswordError encodes an error that a password does not meet strength
// requirements. It is handled specially in errors.go as it gets transformed to
// a HTTPError with a special weak_password field that encodes the Reasons
// slice.
type WeakPasswordError struct {
	Message string   `json:"message,omitempty"`
	Reasons []string `json:"reasons,omitempty"`
}

func (e *WeakPasswordError) Error() string { _ = "STUB: not implemented"; return "" }

func (a *API) checkPasswordStrength(ctx context.Context, password string) error {
	_ = "STUB: not implemented"
	return nil
}
