package siws

import (
	"regexp"
)

var domainPattern = regexp.MustCompile(`^(localhost|(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,})(?::\d{1,5})?$`)

func IsValidDomain(domain string) bool { _ = "STUB: not implemented"; return false }

var validSolanaNetworksPattern = regexp.MustCompile("^solana:(main|dev|test|local)net$")

func IsValidSolanaNetwork(network string) bool { _ = "STUB: not implemented"; return false }
