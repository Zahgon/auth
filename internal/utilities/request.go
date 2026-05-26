package utilities

import (
	"net/http"
	"regexp"

	"github.com/supabase/auth/internal/conf"
)

func getIPAddressWithXFF(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// GetIPAddress returns the real IP address of the HTTP request.
func GetIPAddress(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// GetBodyBytes reads the whole request body properly into a byte array.
func GetBodyBytes(req *http.Request) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GetReferrer(r *http.Request, config *conf.GlobalConfiguration) string {
	_ = "STUB: not implemented"
	// try get redirect url from query or post data first
	return ""
}

// instead try referrer header value

var decimalIPAddressPattern = regexp.MustCompile("^[0-9]+$")
var regularHostname = regexp.MustCompile("^[a-zA-Z0-9]([a-zA-Z0-9.-]*[a-zA-Z0-9])?$")

func IsRedirectURLValid(config *conf.GlobalConfiguration, redirectURL string) bool {
	_ = "STUB: not implemented"
	return false
}

// As long as the referrer came from the site, we will redirect back there

// redirect URL is for some reason invalid

// IP address in decimal form also not allowed in redirects!

// hostname uses characters that are not typically used

// For case when user came from mobile app or other permitted resource - redirect back

// only match without the fragment

// getRedirectTo tries extract redirect url from header or from query params
func getRedirectTo(r *http.Request) (reqref string) { _ = "STUB: not implemented"; return "" }
