package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

const OAuthClientStateTimeout = 5 * time.Minute

type OAuthClientState struct {
	ID           uuid.UUID `json:"id" db:"id"`
	ProviderType string    `json:"provider_type" db:"provider_type"`
	CodeVerifier *string   `json:"code_verifier,omitempty" db:"code_verifier"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

func (OAuthClientState) TableName() string { _ = "STUB: not implemented"; return "" }

func NewOAuthClientState(providerType string, codeVerifier *string) *OAuthClientState {
	_ = "STUB: not implemented"
	return nil
}

func FindAndDeleteOAuthClientStateByID(tx *storage.Connection, id uuid.UUID) (*OAuthClientState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *OAuthClientState) IsExpired() bool { _ = "STUB: not implemented"; return false }
