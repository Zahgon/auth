package observability

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
)

const (
	LOG_SQL_ALL       = "all"
	LOG_SQL_NONE      = "none"
	LOG_SQL_STATEMENT = "statement"
)

var (
	loggingOnce sync.Once
)

type CustomFormatter struct {
	logrus.JSONFormatter
}

func NewCustomFormatter() *CustomFormatter { _ = "STUB: not implemented"; return nil }

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	// logrus doesn't support formatting the time in UTC so we need to use a custom formatter
	return nil, nil
}

func ConfigureLogging(config *conf.LoggingConfig) error { _ = "STUB: not implemented"; return nil }

// use a file if you want

//#nosec G302 -- Log files should be rw-rw-r--

func setPopLogger(sql string) { _ = "STUB: not implemented"; return }

// Special case SQL logging since we have 2 extra flags to check
