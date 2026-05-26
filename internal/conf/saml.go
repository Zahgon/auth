package conf

import (
	"crypto/rsa"
	"crypto/x509"
	"time"
)

// SAMLConfiguration holds configuration for native SAML support.
type SAMLConfiguration struct {
	Enabled                  bool          `json:"enabled"`
	PrivateKey               string        `json:"-" split_words:"true"`
	AllowEncryptedAssertions bool          `json:"allow_encrypted_assertions" split_words:"true"`
	RelayStateValidityPeriod time.Duration `json:"relay_state_validity_period" split_words:"true"`

	RSAPrivateKey *rsa.PrivateKey   `json:"-"`
	RSAPublicKey  *rsa.PublicKey    `json:"-"`
	Certificate   *x509.Certificate `json:"-"`

	ExternalURL string `json:"external_url,omitempty" split_words:"true"`

	RateLimitAssertion float64 `default:"15" split_words:"true"`
}

func (c *SAMLConfiguration) GoString() string { _ = "STUB: not implemented"; return "" }
func (c *SAMLConfiguration) String() string   { _ = "STUB: not implemented"; return "" }

func (c *SAMLConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

// PopulateFields fills the configuration details based off the provided
// parameters.
func (c *SAMLConfiguration) PopulateFields(externalURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// PopulateFields fills the configuration details based off the provided
// parameters.
func (c *SAMLConfiguration) populateFields(externalURL string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	// errors are intentionally ignored since they should have been handled
	// within #Validate()
	return nil, nil
}

// SAML does not care much about the contents of the certificate, it
// only uses it as a vessel for the public key; therefore we set these
// fixed values.
// Please avoid modifying or adding new values to this template as they
// will change the exposed SAML certificate, requiring users of
// GoTrue to re-establish a connection between their Identity Provider
// and their running GoTrue instances.

func (c *SAMLConfiguration) createCertificate(certTemplate *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SAMLConfiguration) parseCertificateDer(certDer []byte) error {
	_ = "STUB: not implemented"
	return nil
}
