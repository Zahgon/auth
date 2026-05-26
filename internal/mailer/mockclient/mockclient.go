package mockclient

import (
	"net/http"
	"net/url"

	"github.com/supabase/auth/internal/models"
)

// MockMailer implements the mailer.Mailer interface for testing
type MockMailer struct {
	InviteMailCalls         []InviteMailCall
	ConfirmationMailCalls   []ConfirmationMailCall
	RecoveryMailCalls       []RecoveryMailCall
	MagicLinkMailCalls      []MagicLinkMailCall
	EmailChangeMailCalls    []EmailChangeMailCall
	ReauthenticateMailCalls []ReauthenticateMailCall
	GetEmailActionLinkCalls []GetEmailActionLinkCall

	PasswordChangedMailCalls     []PasswordChangedMailCall
	EmailChangedMailCalls        []EmailChangedMailCall
	PhoneChangedMailCalls        []PhoneChangedMailCall
	IdentityLinkedMailCalls      []IdentityLinkedMailCall
	IdentityUnlinkedMailCalls    []IdentityUnlinkedMailCall
	MFAFactorEnrolledMailCalls   []MFAFactorEnrolledMailCall
	MFAFactorUnenrolledMailCalls []MFAFactorUnenrolledMailCall
}

type InviteMailCall struct {
	User        *models.User
	OTP         string
	ReferrerURL string
	ExternalURL *url.URL
}

type ConfirmationMailCall struct {
	User        *models.User
	OTP         string
	ReferrerURL string
	ExternalURL *url.URL
}

type RecoveryMailCall struct {
	User        *models.User
	OTP         string
	ReferrerURL string
	ExternalURL *url.URL
}

type MagicLinkMailCall struct {
	User        *models.User
	OTP         string
	ReferrerURL string
	ExternalURL *url.URL
}

type EmailChangeMailCall struct {
	User        *models.User
	OTPNew      string
	OTPCurrent  string
	ReferrerURL string
	ExternalURL *url.URL
}

type ReauthenticateMailCall struct {
	User *models.User
	OTP  string
}

type GetEmailActionLinkCall struct {
	User        *models.User
	ActionType  string
	ReferrerURL string
	ExternalURL *url.URL
	Result      string
	Error       error
}

type PasswordChangedMailCall struct {
	User *models.User
}

type EmailChangedMailCall struct {
	User     *models.User
	OldEmail string
}

type PhoneChangedMailCall struct {
	User     *models.User
	OldPhone string
}

type IdentityLinkedMailCall struct {
	User     *models.User
	Provider string
}

type IdentityUnlinkedMailCall struct {
	User     *models.User
	Provider string
}

type MFAFactorEnrolledMailCall struct {
	User       *models.User
	FactorType string
}

type MFAFactorUnenrolledMailCall struct {
	User       *models.User
	FactorType string
}

func (m *MockMailer) InviteMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) ConfirmationMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) RecoveryMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) MagicLinkMail(r *http.Request, user *models.User, otp, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) EmailChangeMail(r *http.Request, user *models.User, otpNew, otpCurrent, referrerURL string, externalURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) ReauthenticateMail(r *http.Request, user *models.User, otp string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) GetEmailActionLink(user *models.User, actionType, referrerURL string, externalURL *url.URL) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MockMailer) PasswordChangedNotificationMail(r *http.Request, user *models.User) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) EmailChangedNotificationMail(r *http.Request, user *models.User, oldEmail string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) PhoneChangedNotificationMail(r *http.Request, user *models.User, oldPhone string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) IdentityLinkedNotificationMail(r *http.Request, user *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) IdentityUnlinkedNotificationMail(r *http.Request, user *models.User, provider string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) MFAFactorEnrolledNotificationMail(r *http.Request, user *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) MFAFactorUnenrolledNotificationMail(r *http.Request, user *models.User, factorType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMailer) Reset() { _ = "STUB: not implemented"; return }
