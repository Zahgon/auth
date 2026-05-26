package reloader

import (
	"net/http"
	"sync/atomic"
)

// AtomicHandler provides an atomic http.Handler implementation, allowing safe
// handler replacement at runtime. AtomicHandler must be initialized with a call
// to NewAtomicHandler. It will never panic and is safe for concurrent use.
type AtomicHandler struct {
	val atomic.Value
}

// atomicHandlerValue is the value stored within an atomicHandler.
type atomicHandlerValue struct{ http.Handler }

// NewAtomicHandler creates a new AtomicHandler ready for use.
func NewAtomicHandler(h http.Handler) *AtomicHandler { _ = "STUB: not implemented"; return nil }

// String implements fmt.Stringer by returning a string literal.
func (ah *AtomicHandler) String() string { _ = "STUB: not implemented"; return "" }

// Store will update this http.Handler to serve future requests using h.
func (ah *AtomicHandler) Store(h http.Handler) { _ = "STUB: not implemented"; return }

// load will return the underlying http.Handler used to serve requests.
func (ah *AtomicHandler) load() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// ServeHTTP implements the standard libraries http.Handler interface by
// atomically passing the request along to the most recently stored handler.
func (ah *AtomicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
