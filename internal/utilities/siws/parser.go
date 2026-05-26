package siws

import (
	"net/url"
	"regexp"
	"time"
)

// SIWSMessage is the final structured form of a parsed SIWS message.
type SIWSMessage struct {
	Raw string

	Domain         string
	Address        string
	Statement      string
	URI            *url.URL
	Version        string
	Nonce          string
	IssuedAt       time.Time
	ChainID        string
	NotBefore      time.Time
	RequestID      string
	ExpirationTime time.Time
	Resources      []*url.URL
}

const headerSuffix = " wants you to sign in with your Solana account:"

var addressPattern = regexp.MustCompile("^[a-zA-Z0-9]{32,44}$")

func ParseMessage(raw string) (*SIWSMessage, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse first line exactly

func (m *SIWSMessage) VerifySignature(signature []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// try to verify just the signed message (in accordance with https://github.com/phantom/sign-in-with-solana

// if that didn't work, try to verify the signed message as if it was signed via Ledger (https://docs.anza.xyz/proposals/off-chain-message-signing)

// Write 16-byte prefix

// Write single-byte fields
// version

// Write domain, padded/truncated to 32 bytes

// message format = ascii
// signer num = 1

// Write pubkey

// Write message length (2 bytes, little endian)

// Write message
