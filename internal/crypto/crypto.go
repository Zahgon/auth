package crypto

import (
	"io"
)

// GenerateOtp generates a random n digit otp
func GenerateOtp(digits int) string { _ = "STUB: not implemented"; return "" }

func generateOtp(r io.Reader, digits int) string {
	_ = "STUB: not implemented"
	// TODO(cstockton): Change the code to be below and propagate errors so we
	// can have non-panicing bounds checks. This is just a defensive change so
	// if someone changes OTP length in the future we don't end up with an
	// overflowed float64 / panic.
	//
	//	upper := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)
	//	val := must(rand.Int(r, upper))
	return ""
}

// adds a variable zero-padding to the left to ensure otp is uniformly random

func GenerateTokenHash(emailOrPhone, otp string) string { _ = "STUB: not implemented"; return "" }

// Generated a random secure integer from [0, max[
func secureRandomInt(max int) int { _ = "STUB: not implemented"; return 0 }

type EncryptedString struct {
	KeyID     string `json:"key_id"`
	Algorithm string `json:"alg"`
	Data      []byte `json:"data"`
	Nonce     []byte `json:"nonce,omitempty"`
}

func (es *EncryptedString) IsValid() bool {
	_ = "STUB: not implemented"
	// cipher.NewGCM() is always used, which panics on a Nonce that is not 12 bytes
	// enforce that the nonce length and other values are correct
	return false
}

// ShouldReEncrypt tells you if the value encrypted needs to be encrypted again with a newer key.
func (es *EncryptedString) ShouldReEncrypt(encryptionKeyID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (es *EncryptedString) Decrypt(id string, decryptionKeys map[string]string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G407

func ParseEncryptedString(str string) *EncryptedString { _ = "STUB: not implemented"; return nil }

func (es *EncryptedString) String() string { _ = "STUB: not implemented"; return "" }

func deriveSymmetricKey(id, keyID, keyBase64URL string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Since we use AES-GCM here, the same symmetric key *must not be used
// more than* 2^32 times. But, that's not that much. Suppose a system
// with 100 million users, then a user can only change their password
// 42 times. To prevent this, the actual symmetric key is derived by
// using HKDF using the encryption key and the "ID" of the object
// containing the encryption string. Ideally this ID is a UUID.  This
// has the added benefit that the encrypted string is bound to that
// specific object, and can't accidentally be "moved" to other objects
// without changing their ID to the original one.

func NewEncryptedString(id string, data []byte, keyID string, keyBase64URL string) (*EncryptedString, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G407

// SecureAlphanumeric generates a secure random alphanumeric string using standard library
func SecureAlphanumeric(length int) string { _ = "STUB: not implemented"; return "" }

// Calculate bytes needed for desired length
// base32 encoding: 5 bytes -> 8 chars

// Use standard library's base32 without padding
