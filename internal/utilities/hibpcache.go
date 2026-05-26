package utilities

import (
	"context"
	"sync"

	"github.com/bits-and-blooms/bloom/v3"
)

const (
	// hibpHashLength is the length of a hex-encoded SHA1 hash.
	hibpHashLength = 40
	// hibpHashPrefixLength is the length of the hashed password prefix.
	hibpHashPrefixLength = 5
)

type HIBPBloomCache struct {
	sync.RWMutex

	n      uint
	items  uint
	filter *bloom.BloomFilter
}

func NewHIBPBloomCache(n uint, fp float64) *HIBPBloomCache { _ = "STUB: not implemented"; return nil }

func (c *HIBPBloomCache) Cap() uint { _ = "STUB: not implemented"; return 0 }

func (c *HIBPBloomCache) Add(ctx context.Context, prefix []byte, suffixes [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// clear the filter if 80% full to keep the actual false
// positive rate low

// reduce memory footprint when this happens

func (c *HIBPBloomCache) Contains(ctx context.Context, prefix, suffix []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
