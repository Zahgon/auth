package utilities

import (
	"context"
	"sync"
)

type contextKey string

func (c contextKey) String() string { _ = "STUB: not implemented"; return "" }

const (
	requestIDKey = contextKey("request_id")
)

// WithRequestID adds the provided request ID to the context.
func WithRequestID(ctx context.Context, id string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetRequestID reads the request ID from the context.
func GetRequestID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// WaitForCleanup waits until all long-running goroutines shut
// down cleanly or until the provided context signals done.
func WaitForCleanup(ctx context.Context, wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }
