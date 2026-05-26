package siwe

import (
	"regexp"
)

var domainPattern = regexp.MustCompile(`^(localhost|(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,})(?::\d{1,5})?$`)

func IsValidDomain(domain string) bool { _ = "STUB: not implemented"; return false }

func isValidEthereumNetwork(network string) bool {
	_ = "STUB: not implemented"
	// REF: https://eips.ethereum.org/EIPS/eip-155
	// parse the network as an int first (not a string)
	return false
}
