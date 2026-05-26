package models

// OAuth/OIDC scope constants
const (
	ScopeOpenID  = "openid"
	ScopeEmail   = "email"
	ScopeProfile = "profile"
	ScopePhone   = "phone"
)

// SupportedOAuthScopes defines all OAuth/OIDC scopes supported by the server
var SupportedOAuthScopes = []string{
	ScopeOpenID,
	ScopeProfile,
	ScopeEmail,
	ScopePhone,
}

// IsSupportedScope checks if a scope is in the supported scopes list
func IsSupportedScope(scope string) bool { _ = "STUB: not implemented"; return false }

// ParseScopeString parses a space-separated scope string into a slice
func ParseScopeString(scopeString string) []string { _ = "STUB: not implemented"; return nil }

// Always return empty slice instead of nil for consistency

// HasScope checks if the given scope list includes a specific scope
func HasScope(scopes []string, scope string) bool { _ = "STUB: not implemented"; return false }

// HasAllScopes checks if the granted scopes include all of the requested scopes
func HasAllScopes(grantedScopes, requestedScopes []string) bool {
	_ = "STUB: not implemented"
	return false
}
