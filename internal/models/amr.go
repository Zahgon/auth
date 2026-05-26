package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/storage"
)

type AMRClaim struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	SessionID            uuid.UUID `json:"session_id" db:"session_id"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	AuthenticationMethod *string   `json:"authentication_method" db:"authentication_method"`
}

func (AMRClaim) TableName() string { _ = "STUB: not implemented"; return "" }

func (cl *AMRClaim) IsAAL2Claim() bool { _ = "STUB: not implemented"; return false }

func AddClaimToSession(tx *storage.Connection, sessionId uuid.UUID, authenticationMethod AuthenticationMethod) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AMRClaim) GetAuthenticationMethod() string { _ = "STUB: not implemented"; return "" }
