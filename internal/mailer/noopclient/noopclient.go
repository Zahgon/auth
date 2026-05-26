// Package noopclient provides an implementation of mailer.Client that simply
// does nothing.
package noopclient

import (
	"context"
	"time"
)

type Client struct {
	Delay time.Duration
}

func New() *Client { _ = "STUB: not implemented"; return nil }

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
