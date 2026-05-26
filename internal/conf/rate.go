package conf

import (
	"time"
)

const defaultOverTime = time.Hour

const (
	BurstRateType    = "burst"
	IntervalRateType = "interval"
)

type Rate struct {
	Events   float64       `json:"events,omitempty"`
	OverTime time.Duration `json:"over_time,omitempty"`
	typ      string        `json:"-"`
	val      string        `json:"-"`
}

func (r *Rate) GetRateType() string { _ = "STUB: not implemented"; return "" }

// Decode is used by envconfig to parse the env-config string to a Rate value.
func (r *Rate) Decode(value string) error { _ = "STUB: not implemented"; return nil }

// 52 because the uint needs to fit in a float64

func (r Rate) GetRateValue() string { _ = "STUB: not implemented"; return "" }

func (r *Rate) String() string { _ = "STUB: not implemented"; return "" }
