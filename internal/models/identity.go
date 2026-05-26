package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type Identity struct {
	// returned as identity_id in JSON for backward compatibility with the interface exposed by the client library
	// see https://github.com/supabase/gotrue-js/blob/c9296bbc27a2f036af55c1f33fca5930704bd021/src/lib/types.ts#L230-L240
	ID uuid.UUID `json:"identity_id" db:"id"`
	// returned as id in JSON for backward compatibility with the interface exposed by the client library
	// see https://github.com/supabase/gotrue-js/blob/c9296bbc27a2f036af55c1f33fca5930704bd021/src/lib/types.ts#L230-L240
	ProviderID   string             `json:"id" db:"provider_id"`
	UserID       uuid.UUID          `json:"user_id" db:"user_id"`
	IdentityData JSONMap            `json:"identity_data,omitempty" db:"identity_data"`
	Provider     string             `json:"provider" db:"provider"`
	LastSignInAt *time.Time         `json:"last_sign_in_at,omitempty" db:"last_sign_in_at"`
	CreatedAt    time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" db:"updated_at"`
	Email        storage.NullString `json:"email,omitempty" db:"email" rw:"r"`
}

func (Identity) TableName() string { _ = "STUB: not implemented"; return "" }

// GetEmail returns the user's email as a string
func (i *Identity) GetEmail() string { _ = "STUB: not implemented"; return "" }

// NewIdentity returns an identity associated to the user's id.
func NewIdentity(user *User, provider string, identityData map[string]interface{}) (*Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Identity) BeforeCreate(tx *pop.Connection) error { _ = "STUB: not implemented"; return nil }

func (i *Identity) BeforeUpdate(tx *pop.Connection) error { _ = "STUB: not implemented"; return nil }

func (i *Identity) IsForSSOProvider() bool { _ = "STUB: not implemented"; return false }

// FindIdentityById searches for an identity with the matching id and provider given.
func FindIdentityByIdAndProvider(tx *storage.Connection, providerId, provider string) (*Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindIdentitiesByUserID returns all identities associated to a user ID.
func FindIdentitiesByUserID(tx *storage.Connection, userID uuid.UUID) ([]*Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindProvidersByUser returns all providers associated to a user
func FindProvidersByUser(tx *storage.Connection, user *User) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateIdentityData sets all identity_data from a map of updates,
// ensuring that it doesn't override attributes that are not
// in the provided map.
func (i *Identity) UpdateIdentityData(tx *storage.Connection, updates map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// pop doesn't support updates on tables with composite primary keys so we use a raw query here.
