package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
)

// Connection is the interface a storage provider must implement. Do not copy
// a storage connection
type Connection struct {
	*pop.Connection
	sqldb *sql.DB
}

// Dial will connect to that storage engine
func Dial(config *conf.GlobalConfiguration) (*Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialContext(
	ctx context.Context,
	config *conf.GlobalConfiguration,
) (*Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// // GetSqlDB returns the underlying *sql.DB and true or nil if no db could be obtained.
// func (c *Connection) GetSqlDB() (*sql.DB, bool) { return c.sqldb, c.sqldb != nil }

// Copy will return a copy of this connection. It must be instead of using a
// struct literal from external packages.
func (c *Connection) Copy() *Connection { _ = "STUB: not implemented"; return nil }

func newConnectionDetails(
	config *conf.GlobalConfiguration,
) (*pop.ConnectionDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(cstockton): I'm preserving the Mutation here for now because I'm not
// sure what side effects changing this could have. But it should probably go
// inside the Validate() function in conf package or somewhere else.
func applyDBDriver(
	config *conf.GlobalConfiguration,
	cd *pop.ConnectionDetails,
) error {
	_ = "STUB: not implemented"
	return nil
}

// pop v5 uses pgx as the default PostgreSQL driver

// sqlx needs to be informed that the new instrumented
// driver has the same semantics as the
// non-instrumented driver

// NOTE: I couldn't find any way to obtain the store when wrapped with context
// due to the private store field in pop.contextStore.
func popConnToStd(db *pop.Connection) (sqldb *sql.DB, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Get element stored in the pop.store interface within field db.Store.
// *pop.dB

// dbval should contain a pointer to struct with layout of pop.dB:
//
//   type dB struct {
//     *sqlx.DB
//   }
//
// *sqlx.DB

// dbval should now be a pointer to a struct with layout like sqlx.DB:
//
//   type DB struct {
//     *sql.DB
//   }
//
// *sql.DB

// dbval should now be (*sql.DB) get an iface and try to cast.

// ApplyConfig will apply the given config to this *Connection, potentially
// adjusting the underlying *sql.DB's current settings.
//
// When config.DB.ConnPercentage is set to a non-zero value ApplyConfig attempts
// to set the MaxOpenConns and MaxIdleConns to a percentage based value. It does
// this by opening a connection to the server and calling
// `SHOW max_connections;` to determine the connection limits. If this operation
// fails it applies no configuration changes at all and returns an error.
func (c *Connection) ApplyConfig(
	ctx context.Context,
	config *conf.GlobalConfiguration,
	le *logrus.Entry,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Config values

// Server values

// Limit values

func (c *Connection) getConnLimits(
	ctx context.Context,
	dbCfg *conf.DBConfiguration,
) (*ConnLimits, error) {
	_ = "STUB: not implemented"
	// Set the connection limits to the fixed values in config
	return nil, nil
}

// Always fetch max conns because it is useful for logging.

// pct based conn limits are disabled

// pct conn limits are enabled, try to determine what they should be

func (c *Connection) applyPercentageLimits(
	dbCfg *conf.DBConfiguration,
	maxConns int,
	cl *ConnLimits,
) error {
	_ = "STUB: not implemented"
	return nil
	// set this here too for unit tests
}

// pct based conn limits are disabled

// If maxConns is 0 it means our role or db is not allowing conns right
// now and we do nothing.

// Ensure the conn pct isn't OOB

// maxConns > 0 so we may calculate the percentage.

// We set max idle conns to the max open conns.

// return the percentage based conn limits

// showMaxConns retrieves the max_connections from the db.
func (c *Connection) showMaxConns(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const (
	connLimitsErrorStrategy      = "error"
	connLimitsFixedStrategy      = "fixed"
	connLimitsPercentageStrategy = "percentage"
)

// ConnLimits represents the connection limits for the underlying *sql.DB.
type ConnLimits struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	ServerMaxConns  int
	Strategy        string
}

func newConnLimitsFromConfig(dbCfg *conf.DBConfiguration) *ConnLimits {
	_ = "STUB: not implemented"
	return nil
}

func registerOpenTelemetryDatabaseStats(config *conf.GlobalConfiguration, sqldb *sql.DB) {
	_ = "STUB: not implemented"
	return
}

type CommitWithError struct {
	Err error
}

func (e *CommitWithError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *CommitWithError) Cause() error {
	_ = "STUB: not implemented"

	// NewCommitWithError creates an error that can be returned in a pop transaction
	// without rolling back the transaction. This should only be used in cases where
	// you want the transaction to commit but return an error message to the user.
	return nil
}

func NewCommitWithError(err error) *CommitWithError { _ = "STUB: not implemented"; return nil }

func (c *Connection) Transaction(fn func(*Connection) error) error {
	_ = "STUB: not implemented"
	return nil
}

// there exists a race condition when the context deadline is exceeded
// and whether the transaction has been committed or not
// e.g. if the context deadline has exceeded but the transaction has already been committed,
// it won't be possible to perform a rollback on the transaction since the transaction has been closed

// WithContext returns a new connection with an updated context. This is
// typically used for tracing as the context contains trace span information.
func (c *Connection) WithContext(ctx context.Context) *Connection {
	_ = "STUB: not implemented"
	return nil
}

func getExcludedColumns(model interface{}, includeColumns ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get all columns and remove included to get excluded set

// gobuffalo updates the updated_at column automatically
