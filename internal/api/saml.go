package api

import (
	"net/http"

	"github.com/crewjam/saml"
)

// getSAMLServiceProvider generates a new service provider object with the
// (optionally) provided descriptor (metadata) for the identity provider.
func (a *API) getSAMLServiceProvider(identityProvider *saml.EntityDescriptor, idpInitiated bool) *saml.ServiceProvider {
	_ = "STUB: not implemented"
	return nil
}

// this should not fail as a.config should have been validated using #Validate()

// this should not fail as a.config should have been validated using #Validate()

// SAMLMetadata serves GoTrue's SAML Service Provider metadata file.
func (a *API) SAMLMetadata(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// 5 year expiration, comparable to what GSuite does

// we set this to false since the IdP initiated flow can only
// sign the Assertion, and not the full Request
// unfortunately this is hardcoded in the crewjam library if
// signatures (instead of encryption) are supported
// https://github.com/crewjam/saml/blob/v0.4.8/service_provider.go#L217

// advertize the requested NameID formats (either persistent or email address)

// only advertize key as usable for encryption if allowed

// cache at CDN for 10 minutes

// #nosec G705
