// Package reloader provides support for live configuration reloading.
package reloader

import (
	"context"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/supabase/auth/internal/conf"
	"golang.org/x/sync/errgroup"
)

const (

	// tickerInterval is the maximum latency between configuration reloads.
	tickerInterval = time.Second
)

type ConfigFunc func(*conf.GlobalConfiguration)

type Reloader struct {
	watchDir string
	rc       conf.ReloadingConfiguration

	// Below here is for DI
	tickerIval time.Duration
	watchFn    func() (watcher, error)
	reloadFn   func(dir string) (*conf.GlobalConfiguration, error)
	addDirFn   func(ctx context.Context, wr watcher, dir string) error
}

func NewReloader(rc conf.ReloadingConfiguration, watchDir string) *Reloader {
	_ = "STUB: not implemented"
	return nil
}

// reload attempts to create a new *conf.GlobalConfiguration after loading the
// currently configured watchDir.
func (rl *Reloader) reload() (*conf.GlobalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reloadCheckAt checks if reloadConfig should be called, returns true if config
// should be reloaded or false otherwise.
func (rl *Reloader) reloadCheckAt(at, lastUpdate time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// no pending updates

// waiting for reload interval

// Update is pending.

type watchState struct {
	eg errgroup.Group
	fn ConfigFunc
	ch chan struct{}
}

func (o *watchState) notify() { _ = "STUB: not implemented"; return }

func newWatchState(
	fn ConfigFunc,
) *watchState {
	_ = "STUB: not implemented"
	return nil
}

func (rl *Reloader) Watch(ctx context.Context, fn ConfigFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (rl *Reloader) watchReloads(
	ctx context.Context,
	ws *watchState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Check to see if the config is ready to be relaoded.

// Reset the last update time before we try to reload the config.

// Call the callback function with the latest cfg.

func (rl *Reloader) watchPoller(
	ctx context.Context,
	ws *watchState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (rl *Reloader) watchSignal(
	ctx context.Context,
	ws *watchState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (rl *Reloader) watchNotify(
	ctx context.Context,
	ws *watchState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// A simple log dedupe flag to prevent endless noise if config dir missing

// Ignore errors, if watch dir doesn't exist we can add it later.

// On a supported host OS like linux the watcher creation won't fail
// but will when adding a dir with an underlying filesystem without
// notification support. Checking if the directory is watchable when
// you get an error from addDirFn is a way to detect this.

// This is a simple way to solve watch dir being added later or
// being moved and then recreated. I've tested all of these basic
// scenarios and wr.WatchList() does not grow which aligns with
// the documented behavior.

// We only read files ending in .env

func isWatchable(dir string) bool { _ = "STUB: not implemented"; return false }

// defaultAddDirFn adds a dir to a watcher with a common error and sleep
// duration if the directory doesn't exist.
func defaultAddDirFn(ctx context.Context, wr watcher, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultReloadFn(dir string) (*conf.GlobalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type watcher interface {
	Add(path string) error
	Close() error
	Events() chan fsnotify.Event
	Errors() chan error
}

type fsNotifyWatcher struct {
	wr *fsnotify.Watcher
}

func newFSWatcher() (watcher, error) { _ = "STUB: not implemented"; return *new(watcher), nil }

func (o *fsNotifyWatcher) Add(path string) error       { _ = "STUB: not implemented"; return nil }
func (o *fsNotifyWatcher) Close() error                { _ = "STUB: not implemented"; return nil }
func (o *fsNotifyWatcher) Errors() chan error          { _ = "STUB: not implemented"; return nil }
func (o *fsNotifyWatcher) Events() chan fsnotify.Event { _ = "STUB: not implemented"; return nil }

type mockWatcher struct {
	mu      sync.Mutex
	err     error
	eventCh chan fsnotify.Event
	errorCh chan error
	addCh   chan string
}

func newMockWatcher(err error) *mockWatcher { _ = "STUB: not implemented"; return nil }

func (o *mockWatcher) getErr() error { _ = "STUB: not implemented"; return nil }

func (o *mockWatcher) setErr(err error) { _ = "STUB: not implemented"; return }

func (o *mockWatcher) Add(path string) error { _ = "STUB: not implemented"; return nil }

func (o *mockWatcher) Close() error                { _ = "STUB: not implemented"; return nil }
func (o *mockWatcher) Events() chan fsnotify.Event { _ = "STUB: not implemented"; return nil }
func (o *mockWatcher) Errors() chan error          { _ = "STUB: not implemented"; return nil }
