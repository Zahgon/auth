package api

import (
	"net/http"
	"regexp"
	"time"

	"github.com/supabase/auth/internal/observability"

	"github.com/pkg/errors"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

var (
	EmailRateLimitExceeded error = errors.New("email rate limit exceeded")
	emailSendCounter             = observability.ObtainMetricCounter("global_auth_email_send_operations_total", "Number of email send operations")
	emailErrorsCounter           = observability.ObtainMetricCounter("global_auth_email_send_errors_total", "Number of email send errors")
)

type GenerateLinkParams struct {
	Type       string                 `json:"type"`
	Email      string                 `json:"email"`
	NewEmail   string                 `json:"new_email"`
	Password   string                 `json:"password"`
	Data       map[string]interface{} `json:"data"`
	RedirectTo string                 `json:"redirect_to"`
}

type GenerateLinkResponse struct {
	models.User
	ActionLink       string `json:"action_link"`
	EmailOtp         string `json:"email_otp"`
	HashedToken      string `json:"hashed_token"`
	VerificationType string `json:"verification_type"`
	RedirectTo       string `json:"redirect_to"`
}

func (a *API) adminGenerateLink(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// password generation must always succeed

/* <- isSSOUser */

/* <- isSSOUser */

// you should never use SignupParams with
// password here to generate a new user, use
// signupUser which is a model generated from
// SignupParams above

func (a *API) sendConfirmation(r *http.Request, tx *storage.Connection, u *models.User, flowType models.FlowType) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendInvite(r *http.Request, tx *storage.Connection, u *models.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendPasswordRecovery(r *http.Request, tx *storage.Connection, u *models.User, flowType models.FlowType) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendReauthenticationOtp(r *http.Request, tx *storage.Connection, u *models.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendMagicLink(r *http.Request, tx *storage.Connection, u *models.User, flowType models.FlowType) error {
	_ = "STUB: not implemented"
	return nil
}

// since Magic Link is just a recovery with a different template and behaviour
// around new users we will reuse the recovery db timer to prevent potential abuse

// sendEmailChange sends out an email change token to the new email.
func (a *API) sendEmailChange(r *http.Request, tx *storage.Connection, u *models.User, email string, flowType models.FlowType) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendPasswordChangedNotification(r *http.Request, tx *storage.Connection, u *models.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendEmailChangedNotification(r *http.Request, tx *storage.Connection, u *models.User, oldEmail string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendPhoneChangedNotification(r *http.Request, tx *storage.Connection, u *models.User, oldPhone string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendIdentityLinkedNotification(r *http.Request, tx *storage.Connection, u *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendIdentityUnlinkedNotification(r *http.Request, tx *storage.Connection, u *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendMFAFactorEnrolledNotification(r *http.Request, tx *storage.Connection, u *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) sendMFAFactorUnenrolledNotification(r *http.Request, tx *storage.Connection, u *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) validateEmail(email string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateSentWithinFrequencyLimit(sentAt *time.Time, frequency time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

var emailLabelPattern = regexp.MustCompile("[+][^@]+@")

func (a *API) checkEmailAddressAuthorization(email string) bool {
	_ = "STUB: not implemented"
	return false
}

// allow labelled emails when authorization rules are in place

type sendEmailParams struct {
	emailActionType     string
	otp                 string
	otpNew              string
	tokenHashWithPrefix string
	oldEmail            string
	oldPhone            string
	provider            string
	factorType          string
}

func (a *API) sendEmail(r *http.Request, tx *storage.Connection, u *models.User, params sendEmailParams) error {
	_ = "STUB: not implemented"
	return nil
}

// first check that the user can update their address to the
// new one in u.EmailChange

// if secure email change is enabled, check that the user
// account (which could have been created before the authorized
// address authorization restriction was enabled) can even
// receive the confirmation message to the existing address

// if the number of events is set to zero, we immediately apply rate limits.

// TODO(km): Deprecate this behaviour - rate limits should still be applied to autoconfirm

// apply rate limiting before the email is sent out

// When secure email change is disabled, we place the token for the new email on emailData.Token

// BUG(cstockton): This introduced a bug which mismatched the token
// and hash fields, such that:
//
// 	EmailData.TokenHashNew = Hash(CurEmail, EmailData.Token)
// 	EmailData.TokenHash    = Hash(NewEmail, EmailData.TokenNew)
//
// Specifically with email changes we should look to fix this
// behavior in a BC way to maintain that:
//
//   Token      Always contains the Token for user.email
//   TokenHash  Always contains the Hash for user.email
//
//   Token      Always contains the Token for user.email_new
//   TokenHash  Always contains the Hash for user.email_new
//

// BUG(cstockton): This matches the current behavior but is not
// intuitive and should be changed in a future release. See the
// comment above for more details.

// Augment the email data for the email send hook with notification-specific fields

// Increment email send operations here, since this metric is meant to count number of mail
// send operations rather than simply number of attempts to send mail
