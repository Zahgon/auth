package templatemailer

import (
	"net/http"
	"net/url"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/models"
)

const (
	InviteTemplate           = "invite"
	ConfirmationTemplate     = "confirmation"
	RecoveryTemplate         = "recovery"
	EmailChangeTemplate      = "email_change"
	MagicLinkTemplate        = "magic_link"
	ReauthenticationTemplate = "reauthentication"

	// Account Changes Notifications
	PasswordChangedNotificationTemplate     = "password_changed_notification"
	EmailChangedNotificationTemplate        = "email_changed_notification"
	PhoneChangedNotificationTemplate        = "phone_changed_notification"
	IdentityLinkedNotificationTemplate      = "identity_linked_notification"
	IdentityUnlinkedNotificationTemplate    = "identity_unlinked_notification"
	MFAFactorEnrolledNotificationTemplate   = "mfa_factor_enrolled_notification"
	MFAFactorUnenrolledNotificationTemplate = "mfa_factor_unenrolled_notification"
)

const defaultInviteMail = `<h2>You've been invited</h2>

<p>You've been invited to create an account. Follow the link below to accept.</p>
<p><a href="{{ .ConfirmationURL }}">Accept invitation</a></p>`

const defaultConfirmationMail = `<h2>Confirm your email address</h2>

<p>Follow the link below to confirm this email address and finish signing up.</p>
<p><a href="{{ .ConfirmationURL }}">Confirm email address</a></p>
`

const defaultRecoveryMail = `<h2>Reset your password</h2>

<p>We received a request to reset your password. Follow the link below to choose a new one.</p>
<p><a href="{{ .ConfirmationURL }}">Reset password</a></p>
<p>If you didn't request this, you can safely ignore this email.</p>`

const defaultMagicLinkMail = `<h2>Your sign-in link</h2>

<p>Follow the link below to sign in. This link expires shortly and can only be used once.</p>
<p><a href="{{ .ConfirmationURL }}">Sign in</a></p>`

const defaultEmailChangeMail = `<h2>Confirm your new email address</h2>

<p>Follow the link below to confirm {{ .NewEmail }} as your new email address.</p>
<p><a href="{{ .ConfirmationURL }}">Confirm new email address</a></p>
<p>If you didn't request this change, you can safely ignore this email.</p>`

const defaultReauthenticateMail = `<h2>Your verification code</h2>

<p>Use the code below to verify your identity. It expires shortly.</p>
<p>{{ .Token }}</p>`

// Account Changes Notifications

// #nosec G101 -- No hardcoded credentials.
const defaultPasswordChangedNotificationMail = `<h2>Your password was changed</h2>

<p>The password for your account was recently changed.</p>
<p>If you didn't make this change, reset your password and contact support immediately.</p>
`
const defaultEmailChangedNotificationMail = `<h2>Your email address was changed</h2>

<p>The email address for your account was changed from {{ .OldEmail }} to {{ .Email }}.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

const defaultPhoneChangedNotificationMail = `<h2>Your phone number was changed</h2>

<p>The phone number for your account was changed from {{ .OldPhone }} to {{ .Phone }}.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

const defaultIdentityLinkedNotificationMail = `<h2>A new sign-in method was linked</h2>

<p>Your {{ .Provider }} account was linked as a new sign-in method for {{ .Email }}.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

const defaultIdentityUnlinkedNotificationMail = `<h2>A sign-in method was removed</h2>

<p>Your {{ .Provider }} account was removed as a sign-in method for {{ .Email }}.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

const defaultMFAFactorEnrolledNotificationMail = `<h2>A new verification method was added</h2>

<p>Sign-in verification method {{ .FactorType }} was added to your account.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

const defaultMFAFactorUnenrolledNotificationMail = `<h2>A verification method was removed</h2>

