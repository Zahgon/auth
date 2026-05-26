package hookshttp

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/conf"
)

const (
	defaultHTTPHookTimeout  = 5 * time.Second
	defaultHTTPHookRetries  = 3
	httpHookBackoffDuration = 2 * time.Second
	payloadLimit            = 200 * 1024 // 200KB
)

type Dispatcher struct {
	hookTimeout   time.Duration
	hookBackoff   time.Duration
	hookRetries   int
	limitResponse int64
}

type Option interface {
	apply(*Dispatcher)
}

type optionFunc func(*Dispatcher)

func (f optionFunc) apply(o *Dispatcher) { _ = "STUB: not implemented"; return }

func WithTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBackoff(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetries(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResponseLimit(n int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(opts ...Option) *Dispatcher { _ = "STUB: not implemented"; return nil }

func (o *Dispatcher) Dispatch(
	ctx context.Context,
	cfg *conf.ExtensibilityPointConfiguration,
	req any,
	res any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Dispatcher) runHTTPHook(
	ctx context.Context,
	hookConfig *conf.ExtensibilityPointConfiguration,
	input any,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By default, Go Client sets encoding to gzip, which does not carry a content length header.

// check if the response body still has excess bytes to be read

// Check for truthy values to allow for flexibility to switch to time duration

// Previously this returned nil, nil when retryAfterHeader was present

func generateSignatures(
	secrets []string,
	msgID uuid.UUID,
	currentTime time.Time,
	inputPayload []byte,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(joel): Handle asymmetric case once library has been upgraded

// Note this function as implemented always returns a nil error.
