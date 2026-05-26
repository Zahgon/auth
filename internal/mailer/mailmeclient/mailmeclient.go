// Package mailmeclient provides an implementation of mailer.Client that uses
// gopkg.in/gomail.v2 to send via SMTP.
package mailmeclient

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
)

// Client lets MailMe send templated mails
type Client struct {
	From      string
	Host      string
	Port      int
	User      string
	Pass      string
	LocalName string

	Logger      logrus.FieldLogger
	MailLogging bool
}

// New returns a new *Mailer based on the given configuration.
func New(globalConfig *conf.GlobalConfiguration) *Client { _ = "STUB: not implemented"; return nil }

// Mail sends a templated mail. It will try to load the template from a URL, and
// otherwise fall back to the default
func (m *Client) Mail(
	ctx context.Context,
	to string,
	subject string,
	body string,
	headers map[string][]string,
	typ string,
) error {
	_ = "STUB: not implemented"
	return nil
}
