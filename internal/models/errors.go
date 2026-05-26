package models

import "errors"

// sentinel error for all not found errors.
var errNotFound = errors.New("not found")

// sentinel error for unique constraint violations.
var errUniqueConstraintViolated = errors.New("unique constraint violated")

// IsNotFoundError returns whether an error represents a "not found" error.
func IsNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

type SessionNotFoundError struct{}

func (e SessionNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e SessionNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// UserNotFoundError represents when a user is not found.
type UserNotFoundError struct{}

func (e UserNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e UserNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// IdentityNotFoundError represents when an identity is not found.
type IdentityNotFoundError struct{}

func (e IdentityNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e IdentityNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// ConfirmationOrRecoveryTokenNotFoundError represents when a confirmation or recovery token is not found.
type ConfirmationOrRecoveryTokenNotFoundError struct{}

func (e ConfirmationOrRecoveryTokenNotFoundError) Error() string {
	_ = "STUB: not implemented"
	return ""
}

func (e ConfirmationOrRecoveryTokenNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false

	// ConfirmationTokenNotFoundError represents when a confirmation token is not found.
}

type ConfirmationTokenNotFoundError struct{}

func (e ConfirmationTokenNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ConfirmationTokenNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false

	// RefreshTokenNotFoundError represents when a refresh token is not found.
}

type RefreshTokenNotFoundError struct{}

func (e RefreshTokenNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e RefreshTokenNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// FactorNotFoundError represents when a user is not found.
type FactorNotFoundError struct{}

func (e FactorNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e FactorNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// ChallengeNotFoundError represents when a user is not found.
type ChallengeNotFoundError struct{}

func (e ChallengeNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ChallengeNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// SSOProviderNotFoundError represents an error when a SSO Provider can't be
// found.
type SSOProviderNotFoundError struct{}

func (e SSOProviderNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e SSOProviderNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// SAMLRelayStateNotFoundError represents an error when a SAML relay state
// can't be found.
type SAMLRelayStateNotFoundError struct{}

func (e SAMLRelayStateNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e SAMLRelayStateNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// FlowStateNotFoundError represents an error when an FlowState can't be
// found.
type FlowStateNotFoundError struct{}

func (e FlowStateNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e FlowStateNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func IsUniqueConstraintViolatedError(err error) bool { _ = "STUB: not implemented"; return false }

type UserEmailUniqueConflictError struct{}

func (e UserEmailUniqueConflictError) Error() string { _ = "STUB: not implemented"; return "" }

func (e UserEmailUniqueConflictError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false
}

type OAuthClientStateNotFoundError struct{}

func (e OAuthClientStateNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e OAuthClientStateNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false

	// CustomOAuthProviderNotFoundError represents an error when a custom OAuth/OIDC provider can't be found
}

type CustomOAuthProviderNotFoundError struct{}

func (e CustomOAuthProviderNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CustomOAuthProviderNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false

	// WebAuthnCredentialNotFoundError represents when a WebAuthn credential can't be found.
}

type WebAuthnCredentialNotFoundError struct{}

func (e WebAuthnCredentialNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e WebAuthnCredentialNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false

	// WebAuthnChallengeNotFoundError represents when a WebAuthn challenge can't be found.
}

type WebAuthnChallengeNotFoundError struct{}

func (e WebAuthnChallengeNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e WebAuthnChallengeNotFoundError) Is(target error) bool {
	_ = "STUB: not implemented"
	return false
}
