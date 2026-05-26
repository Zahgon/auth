package api

import (
	"context"
	"net/http"

	"github.com/crewjam/saml"
	"github.com/supabase/auth/internal/models"
)

// loadSSOProvider looks for an idp_id and first checks it for a "resource_"
// prefix, if present the provider is loaded by resource_id. Otherwise the
// provider is loaded by id.
func (a *API) loadSSOProvider(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// adminSSOProvidersList lists all SSO Identity Providers in the system. Does
// not deal with pagination at this time.
func (a *API) adminSSOProvidersList(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// remove metadata XML so that the returned JSON is not ginormous

type CreateSSOProviderParams struct {
	Type string `json:"type"`

	MetadataURL      string                      `json:"metadata_url"`
	MetadataXML      string                      `json:"metadata_xml"`
	Domains          []string                    `json:"domains"`
	AttributeMapping models.SAMLAttributeMapping `json:"attribute_mapping"`
	NameIDFormat     string                      `json:"name_id_format"`

	ResourceID *string `json:"resource_id,omitempty"`
	Disabled   *bool   `json:"disabled,omitempty"`
}

func (p *CreateSSOProviderParams) validate(forUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// it's valid

func (p *CreateSSOProviderParams) metadata(ctx context.Context) ([]byte, *saml.EntityDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// impossible situation if you called validate() prior

func parseSAMLMetadata(rawMetadata []byte) (*saml.EntityDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchSAMLMetadata(ctx context.Context, url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// adminSSOProvidersCreate creates a new SAML Identity Provider in the system.
func (a *API) adminSSOProvidersCreate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

/* <- forUpdate */

// TODO handle Name, Description, Attribute Mapping

// adminSSOProvidersGet returns an existing SAML Identity Provider in the system.
func (a *API) adminSSOProvidersGet(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// adminSSOProvidersUpdate updates a provider with the provided diff values.
func (a *API) adminSSOProvidersUpdate(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

/* <- forUpdate */

// metadata is being updated

// domains are being "updated" only when params.Domains is not nil, if
// it was nil (but not `[]`) then the caller is expecting not to modify
// the domains

// adminSSOProvidersDelete deletes a SAML identity provider.
func (a *API) adminSSOProvidersDelete(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
