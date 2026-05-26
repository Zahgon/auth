package siwe

import (
	"net/url"
	"regexp"
	"time"
)

// SIWEMessage is the final structured form of a parsed SIWE message.
// REF: https://eips.ethereum.org/EIPS/eip-4361
type SIWEMessage struct {
	Raw string

	Domain         string
	Address        string
	Statement      *string
	URI            url.URL
	Version        string
	ChainID        string
	Nonce          string
	IssuedAt       time.Time
	ExpirationTime *time.Time
	NotBefore      *time.Time
	RequestID      *string
	Resources      []*url.URL
}

const headerSuffix = " wants you to sign in with your Ethereum account:"

var addressPattern = regexp.MustCompile("^0x[a-fA-F0-9]{40}$")

func ParseMessage(raw string) (*SIWEMessage, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse first line exactly

// this is supposed to be REQUIRED >8 chr alphanum but we'll leave it for now for gotrue's nonce impl

// This is supposed to be a pchar (RFC 3986) but generally we'll keep it as any str for now

// VerifySignature validates that the signature was created by the private key
// corresponding to the address in the message. This performs ECDSA recovery
// which is computationally expensive, so it should be called only after
// ParseMessage has validated the message structure.
//
// The signature must be a 65-byte hex string in the format: 0x{R}{S}{V}
// where R and S are 32 bytes each and V is 1 byte.
//
// Returns true if the recovered address matches the message address (case-insensitive).
func (m *SIWEMessage) VerifySignature(signatureHex string) bool {
	_ = "STUB: not implemented"
	return false
}

// Create signature in [R || S || V] format

// Normalize V if needed
// #nosec G602

// #nosec G602

// Recover public key

// Convert to address
