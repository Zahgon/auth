package observability

import (
	"context"
	"sync"
)

var (
	cleanupWaitGroup sync.WaitGroup
)

// WaitForCleanup waits until all observability long-running goroutines shut
// down cleanly or until the provided context signals done.
func WaitForCleanup(ctx context.Context) { _ = "STUB: not implemented"; return }