<p>Sign-in verification method {{ .FactorType }} was removed from your account.</p>
<p>If you didn't make this change, contact support immediately.</p>
`

var (
	templateTypes = []string{
		InviteTemplate,
		ConfirmationTemplate,
		RecoveryTemplate,
		EmailChangeTemplate,
		MagicLinkTemplate,
		ReauthenticationTemplate,

		// Account Changes Notifications
		PasswordChangedNotificationTemplate,
		EmailChangedNotificationTemplate,
		PhoneChangedNotificationTemplate,
		IdentityLinkedNotificationTemplate,
		IdentityUnlinkedNotificationTemplate,
		MFAFactorEnrolledNotificationTemplate,
		MFAFactorUnenrolledNotificationTemplate,
	}
	defaultTemplateSubjects = &conf.EmailContentConfiguration{
		Invite:           "You've been invited",
		Confirmation:     "Confirm your email address",
		Recovery:         "Reset your password",
		MagicLink:        "Your sign-in link",
		EmailChange:      "Confirm your new email address",
		Reauthentication: "{{ .Token }} is your verification code",

		// Account Changes Notifications
		PasswordChangedNotification:     "Your password was changed",
		EmailChangedNotification:        "Your email address was changed",
		PhoneChangedNotification:        "Your phone number was changed",
		IdentityLinkedNotification:      "A new sign-in method was linked to your account",
		IdentityUnlinkedNotification:    "A sign-in method was removed from your account",
		MFAFactorEnrolledNotification:   "A new verification method was added to your account",
		MFAFactorUnenrolledNotification: "A verification method was removed from your account",
	}
	defaultTemplateBodies = &conf.EmailContentConfiguration{
		Invite:           defaultInviteMail,
		Confirmation:     defaultConfirmationMail,
		Recovery:         defaultRecoveryMail,
		MagicLink:        defaultMagicLinkMail,
		EmailChange:      defaultEmailChangeMail,
		Reauthentication: defaultReauthenticateMail,

		// Account Changes Notifications
		PasswordChangedNotification:     defaultPasswordChangedNotificationMail,
		EmailChangedNotification:        defaultEmailChangedNotificationMail,
		PhoneChangedNotification:        defaultPhoneChangedNotificationMail,
		IdentityLinkedNotification:      defaultIdentityLinkedNotificationMail,
		IdentityUnlinkedNotification:    defaultIdentityUnlinkedNotificationMail,
		MFAFactorEnrolledNotification:   defaultMFAFactorEnrolledNotificationMail,
		MFAFactorUnenrolledNotification: defaultMFAFactorUnenrolledNotificationMail,
	}
)

func (m *Mailer) Headers(cfg *conf.GlobalConfiguration, messageType string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// TODO: in the future, use a templating engine to add more contextual data available to headers

// InviteMail sends a invite mail to a new user
func (m *Mailer) InviteMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfirmationMail sends a signup confirmation mail to a new user
func (m *Mailer) ConfirmationMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// ReauthenticateMail sends a reauthentication mail to an authenticated user
func (m *Mailer) ReauthenticateMail(r *http.Request, user *models.User, otp string) error {
	_ = "STUB: not implemented"
	return nil
}

// EmailChangeMail sends an email change confirmation mail to a user
func (m *Mailer) EmailChangeMail(r *http.Request, user *models.User, otpNew, otpCurrent, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoveryMail sends a password recovery mail
func (m *Mailer) RecoveryMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// MagicLinkMail sends a login link mail
func (m *Mailer) MagicLinkMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// GetEmailActionLink returns a magiclink, recovery or invite link based on the actionType passed.
func (m *Mailer) GetEmailActionLink(user *models.User, actionType, referrerURL string, externalURL *url.URL) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *Mailer) PasswordChangedNotificationMail(r *http.Request, user *models.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mailer) EmailChangedNotificationMail(r *http.Request, user *models.User, oldEmail string) error {
	_ = "STUB: not implemented"
	return nil
}

// the new email address that has been set on the account
// the old email address that was on the account before the change

func (m *Mailer) PhoneChangedNotificationMail(r *http.Request, user *models.User, oldPhone string) error {
	_ = "STUB: not implemented"
	return nil
}

// the new phone number that has been set on the account
// the old phone number that was on the account before the change

func (m *Mailer) IdentityLinkedNotificationMail(r *http.Request, user *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

// the provider of the newly linked identity

func (m *Mailer) IdentityUnlinkedNotificationMail(r *http.Request, user *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

// the provider of the unlinked identity

func (m *Mailer) MFAFactorEnrolledNotificationMail(r *http.Request, user *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mailer) MFAFactorUnenrolledNotificationMail(r *http.Request, user *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

type emailParams struct {
	Token      string
	Type       string
	RedirectTo string
}

func getPath(filepath string, params *emailParams) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeRedirectURL(referrerURL string) string { _ = "STUB: not implemented"; return "" }

// if the string contains &, = or # it has not been URL
// encoded by the caller, which means it should be URL
// encoded by us otherwise, it should be taken as-is
