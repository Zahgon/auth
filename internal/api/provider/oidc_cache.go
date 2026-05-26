package provider

import (
	"context"
	"sync"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/sync/singleflight"
)

type oidcCacheEntry struct {
	provider  *oidc.Provider
	fetchedAt time.Time
}

// OIDCProviderCache caches *oidc.Provider objects keyed by issuer URL.
// This avoids re-fetching the OIDC discovery document on every request.
// JWKS fetching/caching is handled separately by the oidc library itself,
// it lazy-loads the key set on the first call to Verify and manages its own cache.
// It uses singleflight to deduplicate concurrent fetches for the same issuer.
type OIDCProviderCache struct {
	mu    sync.RWMutex
	cache map[string]*oidcCacheEntry
	sf    singleflight.Group
	ttl   time.Duration
	now   func() time.Time // injectable for tests
}

// NewOIDCProviderCache creates a new cache with the given TTL.
func NewOIDCProviderCache(ttl time.Duration) *OIDCProviderCache {
	_ = "STUB: not implemented"
	return nil
}

// GetProvider returns a cached *oidc.Provider for the given issuer, fetching
// it via oidc.NewProvider if not cached or expired. Concurrent requests for
// the same issuer are deduplicated via singleflight.
func (c *OIDCProviderCache) GetProvider(ctx context.Context, issuer string) (*oidc.Provider, error) {
	_ = "STUB: not implemented"

	// Fast path: read-lock check
	return nil, nil
}

// Slow path: singleflight fetch

// Serve stale entry if available — keeps auth working during
// transient network failures or issuer outages.

// Invalidate removes a cached provider for the given issuer.
func (c *OIDCProviderCache) Invalidate(issuer string) { _ = "STUB: not implemented"; return }

// Clear removes all cached providers.
func (c *OIDCProviderCache) Clear() { _ = "STUB: not implemented"; return }
