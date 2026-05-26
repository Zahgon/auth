package v0hooks

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/hooks/hookshttp"
	"github.com/supabase/auth/internal/hooks/hookspgfunc"
	"github.com/supabase/auth/internal/storage"
)

type Manager struct {
	config *conf.GlobalConfiguration
	http   *hookshttp.Dispatcher
	pgfunc *hookspgfunc.Dispatcher
}

func NewManager(
	config *conf.GlobalConfiguration,
	httpDr *hookshttp.Dispatcher,
	pgfuncDr *hookspgfunc.Dispatcher,
) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (o *Manager) Enabled(name Name) bool { _ = "STUB: not implemented"; return false }

func configByName(
	cfg *conf.HookConfiguration,
	name Name,
) (*conf.ExtensibilityPointConfiguration, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (o *Manager) InvokeHook(
	conn *storage.Connection,
	r *http.Request,
	input, output any,
) error {
	_ = "STUB: not implemented"
	return nil
}

// invokeHook invokes the hook code. conn can be nil, in which case a new
// transaction is opened. If calling invokeHook within a transaction, always
// pass the current transaction, as pool-exhaustion deadlocks are very easy to
// trigger.
func (o *Manager) invokeHook(
	conn *storage.Connection,
	r *http.Request,
	input, output any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Manager) dispatch(
	ctx context.Context,
	hookConfig *conf.ExtensibilityPointConfiguration,
	conn *storage.Connection,
	input, output any,
) error {
	_ = "STUB: not implemented"
	return nil
}
