package observability

import (
	"context"
	"sync"

	"github.com/supabase/auth/internal/conf"

	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

func Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func openTelemetryResource() *sdkresource.Resource { _ = "STUB: not implemented"; return nil }

func enableOpenTelemetryTracing(ctx context.Context, tc *conf.TracingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// http/json for example

// Register the W3C trace context and baggage propagators so data is
// propagated across services/processes

// #nosec G118 -- Cleanup goroutine intentionally outlives the request; context.Background() is required for shutdown after parent context is cancelled.

var (
	tracingOnce sync.Once
)

// ConfigureTracing sets up global tracing configuration for OpenTracing /
// OpenTelemetry. The context should be the global context. Cancelling this
// context will cancel tracing collection.
func ConfigureTracing(ctx context.Context, tc *conf.TracingConfig) error {
	_ = "STUB: not implemented"
	return nil
}
