package models

import (
	"sync/atomic"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

type Cleaner interface {
	Clean(*storage.Connection) (int, error)
}

type Cleanup struct {
	cleanupStatements []string

	// cleanupNext holds an atomically incrementing value that determines which of
	// the cleanupStatements will be run next.
	cleanupNext uint32

	// cleanupAffectedRows tracks an OpenTelemetry metric on the total number of
	// cleaned up rows.
	cleanupAffectedRows atomic.Int64
}

func NewCleanup(config *conf.GlobalConfiguration) *Cleanup { _ = "STUB: not implemented"; return nil }

// These statements intentionally use SELECT ... FOR UPDATE SKIP LOCKED
// as this makes sure that only rows that are not being used in another
// transaction are deleted. These deletes are thus very quick and
// efficient, as they don't wait on other transactions.

// sessions are deleted after 72 hours to allow refresh tokens
// to be deleted piecemeal; 10 at once so that cascades don't
// overwork the database

// delete anonymous users older than 30 days

// delete sessions with a refreshed_at column

// delete sessions without a refreshed_at column by looking for
// unrevoked refresh_tokens

// Cleanup removes stale entities in the database. You can call it on each
// request or as a periodic background job. It does quick lockless updates or
// deletes, has an execution timeout and acquire timeout so that cleanups do
// not affect performance of other database jobs. Note that calling this does
// not clean up the whole database, but does a small piecemeal clean up each
// time when called.
func (c *Cleanup) Clean(db *storage.Connection) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// #nosec G115
