package models

import (
	"context"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/crypto"
	"github.com/supabase/auth/internal/storage"
)

// User respresents a registered user with email/password authentication
type User struct {
	ID uuid.UUID `json:"id" db:"id"`

	Aud       string             `json:"aud" db:"aud"`
	Role      string             `json:"role" db:"role"`
	Email     storage.NullString `json:"email" db:"email"`
	IsSSOUser bool               `json:"-" db:"is_sso_user"`

	EncryptedPassword *string    `json:"-" db:"encrypted_password"`
	EmailConfirmedAt  *time.Time `json:"email_confirmed_at,omitempty" db:"email_confirmed_at"`
	InvitedAt         *time.Time `json:"invited_at,omitempty" db:"invited_at"`

	Phone            storage.NullString `json:"phone" db:"phone"`
	PhoneConfirmedAt *time.Time         `json:"phone_confirmed_at,omitempty" db:"phone_confirmed_at"`

	ConfirmationToken  string     `json:"-" db:"confirmation_token"`
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty" db:"confirmation_sent_at"`

	// For backward compatibility only. Use EmailConfirmedAt or PhoneConfirmedAt instead.
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty" db:"confirmed_at" rw:"r"`

	RecoveryToken  string     `json:"-" db:"recovery_token"`
	RecoverySentAt *time.Time `json:"recovery_sent_at,omitempty" db:"recovery_sent_at"`

	EmailChangeTokenCurrent  string     `json:"-" db:"email_change_token_current"`
	EmailChangeTokenNew      string     `json:"-" db:"email_change_token_new"`
	EmailChange              string     `json:"new_email,omitempty" db:"email_change"`
	EmailChangeSentAt        *time.Time `json:"email_change_sent_at,omitempty" db:"email_change_sent_at"`
	EmailChangeConfirmStatus int        `json:"-" db:"email_change_confirm_status"`

	PhoneChangeToken  string     `json:"-" db:"phone_change_token"`
	PhoneChange       string     `json:"new_phone,omitempty" db:"phone_change"`
	PhoneChangeSentAt *time.Time `json:"phone_change_sent_at,omitempty" db:"phone_change_sent_at"`

	ReauthenticationToken  string     `json:"-" db:"reauthentication_token"`
	ReauthenticationSentAt *time.Time `json:"reauthentication_sent_at,omitempty" db:"reauthentication_sent_at"`

	LastSignInAt *time.Time `json:"last_sign_in_at,omitempty" db:"last_sign_in_at"`

	AppMetaData  JSONMap `json:"app_metadata" db:"raw_app_meta_data"`
	UserMetaData JSONMap `json:"user_metadata" db:"raw_user_meta_data"`

	Factors    []Factor   `json:"factors,omitempty" has_many:"factors"`
	Identities []Identity `json:"identities" has_many:"identities"`

	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	BannedUntil *time.Time `json:"banned_until,omitempty" db:"banned_until"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IsAnonymous bool       `json:"is_anonymous" db:"is_anonymous"`

	DONTUSEINSTANCEID uuid.UUID `json:"-" db:"instance_id"`
}

func NewUserWithPasswordHash(phone, email, passwordHash, aud string, userData map[string]interface{}) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verify that the hash is a bcrypt hash

// NewUser initializes a new user from an email, password and user data.
func NewUser(phone, email, password, aud string, userData map[string]interface{}) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TableName overrides the table name used by pop
func (User) TableName() string { _ = "STUB: not implemented"; return "" }

func (u *User) HasPassword() bool { _ = "STUB: not implemented"; return false }

// BeforeSave is invoked before the user is saved to the database
func (u *User) BeforeSave(tx *pop.Connection) error { _ = "STUB: not implemented"; return nil }

// IsConfirmed checks if a user has already been
// registered and confirmed.
func (u *User) IsConfirmed() bool { _ = "STUB: not implemented"; return false }

// HasBeenInvited checks if user has been invited
func (u *User) HasBeenInvited() bool { _ = "STUB: not implemented"; return false }

// IsPhoneConfirmed checks if a user's phone has already been
// registered and confirmed.
func (u *User) IsPhoneConfirmed() bool { _ = "STUB: not implemented"; return false }

// SetRole sets the users Role to roleName
func (u *User) SetRole(tx *storage.Connection, roleName string) error {
	_ = "STUB: not implemented"
	return nil
}

// HasRole returns true when the users role is set to roleName
func (u *User) HasRole(roleName string) bool { _ = "STUB: not implemented"; return false }

// GetEmail returns the user's email as a string
func (u *User) GetEmail() string { _ = "STUB: not implemented"; return "" }

// GetPhone returns the user's phone number as a string
func (u *User) GetPhone() string { _ = "STUB: not implemented"; return "" }

// UpdateUserMetaData sets all user data from a map of updates,
// ensuring that it doesn't override attributes that are not
// in the provided map.
func (u *User) UpdateUserMetaData(tx *storage.Connection, updates map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAppMetaData updates all app data from a map of updates
func (u *User) UpdateAppMetaData(tx *storage.Connection, updates map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAppMetaDataProviders updates the provider field in AppMetaData column
func (u *User) UpdateAppMetaDataProviders(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateUserEmail updates the user's email to one of the identity's email
// if the current email used doesn't match any of the identities email
func (u *User) UpdateUserEmailFromIdentities(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// there's an existing identity that uses the same email
// so the user's email can be kept

// the identity's email is not used by another user
// so we can set it as the primary identity

// default to the first identity's email

// SetEmail sets the user's email
func (u *User) SetEmail(tx *storage.Connection, email string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetPhone sets the user's phone
func (u *User) SetPhone(tx *storage.Connection, phone string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *User) SetPassword(ctx context.Context, password string, encrypt bool, encryptionKeyID, encryptionKey string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdatePassword updates the user's password. Use SetPassword outside of a transaction first!
func (u *User) UpdatePassword(tx *storage.Connection, sessionID *uuid.UUID) error {
	_ = "STUB: not implemented"
	// These need to be reset because password change may mean the user no longer trusts the actions performed by the previous password.
	return nil
}

// log out user from all sessions to ensure reauthentication after password change

// log out user from all other sessions to ensure reauthentication after password change

// Authenticate a user from a password
func (u *User) Authenticate(ctx context.Context, tx *storage.Connection, password string, decryptionKeys map[string]string, encrypt bool, encryptionKeyID string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// check if cost exceeds default cost or is too low

// don't bother with encrypting the password in Authenticate
// since it's handled separately

// ConfirmReauthentication resets the reauthentication token
func (u *User) ConfirmReauthentication(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Confirm resets the confimation token and sets the confirm timestamp
func (u *User) Confirm(tx *storage.Connection) error { _ = "STUB: not implemented"; return nil }

// ConfirmPhone resets the confimation token and sets the confirm timestamp
func (u *User) ConfirmPhone(tx *storage.Connection) error { _ = "STUB: not implemented"; return nil }

// UpdateLastSignInAt update field last_sign_in_at for user according to specified field
func (u *User) UpdateLastSignInAt(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfirmEmailChange confirm the change of email for a user
func (u *User) ConfirmEmailChange(tx *storage.Connection, status int) error {
	_ = "STUB: not implemented"
	return nil
}

// no email identity, not an error

// ConfirmPhoneChange confirms the change of phone for a user
func (u *User) ConfirmPhoneChange(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// no phone identity, not an error

// Recover resets the recovery token
func (u *User) Recover(tx *storage.Connection) error { _ = "STUB: not implemented"; return nil }

// HighestPossibleAAL returns the AAL level that this user can obtain. Derived
// from the number of verified MFA factors associated with the user object.
func (u *User) HighestPossibleAAL() AuthenticatorAssuranceLevel {
	_ = "STUB: not implemented"
	return *new(AuthenticatorAssuranceLevel)
}

// CountOtherUsers counts how many other users exist besides the one provided
func CountOtherUsers(tx *storage.Connection, id uuid.UUID) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func findUser(tx *storage.Connection, query string, args ...interface{}) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByEmailAndAudience finds a user with the matching email and audience.
func FindUserByEmailAndAudience(tx *storage.Connection, email, aud string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByPhoneAndAudience finds a user with the matching email and audience.
func FindUserByPhoneAndAudience(tx *storage.Connection, phone, aud string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserByID finds a user matching the provided ID.
func FindUserByID(tx *storage.Connection, id uuid.UUID) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindUserWithRefreshToken finds a user from the provided refresh token. If
// forUpdate is set to true, then the SELECT statement used by the query has
// the form SELECT ... FOR UPDATE SKIP LOCKED. This means that a FOR UPDATE
// lock will only be acquired if there's no other lock. In case there is a
// lock, a IsNotFound(err) error will be returned.
//
// Second value returned is either *models.RefreshToken or *crypto.RefreshToken.
func FindUserWithRefreshToken(tx *storage.Connection, dbEncryption conf.DatabaseEncryptionConfiguration, token string, forUpdate bool) (*User, any, *Session, error) {
	_ = "STUB: not implemented"
	return nil,

		// not a valid refresh token so don't bother looking it up in the database
		*new(any), nil, nil
}

func findUserWithRefreshToken(tx *storage.Connection, dbEncryption conf.DatabaseEncryptionConfiguration, token string, forUpdate bool) (*User, *crypto.RefreshToken, *Session, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// refresh token is not valid

// if the session is not set up to support these encoded refresh tokens it's as if it doesn't exist
// meaning someone is hand-crafting tokens for uuids that exist

// refresh token signature is not valid for this session

func findUserWithLegacyRefreshToken(tx *storage.Connection, token string, forUpdate bool) (*User, any, *Session, error) {
	_ = "STUB: not implemented"
	return nil, *new(any), nil, nil
}

// pop does not provide us with a way to execute FOR UPDATE
// queries which lock the rows affected by the query from
// being accessed by any other transaction that also uses FOR
// UPDATE

// once the rows are locked (if forUpdate was true), we can query again using pop

// otherwise, there's no session for this refresh token

// FindUsersInAudience finds users with the matching audience.
func FindUsersInAudience(tx *storage.Connection, aud string, pageParams *Pagination, sortParams *SortParams, filter string) ([]*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we must specify the collation in order to get case insensitive search for the JSON column

// #nosec G115
// #nosec G115

// IsDuplicatedEmail returns whether a user exists with a matching email and
// audience importantly in the *default* identity linking domain (meaning SSO
// accounts and similar are not considered).
// If a currentUser is provided, we will need to filter out any identities that belong to the current user.
func IsDuplicatedEmail(tx *storage.Connection, email, aud string, currentUser *User, ownDomainProviders []string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// out of an abundance of caution, if nothing was found via the
// identities table we also do a final check on the users table

// IsDuplicatedPhone checks if the phone number already exists in the users table
func IsDuplicatedPhone(tx *storage.Connection, phone, aud string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Ban a user for a given duration.
func (u *User) Ban(tx *storage.Connection, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// IsBanned checks if a user is banned or not
func (u *User) IsBanned() bool { _ = "STUB: not implemented"; return false }

func (u *User) HasMFAEnabled() bool { _ = "STUB: not implemented"; return false }

func (u *User) UpdateBannedUntil(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveUnconfirmedIdentities removes potentially malicious unconfirmed identities from a user (if any)
func (u *User) RemoveUnconfirmedIdentities(tx *storage.Connection, identity *Identity) error {
	_ = "STUB: not implemented"
	return nil
}

// user is unconfirmed so the password should be reset

// user is unconfirmed so existing user_metadata should be overwritten
// to use the current identity metadata

// finally, remove all identities except the current identity being authenticated

// user is unconfirmed so none of the providers associated to it are verified yet
// only the current provider should be kept

// SoftDeleteUser performs a soft deletion on the user by obfuscating and clearing certain fields
func (u *User) SoftDeleteUser(tx *storage.Connection) error { _ = "STUB: not implemented"; return nil }

// set deleted_at time

// set raw_user_meta_data to {}

// set raw_app_meta_data to {}

// SoftDeleteUserIdentities performs a soft deletion on all identities associated to a user
func (u *User) SoftDeleteUserIdentities(tx *storage.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// set identity_data to {}

// updating the identity.ID has to happen last since the primary key is on (provider, id)
// we use RawQuery here instead of UpdateOnly because UpdateOnly relies on the primary key of Identity

func (u *User) FindOwnedFactorByID(tx *storage.Connection, factorID uuid.UUID) (*Factor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (user *User) WebAuthnID() []byte { _ = "STUB: not implemented"; return nil }

func (user *User) WebAuthnName() string { _ = "STUB: not implemented"; return "" }

func (user *User) WebAuthnDisplayName() string { _ = "STUB: not implemented"; return "" }

func (user *User) WebAuthnCredentials() []webauthn.Credential {
	_ = "STUB: not implemented"
	return nil
}

func obfuscateValue(id uuid.UUID, value string) string { _ = "STUB: not implemented"; return "" }

func obfuscateEmail(u *User, email string) string { _ = "STUB: not implemented"; return "" }

func obfuscatePhone(u *User, phone string) string {
	_ = "STUB: not implemented"
	// Field converted from VARCHAR(15) to text
	return ""
}

func obfuscateIdentityProviderId(identity *Identity) string { _ = "STUB: not implemented"; return "" }

// FindUserByPhoneChangeAndAudience finds a user with the matching phone change and audience.
func FindUserByPhoneChangeAndAudience(tx *storage.Connection, phone, aud string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
