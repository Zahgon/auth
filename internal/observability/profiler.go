package observability

import (
	"context"

	"net/http"

	"github.com/supabase/auth/internal/conf"
)

func ConfigureProfiler(ctx context.Context, pc *conf.ProfilerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G118 -- cancel() is called in the shutdown goroutine below; baseContext is for the HTTP server.

// close baseContext

type ProfilerHandler struct{}

func (p *ProfilerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
