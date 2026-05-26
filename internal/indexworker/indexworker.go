package indexworker

import (
	"context"
	"errors"

	"github.com/gobuffalo/pop/v6"
	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
)

// ErrAdvisoryLockAlreadyAcquired is returned when another process already holds the advisory lock
var ErrAdvisoryLockAlreadyAcquired = errors.New("advisory lock already acquired by another process")
var ErrOrioleDBUnsupported = errors.New("index worker does not support OrioleDB tables")

type Outcome string

const (
	OutcomeSuccess Outcome = "success"
	OutcomeFailure Outcome = "failure"
	OutcomeSkipped Outcome = "skipped"
)

// CreateIndexes ensures that the necessary indexes on the users table exist.
// If the indexes already exist and are valid, it skips creation.
// It uses a Postgres advisory lock to prevent concurrent index creation
// by multiple processes.
// Returns an error either from index creation failure (partial or complete) or if the advisory lock
// could not be acquired.
func CreateIndexes(ctx context.Context, config *conf.GlobalConfiguration, le *logrus.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the users table uses OrioleDB, which does not support
// CREATE INDEX CONCURRENTLY or DROP INDEX CONCURRENTLY.

// Try to obtain advisory lock to ensure only one index worker is creating indexes at a time

// Ensure lock is released on function exit

// Check existing indexes and their statuses. If all exist and are valid, skip creation.

// When EnsureUserSearchIndexesExist is true, the user explicitly opted in, so we skip the threshold check entirely.

// First, clean up any invalid indexes from previous interrupted attempts

// Create indexes one by one

// getUsersIndexes returns the list of indexes to create on the users table
func getUsersIndexes(namespace string) []struct {
	name  string
	query string
} {
	_ = "STUB: not implemented"
	// Define indexes to create
	// Note: CONCURRENTLY cannot be used inside a transaction block
	return nil
}

// for exact-match queries, sorting, and prefix searches on email (e.g., email LIKE 'term%')

// for range queries and sorting on created_at and last_sign_in_at

// for exact-match, sorting, and prefix searches on raw_user_meta_data->>'name'

type indexStatus struct {
	IndexName string `db:"index_name"`
	IsValid   bool   `db:"is_valid"`
	IsReady   bool   `db:"is_ready"`
}

// getIndexStatuses checks the status of the given indexes in the specified namespace
func getIndexStatuses(db *pop.Connection, namespace string, indexNames []string) ([]indexStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getApproximateUserCount returns an approximate count of users in the users table to avoid a full table scan
func getApproximateUserCount(db *pop.Connection, namespace string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// isOrioleDBTable checks if the specified table uses the OrioleDB storage engine
// by examining the table's access method in pg_class.
func isOrioleDBTable(db *pop.Connection, namespace, tableName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// dropInvalidIndexes drops any invalid indexes from previous interrupted attempts
func dropInvalidIndexes(db *pop.Connection, le *logrus.Entry, namespace string, indexNames []string) {
	_ = "STUB: not implemented"
	return
}

// Query the system catalog to find invalid indexes (from interrupted CONCURRENTLY operations)
