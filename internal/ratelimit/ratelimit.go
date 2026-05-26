package ratelimit

import (
	"time"

	"github.com/supabase/auth/internal/conf"
)

// Limiter is the interface implemented by rate limiters.
//
// Implementations of Limiter must be safe for concurrent use.
type Limiter interface {

	// Allow should return true if an event should be allowed at the time
	// which it was called, or false otherwise.
	Allow() bool

	// AllowAt should return true if an event should be allowed at the given
	// time, or false otherwise.
	AllowAt(at time.Time) bool

	// Config returns the underlying config value
	Config() conf.Rate
}

// Equal checks to see if two limiters / vals / cfgs are both valid and equal.
func Equal(a, b any) bool { _ = "STUB: not implemented"; return false }

func toRate(v any) conf.Rate { _ = "STUB: not implemented"; return *new(conf.Rate) }

// New returns a new Limiter based on the given config.
//
// When the type is conf.BurstRateType it returns a BurstLimiter, otherwise
// New returns an IntervalLimiter.
func New(r conf.Rate) Limiter { _ = "STUB: not implemented"; return *new(Limiter) }
