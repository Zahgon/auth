package oauthserver

import (
	"net/http"
)

// ClientCredentials represents the extracted client credentials and authentication method used
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
	AuthMethod   string
}

// ExtractClientCredentials extracts OAuth client credentials from the request
// Supports Basic auth header, form body parameters, and JSON body parameters
func ExtractClientCredentials(r *http.Request) (*ClientCredentials, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// First, try Basic auth header: Authorization: Basic base64(client_id:client_secret)
}

// Check Content-Type to determine how to parse body parameters

// Parse JSON body

// Restore the body so other handlers can read it

// Fall back to form parameters

// return error if client_id is not provided

// Determine auth method based on presence of client_secret in body
