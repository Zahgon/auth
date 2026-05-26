package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

type AuthenticatorAssuranceLevel int

const (
	AAL1 AuthenticatorAssuranceLevel = iota
	AAL2
	AAL3
)

func (aal AuthenticatorAssuranceLevel) String() string { _ = "STUB: not implemented"; return "" }

func (aal AuthenticatorAssuranceLevel) PointerString() *string {
	_ = "STUB: not implemented"
	return nil
}

// CompareAAL returns 0 if both AAL levels are equal, > 0 if A is a higher level than B or < 0 if A is a lower level than B.
func CompareAAL(a, b AuthenticatorAssuranceLevel) int { _ = "STUB: not implemented"; return 0 }

func ParseAAL(value *string) AuthenticatorAssuranceLevel {
	_ = "STUB: not implemented"
	return *new(AuthenticatorAssuranceLevel)
}

// AMREntry represents a method that a user has logged in together with the corresponding time
type AMREntry struct {
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
	Provider  string `json:"provider,omitempty"`
}

type Session struct {
	ID     uuid.UUID `json:"-" db:"id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`

	// NotAfter is overriden by timeboxed sessions.
	NotAfter *time.Time `json:"not_after,omitempty" db:"not_after"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	FactorID  *uuid.UUID `json:"factor_id" db:"factor_id"`
	AMRClaims []AMRClaim `json:"amr,omitempty" has_many:"amr_claims"`
	AAL       *string    `json:"aal" db:"aal"`

	RefreshedAt *time.Time `json:"refreshed_at,omitempty" db:"refreshed_at"`
	UserAgent   *string    `json:"user_agent,omitempty" db:"user_agent"`
	IP          *string    `json:"ip,omitempty" db:"ip"`

	Tag           *string    `json:"tag" db:"tag"`
	OAuthClientID *uuid.UUID `json:"oauth_client_id" db:"oauth_client_id"`
	Scopes        *string    `json:"scopes,omitempty" db:"scopes"` // OAuth scopes granted for this session

	RefreshTokenHmacKey *string `json:"-" db:"refresh_token_hmac_key"`
	RefreshTokenCounter *int64  `json:"-" db:"refresh_token_counter"`
}

func (Session) TableName() string { _ = "STUB: not implemented"; return "" }

func (s *Session) IsRecovery() bool { _ = "STUB: not implemented"; return false }

// We want to assert this is a recovery session, skip invalid.

func (s *Session) GetRefreshTokenHmacKey(dbEncryption conf.DatabaseEncryptionConfiguration) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *Session) LastRefreshedAt(refreshTokenTime *time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (s *Session) UpdateOnlyRefreshInfo(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	// TODO(kangmingtay): The underlying database type uses timestamp without timezone,
	// so we need to convert the value to UTC before updating it.
	// In the future, we should add a migration to update the type to contain the timezone.
	return nil
}

func (s *Session) UpdateOnlyRefreshToken(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Session) ReEncryptRefreshTokenHmacKey(tx *storage.Connection, dbEncryption conf.DatabaseEncryptionConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

type SessionValidityReason = int

const (
	SessionValid        SessionValidityReason = iota
	SessionPastNotAfter                       = iota
	SessionPastTimebox                        = iota
	SessionTimedOut                           = iota
	SessionLowAAL                             = iota
)

type SessionValidityConfig struct {
	Timebox           *time.Duration
	InactivityTimeout *time.Duration
	AllowLowAAL       *time.Duration
}

func (s *Session) CheckValidity(config SessionValidityConfig, now time.Time, refreshTokenTime *time.Time, userHighestPossibleAAL AuthenticatorAssuranceLevel) SessionValidityReason {
	_ = "STUB: not implemented"
	return *new(SessionValidityReason)
}

func (s *Session) DetermineTag(tags []string) string { _ = "STUB: not implemented"; return "" }

func NewSession(userID uuid.UUID, factorID *uuid.UUID) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindSessionByID looks up a Session by the provided id. If forUpdate is set
// to true, then the SELECT statement used by the query has the form SELECT ...
// FOR UPDATE SKIP LOCKED. This means that a FOR UPDATE lock will only be
// acquired if there's no other lock. In case there is a lock, a
// IsNotFound(err) error will be retured.
func FindSessionByID(tx *storage.Connection, id uuid.UUID, forUpdate bool) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pop does not provide us with a way to execute FOR UPDATE
// queries which lock the rows affected by the query from
// being accessed by any other transaction that also uses FOR
// UPDATE

// once the rows are locked (if forUpdate was true), we can query again using pop

func FindSessionByUserID(tx *storage.Connection, userId uuid.UUID) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindSessionsByFactorID(tx *storage.Connection, factorID uuid.UUID) ([]*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAllSessionsForUser finds all of the sessions for a user. If forUpdate is
// set, it will first lock on the user row which can be used to prevent issues
// with concurrency. If the lock is acquired, it will return a
// UserNotFoundError and the operation should be retried. If there are no
// sessions for the user, a nil result is returned without an error.
func FindAllSessionsForUser(tx *storage.Connection, userId uuid.UUID, forUpdate bool) ([]*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateFactorAssociatedSessions(tx *storage.Connection, userID, factorID uuid.UUID, aal string) error {
	_ = "STUB: not implemented"
	return nil
}

func InvalidateSessionsWithAALLessThan(tx *storage.Connection, userID uuid.UUID, level string) error {
	_ = "STUB: not implemented"
	return nil
}

// Logout deletes all sessions for a user.
func Logout(tx *storage.Connection, userId uuid.UUID) error { _ = "STUB: not implemented"; return nil }

// LogoutSession deletes the current session for a user
func LogoutSession(tx *storage.Connection, sessionId uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// LogoutAllExceptMe deletes all sessions for a user except the current one
func LogoutAllExceptMe(tx *storage.Connection, sessionId uuid.UUID, userID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// RevokeOAuthSessions deletes all sessions associated with a specific OAuth client for a user
func RevokeOAuthSessions(tx *storage.Connection, userID uuid.UUID, oauthClientID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Session) UpdateAALAndAssociatedFactor(tx *storage.Connection, aal AuthenticatorAssuranceLevel, factorID *uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Session) CalculateAALAndAMR(user *User) (aal AuthenticatorAssuranceLevel, amr []AMREntry, err error) {
	_ = "STUB: not implemented"
	return *new(AuthenticatorAssuranceLevel), nil, nil
}

// SSO users should only have one identity since they are excluded from account linking
// These checks act as a safeguard in the event future changes break this assumption.

// makes sure that the AMR claims are always ordered most-recent first

func (s *Session) GetAAL() string { _ = "STUB: not implemented"; return "" }

func (s *Session) IsAAL2() bool { _ = "STUB: not implemented"; return false }

// FindCurrentlyActiveRefreshToken returns the currently active refresh
// token in the session. This is the last created (ordered by the serial
// primary key) non-revoked refresh token for the session.
func (s *Session) FindCurrentlyActiveRefreshToken(tx *storage.Connection) (*RefreshToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetScopeList returns the scopes as a slice
func (s *Session) GetScopeList() []string { _ = "STUB: not implemented"; return nil }

// HasScope checks if the session has a specific scope
func (s *Session) HasScope(scope string) bool { _ = "STUB: not implemented"; return false }
