package hookspgfunc

import (
	"context"
	"time"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

const (
	defaultTimeout = time.Second * 2
)

type Dispatcher struct {
	db            *storage.Connection
	hookTimeoutMS int
}

type Option interface {
	apply(*Dispatcher)
}

type optionFunc func(*Dispatcher)

func (f optionFunc) apply(o *Dispatcher) { _ = "STUB: not implemented"; return }

func WithTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(db *storage.Connection, opts ...Option) *Dispatcher { _ = "STUB: not implemented"; return nil }

func (o *Dispatcher) Dispatch(
	ctx context.Context,
	cfg *conf.ExtensibilityPointConfiguration,
	tx *storage.Connection,
	req, res any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Dispatcher) runPostgresHook(
	ctx context.Context,
	hookConfig conf.ExtensibilityPointConfiguration,
	tx *storage.Connection,
	input any,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We rely on Postgres timeouts to ensure the function doesn't overrun

// reset the timeout
