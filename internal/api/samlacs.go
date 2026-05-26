package api

import (
	"context"
	"net/http"

	"github.com/crewjam/saml"
	"github.com/supabase/auth/internal/models"
)

func (a *API) samlDestroyRelayState(ctx context.Context, relayState *models.SAMLRelayState) error {
	_ = "STUB: not implemented"
	return nil

	// It's OK to destroy the RelayState, as a user will
	// likely initiate a completely new login flow, instead
	// of reusing the same one.
}

func IsSAMLMetadataStale(idpMetadata *saml.EntityDescriptor, samlProvider models.SAMLProvider) bool {
	_ = "STUB: not implemented"
	return false
}

// if metadata XML does not publish validity or caching information, update once in 24 hours

func (a *API) SamlAcs(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// handleSamlAcs implements the main Assertion Consumer Service endpoint behavior.
func (a *API) handleSamlAcs(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// relay state is a valid UUID, therefore this is likely a SP initiated flow

// TODO: add abuse detection to bind the RelayState UUID with a
// HTTP-Only cookie

// RelayState may be a URL in which case it's the URL where the
// IdP is telling us to redirect the user to

// SAML Artifact responses are possible only when
// RelayState can be used to identify the Identity
// Provider.

// #nosec G709

// RelayState can't be identified, so SAML flow can't continue

// Fail silently but raise warning and continue with existing metadata

// mapping does not identify the email attribute, try to figure it out

// remove all of the parsed claims, so that the rest can go into CustomClaims

// userProvidedData.Provider.Type = "saml"
// userProvidedData.Provider.ID = ssoProvider.ID.String()
// userProvidedData.Provider.SAMLEntityID = ssoProvider.SAMLProvider.EntityID
// userProvidedData.Provider.SAMLInitiatedBy = initiatedBy

// TODO: below
// refreshTokenParams.SSOProviderID = ssoProvider.ID
// refreshTokenParams.InitiatedByProvider = initiatedBy == "idp"
// refreshTokenParams.NotBefore = assertion.NotBefore()
// refreshTokenParams.NotAfter = assertion.NotAfter()

// accounts potentially created via SAML can contain non-unique email addresses in the auth.users table

// PKCE flow: update flow state with user ID
// Re-fetch with FOR UPDATE lock inside the transaction to prevent concurrent claims

// PKCE flow: redirect with auth code

// #nosec G710

// Record login for analytics - only when token is issued (not during pkce authorize)

// #nosec G710
