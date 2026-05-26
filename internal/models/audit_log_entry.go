package models

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

type AuditAction string
type auditLogType string

const (
	LoginAction                     AuditAction = "login"
	LogoutAction                    AuditAction = "logout"
	InviteAcceptedAction            AuditAction = "invite_accepted"
	UserSignedUpAction              AuditAction = "user_signedup"
	UserInvitedAction               AuditAction = "user_invited"
	UserDeletedAction               AuditAction = "user_deleted"
	UserModifiedAction              AuditAction = "user_modified"
	UserRecoveryRequestedAction     AuditAction = "user_recovery_requested"
	UserReauthenticateAction        AuditAction = "user_reauthenticate_requested"
	UserConfirmationRequestedAction AuditAction = "user_confirmation_requested"
	UserRepeatedSignUpAction        AuditAction = "user_repeated_signup"
	UserUpdatePasswordAction        AuditAction = "user_updated_password"
	TokenRevokedAction              AuditAction = "token_revoked"
	TokenRefreshedAction            AuditAction = "token_refreshed"
	GenerateRecoveryCodesAction     AuditAction = "generate_recovery_codes"
	EnrollFactorAction              AuditAction = "factor_in_progress"
	UnenrollFactorAction            AuditAction = "factor_unenrolled"
	CreateChallengeAction           AuditAction = "challenge_created"
	VerifyFactorAction              AuditAction = "verification_attempted"
	DeleteFactorAction              AuditAction = "factor_deleted"
	DeleteRecoveryCodesAction       AuditAction = "recovery_codes_deleted"
	UpdateFactorAction              AuditAction = "factor_updated"
	MFACodeLoginAction              AuditAction = "mfa_code_login"
	IdentityUnlinkAction            AuditAction = "identity_unlinked"
	PasskeyCreatedAction            AuditAction = "passkey_created"
	PasskeyUpdatedAction            AuditAction = "passkey_updated"
	PasskeyDeletedAction            AuditAction = "passkey_deleted"

	account       auditLogType = "account"
	team          auditLogType = "team"
	token         auditLogType = "token"
	user          auditLogType = "user"
	factor        auditLogType = "factor"
	recoveryCodes auditLogType = "recovery_codes"
)

var ActionLogTypeMap = map[AuditAction]auditLogType{
	LoginAction:                     account,
	LogoutAction:                    account,
	InviteAcceptedAction:            account,
	UserSignedUpAction:              team,
	UserInvitedAction:               team,
	UserDeletedAction:               team,
	TokenRevokedAction:              token,
	TokenRefreshedAction:            token,
	UserModifiedAction:              user,
	UserRecoveryRequestedAction:     user,
	UserConfirmationRequestedAction: user,
	UserRepeatedSignUpAction:        user,
	UserUpdatePasswordAction:        user,
	GenerateRecoveryCodesAction:     user,
	EnrollFactorAction:              factor,
	UnenrollFactorAction:            factor,
	CreateChallengeAction:           factor,
	VerifyFactorAction:              factor,
	DeleteFactorAction:              factor,
	UpdateFactorAction:              factor,
	MFACodeLoginAction:              factor,
	DeleteRecoveryCodesAction:       recoveryCodes,
	PasskeyCreatedAction:            user,
	PasskeyUpdatedAction:            user,
	PasskeyDeletedAction:            user,
}

// AuditLogEntry is the database model for audit log entries.
type AuditLogEntry struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Payload   JSONMap   `json:"payload" db:"payload"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IPAddress string    `json:"ip_address" db:"ip_address"`

	DONTUSEINSTANCEID uuid.UUID `json:"-" db:"instance_id"`
}

func (AuditLogEntry) TableName() string { _ = "STUB: not implemented"; return "" }

func NewAuditLogEntry(config conf.AuditLogConfiguration, r *http.Request, tx *storage.Connection, actor *User, action AuditAction, ipAddress string, traits map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// AUDIT LOGGING FIX: Log each audit event immediately as a separate log entry
//
// BUG: The observability.LogEntrySetFields() above adds to request context, causing
// multiple audit events in the same request to overwrite each other. For example,
// refresh token requests call NewAuditLogEntry() twice (token_refreshed, then
// token_revoked) but only the last event (token_revoked) was logged.
//
// SOLUTION: Create immediate separate log entries with "auth_audit_event" key.
// This ensures all audit events are captured without overwriting.
//
// TRANSITION: We keep the existing "auth_event" for backward compatibility during
// the transition period. This fix may impact metrics that count audit events,
// as previously missing events (like token_refreshed) will now appear in logs.
// Eventually, we should remove the observability.LogEntrySetFields() call above
// once new logging is proven stable.

func FindAuditLogEntries(tx *storage.Connection, filterColumns []string, filterValue string, pageParams *Pagination) ([]*AuditLogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G115
// #nosec G115
