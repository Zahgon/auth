// Package apitask provides a background execution context for background work
// that limits the execution time to the current request.
package apitask

import (
	"context"
	"errors"
	"net/http"
	"sync"
)

// ErrTask is the base of all errors originating from apitasks.
var ErrTask = errors.New("apitask")

// Middleware wraps next with an http.Handler which adds apitasks handling
// to the request context and waits for all tasks to exit before returning.
func Middleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Task is implemented by objects which may be ran in the background.
type Task interface {

	// Type return a basic name for a task. It is not expected to be consistent
	// with the underlying type, but it should be low cardinality.
	Type() string

	// Run should run this task.
	Run(context.Context) error
}

type taskFunc struct {
	typ string
	fn  func(context.Context) error
}

func (o *taskFunc) Type() string { _ = "STUB: not implemented"; return "" }

func (o *taskFunc) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func Func(typ string, fn func(context.Context) error) Task {
	_ = "STUB: not implemented"
	return *new(Task)
}

// Run will run a request-scoped background task in a separate goroutine
// immediately if the current context supports it. Otherwise it makes an
// immediate blocking call to task.Run(ctx).
//
// It is invalid to call Run within a tasks Run method.
func Run(ctx context.Context, task Task) error { _ = "STUB: not implemented"; return nil }

// Wait will wait for all currently running request-scoped background tasks to
// complete before returning.
func Wait(ctx context.Context) { _ = "STUB: not implemented"; return }

// With sets up the given context for adding request-scoped background tasks.
func With(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var ctxKey = new(int)

func from(ctx context.Context) (*requestWorker, bool) { _ = "STUB: not implemented"; return nil, false }

type requestWorker struct {
	mu   sync.Mutex
	wg   sync.WaitGroup
	done bool
}

func (o *requestWorker) wait() { _ = "STUB: not implemented"; return }

func (o *requestWorker) run(ctx context.Context, task Task) error {
	_ = "STUB: not implemented"
	return nil
}
