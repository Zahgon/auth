package ratelimit

import (
	"time"

	"github.com/supabase/auth/internal/conf"
	"golang.org/x/time/rate"
)

const defaultOverTime = time.Hour

// BurstLimiter wraps the golang.org/x/time/rate package.
type BurstLimiter struct {
	rl  *rate.Limiter
	cfg conf.Rate
}

// NewBurstLimiter returns a rate limiter configured using the given conf.Rate.
//
// The returned Limiter will be configured with a token bucket containing a
// single token, which will fill up at a rate of 1 event per r.OverTime with
// an initial burst amount of r.Events.
//
// For example:
//   - 1/10s  is 1 events per 10 seconds with burst of 1.
//   - 1/2s   is 1 events per 2  seconds with burst of 1.
//   - 10/10s is 1 events per 10 seconds with burst of 10.
//
// If Rate.Events is <= 0, the burst amount will be set to 0. Since b is the
// initial size of the token bucket this means no events will be permitted at
// all.
//
// See Example_newBurstLimiter for a visualization.
func NewBurstLimiter(r conf.Rate) *BurstLimiter {
	_ = "STUB: not implemented"
	// The rate limiter deals in events per second.
	return nil
}

// BurstLimiter will have an initial token bucket of size `e`. It will
// be refilled at a rate of 1 per duration `d` indefinitely.

// Allow implements Limiter by calling AllowAt with the current time.
func (l *BurstLimiter) Allow() bool { _ = "STUB: not implemented"; return false }

// AllowAt implements Limiter by calling the underlying x/time/rate.Limiter
// with the given time.
func (l *BurstLimiter) AllowAt(at time.Time) bool { _ = "STUB: not implemented"; return false }

func (l *BurstLimiter) String() string { _ = "STUB: not implemented"; return "" }

func (l *BurstLimiter) Config() conf.Rate { _ = "STUB: not implemented"; return *new(conf.Rate) }
