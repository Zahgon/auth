package crypto

import (
	"errors"

	"github.com/gofrs/uuid"
)

func GenerateRefreshTokenHmacKey() []byte { _ = "STUB: not implemented"; return nil }

const refreshTokenChecksumLength = 4
const refreshTokenSignatureLength = 16
const minRefreshTokenLength = 1 + 16 + 1 + refreshTokenSignatureLength + refreshTokenChecksumLength
const maxRefreshTokenLength = minRefreshTokenLength + 8

// RefreshToken is an object that encodes a cryptographically authenticated
// (signed) message containing a version, session ID and monotonically
// increasing non-negative counter.
//
// The signature is a truncated (first 128 bits) of HMAC-SHA-256, which saves
// on encoded length without sacrificing security. The checksum of 4 bytes at
// the end is to lessen the load on the server with invalid strings (those that
// are not likely to be a proper refresh token).
type RefreshToken struct {
	Raw []byte

	Version   byte
	SessionID uuid.UUID
	Counter   int64
	Signature []byte
}

func (RefreshToken) TableName() string { _ = "STUB: not implemented"; return "" }

func (r *RefreshToken) CheckSignature(hmacSha256Key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RefreshToken) Encode(hmacSha256Key []byte) string { _ = "STUB: not implemented"; return "" }

// Note on truncating the HMAC-SHA-256 output:
// This does not impact security as the brute-force space is 2^128 and
// the collision space is 2^64, both unattainable in practice.

var (
	ErrRefreshTokenLength          = errors.New("crypto: refresh token length is not valid")
	ErrRefreshTokenUnknownVersion  = errors.New("crypto: refresh token version is not 0")
	ErrRefreshTokenChecksumInvalid = errors.New("crypto: refresh token checksum is not valid")
	ErrRefreshTokenCounterInvalid  = errors.New("crypto: refresh token's counter is not valid")
)

func safeInt64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

func safeUint64(v int64) uint64 { _ = "STUB: not implemented"; return 0 }

func ParseRefreshToken(token string) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
