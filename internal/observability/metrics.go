package observability

import (
	"context"
	"sync"

	"github.com/supabase/auth/internal/conf"

	"go.opentelemetry.io/otel/metric"
)

func Meter(instrumentationName string, opts ...metric.MeterOption) metric.Meter {
	_ = "STUB: not implemented"
	return *new(metric.Meter)
}

func ObtainMetricCounter(name, desc string) metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func enablePrometheusMetrics(ctx context.Context, mc *conf.MetricsConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G118 -- Cleanup goroutine intentionally manages server lifecycle independent of request context.

// #nosec G118 -- cancel() is called in the shutdown goroutine below; baseContext is for the HTTP server.

// to mitigate a Slowloris attack

// close baseContext

func enableOpenTelemetryMetrics(ctx context.Context, mc *conf.MetricsConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G118 -- Cleanup goroutine intentionally outlives the request; context.Background() is required for shutdown after parent context is cancelled.

// #nosec G118 -- Cleanup goroutine intentionally outlives the request; context.Background() is required for shutdown after parent context is cancelled.

// http/json for example

var (
	metricsOnce *sync.Once = &sync.Once{}
)

func ConfigureMetrics(ctx context.Context, mc *conf.MetricsConfig) error {
	_ = "STUB: not implemented"
	return nil
}
