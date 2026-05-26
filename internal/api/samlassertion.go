package api

import (
	"time"

	"github.com/crewjam/saml"
	"github.com/supabase/auth/internal/models"
)

type SAMLAssertion struct {
	*saml.Assertion
}

const (
	SAMLSubjectIDAttributeName = "urn:oasis:names:tc:SAML:attribute:subject-id"
)

// Attribute returns the first matching attribute value in the attribute
// statements where name equals the official SAML attribute Name or
// FriendlyName. Returns nil if such an attribute can't be found.
func (a *SAMLAssertion) Attribute(name string) []saml.AttributeValue {
	_ = "STUB: not implemented"
	return nil
}

// UserID returns the best choice for a persistent user identifier on the
// Identity Provider side. Don't assume the format of the string returned, as
// it's Identity Provider specific.
func (a *SAMLAssertion) UserID() string {
	_ = "STUB: not implemented"
	// First we look up the SAMLSubjectIDAttributeName in the attribute
	// section of the assertion, as this is the preferred way to
	// persistently identify users in SAML 2.0.
	// See: https://docs.oasis-open.org/security/saml-subject-id-attr/v1.0/cs01/saml-subject-id-attr-v1.0-cs01.html#_Toc536097226
	return ""
}

// Otherwise, fall back to the SubjectID value.

// SubjectID returns the user identifier in present in the Subject section of
// the SAML assertion. Note that this way of identifying the Subject is
// generally superseded by the SAMLSubjectIDAttributeName assertion attribute;
// tho must be present in all assertions. It can have a few formats, of which
// the most important are: saml.EmailAddressNameIDFormat (meaning the user ID
// is an email address), saml.PersistentNameIDFormat (the user ID is an opaque
// string that does not change with each assertion, e.g. UUID),
// saml.TransientNameIDFormat (the user ID changes with each assertion -- can't
// be used to identify a user). The boolean returned identifies if the user ID
// is persistent. If it's an email address, it's lowercased just in case.
func (a *SAMLAssertion) SubjectID() (string, bool) { _ = "STUB: not implemented"; return "", false }

// all other NameID formats are regarded as persistent

// Email returns the best guess for an email address.
func (a *SAMLAssertion) Email() string { _ = "STUB: not implemented"; return "" }

// Process processes this assertion according to the SAMLAttributeMapping. Never returns nil.
func (a *SAMLAssertion) Process(mapping models.SAMLAttributeMapping) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// NotBefore extracts the time before which this assertion should not be
// considered.
func (a *SAMLAssertion) NotBefore() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NotAfter extracts the time at which or after this assertion should not be
// considered.
func (a *SAMLAssertion) NotAfter() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
