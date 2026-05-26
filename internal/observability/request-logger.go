package observability

import (
	"context"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
)

func AddRequestID(globalConfig *conf.GlobalConfiguration) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func NewStructuredLogger(logger *logrus.Logger, config *conf.GlobalConfiguration) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type structuredLogger struct {
	Logger *logrus.Logger
	Config *conf.GlobalConfiguration
}

func (l *structuredLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	_ = "STUB: not implemented"
	return *new(chimiddleware.LogEntry)
}

// logEntry implements the chiMiddleware.LogEntry interface
type logEntry struct {
	Entry *logrus.Entry
}

// NewLogEntry returns a new chimiddleware.LogEntry from a *logrus.Entry.
func NewLogEntry(le *logrus.Entry) chimiddleware.LogEntry {
	_ = "STUB: not implemented"
	return *new(chimiddleware.LogEntry)
}

func (e *logEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	_ = "STUB: not implemented"
	return
}

func (e *logEntry) Panic(v interface{}, stack []byte) { _ = "STUB: not implemented"; return }

func GetLogEntry(r *http.Request) *logEntry { _ = "STUB: not implemented"; return nil }

func GetLogEntryFromContext(ctx context.Context) *logEntry { _ = "STUB: not implemented"; return nil }

func SetLogEntryWithContext(
	ctx context.Context,
	entry chimiddleware.LogEntry,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func LogEntrySetField(r *http.Request, key string, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func LogEntrySetFields(r *http.Request, fields logrus.Fields) { _ = "STUB: not implemented"; return }
