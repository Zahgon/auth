package ratelimit

import (
	"sync"
	"time"

	"github.com/supabase/auth/internal/conf"
)

// IntervalLimiter will limit the number of calls to Allow per interval.
type IntervalLimiter struct {
	mu     sync.Mutex
	cfg    conf.Rate
	ival   time.Duration // Count is reset and time updated every ival.
	events int           // Events allowed per ival.

	// Guarded by mu.
	last  time.Time // When the limiter was last reset.
	count int       // Total calls to Allow() since time.
}

// NewIntervalLimiter returns a rate limiter using the given conf.Rate.
func NewIntervalLimiter(r conf.Rate) *IntervalLimiter { _ = "STUB: not implemented"; return nil }

// Allow implements Limiter by calling AllowAt with the current time.
func (rl *IntervalLimiter) Allow() bool { _ = "STUB: not implemented"; return false }

// AllowAt implements Limiter by checking if the current number of permitted
// events within this interval would permit 1 additional event at the current
// time.
//
// When called with a time outside the current active interval the counter is
// reset, meaning it can be vulnerable at the edge of it's intervals so avoid
// small intervals.
func (rl *IntervalLimiter) AllowAt(at time.Time) bool { _ = "STUB: not implemented"; return false }

func (rl *IntervalLimiter) allowAt(at time.Time) bool { _ = "STUB: not implemented"; return false }

func (l *IntervalLimiter) String() string { _ = "STUB: not implemented"; return "" }

func (l *IntervalLimiter) Config() conf.Rate { _ = "STUB: not implemented"; return *new(conf.Rate) }
