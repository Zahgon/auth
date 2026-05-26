package utilities

// PostgresError is a custom error struct for marshalling Postgres errors to JSON.
type PostgresError struct {
	Code           string `json:"code"`
	HttpStatusCode int    `json:"-"`
	Message        string `json:"message"`
	Hint           string `json:"hint,omitempty"`
	Detail         string `json:"detail,omitempty"`
}

// NewPostgresError returns a new PostgresError if the error was from a publicly
// accessible Postgres error.
func NewPostgresError(err error) *PostgresError { _ = "STUB: not implemented"; return nil }

func (pg *PostgresError) IsUniqueConstraintViolated() bool {
	_ = "STUB: not implemented"
	// See https://www.postgresql.org/docs/current/errcodes-appendix.html for list of error codes
	return false
}

// isPubliclyAccessiblePostgresError checks if the Postgres error should be
// made accessible.
func isPubliclyAccessiblePostgresError(code string) bool { _ = "STUB: not implemented"; return false }

// default response

// getHttpStatusCodeFromPostgresErrorCode maps a Postgres error code to a HTTP
// status code. Returns 0 if the code doesn't map to a given postgres error code.
func getHttpStatusCodeFromPostgresErrorCode(code string) int { _ = "STUB: not implemented"; return 0 }

// Use custom HTTP status code if Postgres error was triggered with `PTXXX`
// code. This is consistent with PostgREST's behaviour as well.
