package models

import (
	"database/sql/driver"
	"net/url"
	"time"

	"github.com/crewjam/saml"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type SSOProvider struct {
	ID           uuid.UUID    `db:"id" json:"id"`
	ResourceID   *string      `db:"resource_id" json:"resource_id,omitempty"`
	Disabled     *bool        `db:"disabled" json:"disabled"`
	SAMLProvider SAMLProvider `has_one:"saml_providers" fk_id:"sso_provider_id" json:"saml,omitempty"`
	SSODomains   []SSODomain  `has_many:"sso_domains" fk_id:"sso_provider_id" json:"domains"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (p SSOProvider) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (p SSOProvider) TableName() string { _ = "STUB: not implemented"; return "" }

func (p SSOProvider) Type() string { _ = "STUB: not implemented"; return "" }

type SAMLAttribute struct {
	Name    string      `json:"name,omitempty"`
	Names   []string    `json:"names,omitempty"`
	Default interface{} `json:"default,omitempty"`
	Array   bool        `json:"array,omitempty"`
}

type SAMLAttributeMapping struct {
	Keys map[string]SAMLAttribute `json:"keys,omitempty"`
}

func (m *SAMLAttributeMapping) Equal(o *SAMLAttributeMapping) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *SAMLAttributeMapping) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (m SAMLAttributeMapping) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

type SAMLProvider struct {
	ID uuid.UUID `db:"id" json:"-"`

	SSOProvider   *SSOProvider `belongs_to:"sso_providers" json:"-"`
	SSOProviderID uuid.UUID    `db:"sso_provider_id" json:"-"`

	EntityID    string  `db:"entity_id" json:"entity_id"`
	MetadataXML string  `db:"metadata_xml" json:"metadata_xml,omitempty"`
	MetadataURL *string `db:"metadata_url" json:"metadata_url,omitempty"`

	AttributeMapping SAMLAttributeMapping `db:"attribute_mapping" json:"attribute_mapping,omitempty"`

	NameIDFormat *string `db:"name_id_format" json:"name_id_format,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (p SAMLProvider) TableName() string { _ = "STUB: not implemented"; return "" }

func (p SAMLProvider) EntityDescriptor() (*saml.EntityDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SSODomain struct {
	ID uuid.UUID `db:"id" json:"-"`

	SSOProvider   *SSOProvider `belongs_to:"sso_providers" json:"-"`
	SSOProviderID uuid.UUID    `db:"sso_provider_id" json:"-"`

	Domain string `db:"domain" json:"domain"`

	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (d SSODomain) TableName() string { _ = "STUB: not implemented"; return "" }

type SAMLRelayState struct {
	ID uuid.UUID `db:"id"`

	SSOProviderID uuid.UUID `db:"sso_provider_id"`

	RequestID string  `db:"request_id"`
	ForEmail  *string `db:"for_email"`

	RedirectTo string `db:"redirect_to"`

	CreatedAt   time.Time  `db:"created_at" json:"-"`
	UpdatedAt   time.Time  `db:"updated_at" json:"-"`
	FlowStateID *uuid.UUID `db:"flow_state_id" json:"flow_state_id,omitempty"`
	FlowState   *FlowState `db:"-" json:"flow_state,omitempty" belongs_to:"flow_state"`
}

func (s SAMLRelayState) TableName() string { _ = "STUB: not implemented"; return "" }

func FindSAMLProviderByEntityID(tx *storage.Connection, entityId string) (*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSSOProviderByID(tx *storage.Connection, id uuid.UUID) (*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSSOProviderByResourceID(tx *storage.Connection, id string) (*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSSOProviderForEmailAddress(tx *storage.Connection, emailAddress string) (*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSSOProviderByDomain(tx *storage.Connection, domain string) (*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindAllSSOProviders(tx *storage.Connection) ([]SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	resourceIDFilter       = "resource_id"
	resourceIDPrefixFilter = "resource_id_prefix"
)

// FindAllSSOProvidersByFilter finds SSO Providers with the matching filter.
func FindAllSSOProvidersByFilter(
	tx *storage.Connection,
	queryValues url.Values,
) ([]*SSOProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSAMLRelayStateByID(tx *storage.Connection, id uuid.UUID) (*SAMLRelayState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
