package apiworker

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/mailer/templatemailer"
	"github.com/supabase/auth/internal/storage"
)

// Worker is a simple background worker for async tasks.
type Worker struct {
	le *logrus.Entry
	tc *templatemailer.Cache
	db *storage.Connection

	// Notifies worker the cfg has been updated.
	cfgCh chan struct{}

	// workMu must be held for calls to Work
	workMu sync.Mutex

	// mu must be held for field access below here
	mu  sync.Mutex
	cfg *conf.GlobalConfiguration
}

// New will return a new *Worker instance.
func New(
	cfg *conf.GlobalConfiguration,
	tc *templatemailer.Cache,
	db *storage.Connection,
	le *logrus.Entry,
) *Worker {
	_ = "STUB: not implemented"
	return nil
}

func (o *Worker) putConfig(cfg *conf.GlobalConfiguration) { _ = "STUB: not implemented"; return }

func (o *Worker) getConfig() *conf.GlobalConfiguration { _ = "STUB: not implemented"; return nil }

// ReloadConfig notifies the worker a new configuration is available.
func (o *Worker) ReloadConfig(cfg *conf.GlobalConfiguration) { _ = "STUB: not implemented"; return }

// Work will run background workers.
func (o *Worker) Work(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (o *Worker) configNotifier(
	ctx context.Context,
	notifyCh ...chan<- struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// When we get a config update, notify each worker to wake up

func (o *Worker) dbWorker(ctx context.Context, cfgCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// templateWorker will periodically reload the templates in the background as
// long as the system remains active.
func (o *Worker) templateWorker(ctx context.Context, cfgCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Reload templates right away on Work.

// Either ticker fired or we got a config update.

func (o *Worker) maybeReloadTemplates(
	ctx context.Context,
	cfg *conf.GlobalConfiguration,
) {
	_ = "STUB: not implemented"
	return
}

func (o *Worker) indexWorker(ctx context.Context, cfgCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Worker) maybeCreateIndexes(
	ctx context.Context,
	cfg *conf.GlobalConfiguration,
	le *logrus.Entry,
) {
	_ = "STUB: not implemented"
	return
}
