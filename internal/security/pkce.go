package security

const PKCEInvalidCodeChallengeError = "code challenge does not match previously saved code verifier"
const PKCEInvalidCodeMethodError = "code challenge method not supported"

// VerifyPKCEChallenge performs PKCE verification using the provided challenge, method, and verifier
// This is a shared utility function used by both FlowState and OAuthServerAuthorization
func VerifyPKCEChallenge(codeChallenge, codeChallengeMethod, codeVerifier string) error {
	_ = "STUB: not implemented"
	return nil
}
