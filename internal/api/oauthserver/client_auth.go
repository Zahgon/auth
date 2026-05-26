package oauthserver

import (
	"github.com/supabase/auth/internal/models"
)

// InferClientTypeFromAuthMethod infers client type from token_endpoint_auth_method
func InferClientTypeFromAuthMethod(authMethod string) string { _ = "STUB: not implemented"; return "" }

// Default to confidential

// GetValidAuthMethodsForClientType returns the valid authentication methods for a client type
func GetValidAuthMethodsForClientType(clientType string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Unknown client type

// ValidateClientTypeConsistency validates consistency between client_type and token_endpoint_auth_method
func ValidateClientTypeConsistency(clientType, authMethod string) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip validation if either is not provided

// IsValidAuthMethodForClientType checks if the auth method is valid for the given client type
func IsValidAuthMethodForClientType(clientType, authMethod string) bool {
	_ = "STUB: not implemented"
	return false
}

// DetermineClientType determines the final client type using the priority:
// 1. Explicit client_type
// 2. Inferred from token_endpoint_auth_method
// 3. Default to confidential
func DetermineClientType(explicitClientType, authMethod string) string {
	_ = "STUB: not implemented"
	// Priority 1: Explicit client_type
	return ""
}

// Priority 2: Infer from token_endpoint_auth_method

// Priority 3: Default to confidential

// ValidateClientAuthentication validates client authentication based on client type
func ValidateClientAuthentication(client *models.OAuthServerClient, providedSecret string) error {
	_ = "STUB: not implemented"
	return nil

	// Public clients should not provide client secrets
}

// Confidential clients must provide a valid client secret

// GetAllValidAuthMethods returns all supported authentication methods
func GetAllValidAuthMethods() []string { _ = "STUB: not implemented"; return nil }

// ValidateClientAuthMethod validates the authentication method used matches the registered method
func ValidateClientAuthMethod(client *models.OAuthServerClient, usedMethod string) error {
	_ = "STUB: not implemented"
	return nil
}
