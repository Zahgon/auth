package reloader

import (
	"context"
	"io/fs"
	"sync"
	"time"
)

const (
	pollerMaxScan  = 1
	pollerMaxFiles = 1000
)

type poller struct {
	pollMu    sync.Mutex
	dir       string
	cur, prev *pollerState
}

type pollerFile struct {
	name string
	size int64
	mode fs.FileMode
	mod  time.Time
	dir  bool
}

type pollerState struct {
	updatedAt time.Time
	files     map[string]*pollerFile
}

func (o *pollerState) reset() { _ = "STUB: not implemented"; return }

func newPollerState() *pollerState { _ = "STUB: not implemented"; return nil }

func newPollerFile(fi fs.FileInfo) *pollerFile { _ = "STUB: not implemented"; return nil }

func newPoller(watchDir string) *poller { _ = "STUB: not implemented"; return nil }

func (o *poller) watch(
	ctx context.Context,
	ival time.Duration,
	notifyFn func(),
	errFn func(error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *poller) poll(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (o *poller) scan(
	ctx context.Context,
	ps *pollerState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *poller) scanFile(
	ctx context.Context,
	ps *pollerState,
	f fs.ReadDirFile,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *poller) scanEntries(ps *pollerState, ents []fs.DirEntry) {
	_ = "STUB: not implemented"
	return
}
