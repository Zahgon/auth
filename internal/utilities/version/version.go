package version

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

func InitVersionMetrics(ctx context.Context, ver string) error {
	_ = "STUB: not implemented"
	return nil
}

func initGauge(
	ctx context.Context,
	typ string,
	val uint64,
	gaugeFunc initGaugeFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

type initGaugeFunc func(
	name string,
	options ...metric.Int64GaugeOption,
) (metric.Int64Gauge, error)

func initGaugeOtel(name string, options ...metric.Int64GaugeOption) (metric.Int64Gauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge), nil
}

func initMetrics(ctx context.Context, vi *versionInfo) error { _ = "STUB: not implemented"; return nil }

type versionInfo struct {
	Original string
	Major    uint64
	Minor    uint64
	Patch    uint64
	RC       uint64
}

func parseSemver(ver string) (*versionInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func normalizeVersion(ver string) string { _ = "STUB: not implemented"; return "" }
