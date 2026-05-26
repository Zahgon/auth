package v0hooks

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase/auth/internal/mailer"
	"github.com/supabase/auth/internal/models"
)

type Name string

const (
	SendSMS              Name = "send-sms"
	SendEmail            Name = "send-email"
	CustomizeAccessToken Name = "customize-access-token"
	MFAVerification      Name = "mfa-verification"
	PasswordVerification Name = "password-verification"
	BeforeUserCreated    Name = "before-user-created"
	AfterUserCreated     Name = "after-user-created"
)

const (
	HookRejection = "reject"
)

const (
	DefaultMFAHookRejectionMessage      = "Further MFA verification attempts will be rejected."
	DefaultPasswordHookRejectionMessage = "Further password verification attempts will be rejected."
)

type Metadata struct {
	UUID uuid.UUID `json:"uuid"`
	Time time.Time `json:"time"`

	// Hook name
	Name Name `json:"name,omitempty"`

	// IP Address of the request, if present
	IPAddress string `json:"ip_address,omitempty"`
}

func NewMetadata(r *http.Request, name Name) *Metadata { _ = "STUB: not implemented"; return nil }

type BeforeUserCreatedInput struct {
	Metadata *Metadata    `json:"metadata"`
	User     *models.User `json:"user"`
}

func NewBeforeUserCreatedInput(
	r *http.Request,
	user *models.User,
) *BeforeUserCreatedInput {
	_ = "STUB: not implemented"
	return nil
}

type BeforeUserCreatedOutput struct {
}

type AfterUserCreatedInput struct {
	Metadata *Metadata    `json:"metadata"`
	User     *models.User `json:"user"`
}

func NewAfterUserCreatedInput(
	r *http.Request,
	user *models.User,
) *AfterUserCreatedInput {
	_ = "STUB: not implemented"
	return nil
}

type AfterUserCreatedOutput struct{}

// TODO(joel): Move this to phone package
type SMS struct {
	OTP     string `json:"otp,omitempty"`
	SMSType string `json:"sms_type,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

// AccessTokenClaims is a struct thats used for JWT claims
type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Email                         string                 `json:"email"`
	Phone                         string                 `json:"phone"`
	AppMetaData                   map[string]interface{} `json:"app_metadata"`
	UserMetaData                  map[string]interface{} `json:"user_metadata"`
	Role                          string                 `json:"role"`
	AuthenticatorAssuranceLevel   string                 `json:"aal,omitempty"`
	AuthenticationMethodReference []models.AMREntry      `json:"amr,omitempty"`
	SessionId                     string                 `json:"session_id,omitempty"`
	IsAnonymous                   bool                   `json:"is_anonymous"`
	ClientID                      string                 `json:"client_id,omitempty"`
	Scope                         string                 `json:"scope,omitempty"`
}

type MFAVerificationAttemptInput struct {
	Metadata   *Metadata `json:"metadata"`
	UserID     uuid.UUID `json:"user_id"`
	FactorID   uuid.UUID `json:"factor_id"`
	FactorType string    `json:"factor_type"`
	Valid      bool      `json:"valid"`
}

func NewMFAVerificationAttemptInput(
	r *http.Request,
	userID uuid.UUID,
	factorID uuid.UUID,
	factorType string,
	valid bool,
) *MFAVerificationAttemptInput {
	_ = "STUB: not implemented"
	return nil
}

type MFAVerificationAttemptOutput struct {
	Decision string `json:"decision"`
	Message  string `json:"message"`
}

type PasswordVerificationAttemptInput struct {
	Metadata *Metadata `json:"metadata"`
	UserID   uuid.UUID `json:"user_id"`
	Valid    bool      `json:"valid"`
}

func NewPasswordVerificationAttemptInput(
	r *http.Request,
	userID uuid.UUID,
	valid bool,
) *PasswordVerificationAttemptInput {
	_ = "STUB: not implemented"
	return nil
}

type PasswordVerificationAttemptOutput struct {
	Decision         string `json:"decision"`
	Message          string `json:"message"`
	ShouldLogoutUser bool   `json:"should_logout_user"`
}

type CustomAccessTokenInput struct {
	Metadata             *Metadata          `json:"metadata"`
	UserID               uuid.UUID          `json:"user_id"`
	Claims               *AccessTokenClaims `json:"claims"`
	AuthenticationMethod string             `json:"authentication_method"`
}

func NewCustomAccessTokenInput(
	r *http.Request,
	userID uuid.UUID,
	claims *AccessTokenClaims,
	authenticationMethod string,
) *CustomAccessTokenInput {
	_ = "STUB: not implemented"
	return nil
}

type CustomAccessTokenOutput struct {
	Claims map[string]any `json:"claims"`
}

func (o *CustomAccessTokenOutput) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// First check if the claims field is missing

// This check allows us to skip an additional unmarshal for valid inputs

// The Claims field is not a map[string]any so we unmarshal again just
// to get the correct error type.

type SendSMSInput struct {
	Metadata *Metadata    `json:"metadata"`
	User     *models.User `json:"user,omitempty"`
	SMS      SMS          `json:"sms,omitempty"`
}

func NewSendSMSInput(
	r *http.Request,
	user *models.User,
	sms SMS,
) *SendSMSInput {
	_ = "STUB: not implemented"
	return nil
}

type SendSMSOutput struct {
}

type SendEmailInput struct {
	Metadata  *Metadata        `json:"metadata"`
	User      *models.User     `json:"user"`
	EmailData mailer.EmailData `json:"email_data"`
}

func NewSendEmailInput(
	r *http.Request,
	user *models.User,
	emailData mailer.EmailData,
) *SendEmailInput {
	_ = "STUB: not implemented"
	return nil
}

type SendEmailOutput struct {
}
