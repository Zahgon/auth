package sbff

import (
	"errors"
	"net/http"

	"github.com/supabase/auth/internal/conf"
)

// HeaderName is the Sb-Forwarded-For header name. It is all lowercase here as HTTP header names
// are not case-sensitive.
const HeaderName = "sb-forwarded-for"

var (
	ctxKeySBFF = &struct{}{}

	ErrHeaderNotFound = errors.New("Sb-Forwarded-For header not found")
	ErrHeaderInvalid  = errors.New("invalid Sb-Forwarded-For header value")
)

func parseSBFFHeader(headerVal string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetIPAddress returns the value of the IP address in Sb-Forwarded-For as defined by
// SBForwardedForMiddleware. If no value is present in the request context, this function will
// return ("", false).
func GetIPAddress(r *http.Request) (addr string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

// withIPAddress parses the Sb-Forwarded-For header and adds the leftmost value to the
// request context if it is a valid IP address, then returns a new request with modified context.
// If the leftmost value is not a valid IP address or the header is not set, this function returns
// an error.
func withIPAddress(r *http.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Middleware returns a middleware function that parses the Sb-Forwarded-For header
// and adds the leftmost header value to the request context if GOTRUE_SECURITY_SB_FORWARDED_FOR_ENABLED
// is true and the value is a valid IP address.
func Middleware(cfg *conf.SecurityConfiguration, errCallback func(*http.Request, error)) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
