package storage

import (
	"database/sql"
	"time"
)

type Advisory struct {
	LongWaitDurationSamples int
	Over2WaitingSamples     int
}

type Advisor struct {
	StatsFunc  func() sql.DBStats
	AdviseFunc func(Advisory)
	Interval   time.Duration

	Stats         sql.DBStats
	LastAdvisedAt time.Time

	Iterations int

	WaitDurationSamples []time.Duration
	WaitCountSamples    []int64
}

func (a *Advisor) Start(observeDuration time.Duration) { _ = "STUB: not implemented"; return }

// after server start the db stats are going to be worse, so ignore that period

func (a *Advisor) setup(observeDuration time.Duration) { _ = "STUB: not implemented"; return }

func (a *Advisor) loop() { _ = "STUB: not implemented"; return }

// 1/3 of the observation time was spent waiting for over one sampling interval

// 1/3 of the observation time we saw more than 2 goroutines waiting for a connection
