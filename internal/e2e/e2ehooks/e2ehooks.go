// Package e2ehooks provides utilities for end-to-end testing of hooks.
package e2ehooks

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/e2e/e2eapi"
	"github.com/supabase/auth/internal/hooks/v0hooks"
)

type Instance struct {
	*e2eapi.Instance

	HookServer   *httptest.Server
	HookRecorder *HookRecorder
}

func (o *Instance) Close() error { _ = "STUB: not implemented"; return nil }

func New(globalCfg *conf.GlobalConfiguration) (*Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HandleSuccess() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func HandleJSON(m map[string]any) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type Hook struct {
	mu    sync.Mutex
	name  v0hooks.Name
	calls []*HookCall

	hr http.Handler
}

func NewHook(name v0hooks.Name) *Hook { _ = "STUB: not implemented"; return nil }

// This hooks returns the exact claims given.

func (o *Hook) ClearCalls() { _ = "STUB: not implemented"; return }

func (o *Hook) GetCalls() []*HookCall { _ = "STUB: not implemented"; return nil }

func (o *Hook) SetHandler(hr http.Handler) { _ = "STUB: not implemented"; return }

func (o *Hook) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type HookCall struct {
	Header http.Header
	Body   string
	Dump   string
}

func (o *HookCall) Unmarshal(v any) error { _ = "STUB: not implemented"; return nil }

type HookRecorder struct {
	mux                  *http.ServeMux
	BeforeUserCreated    *Hook
	AfterUserCreated     *Hook
	CustomizeAccessToken *Hook
	MFAVerification      *Hook
	PasswordVerification *Hook
	SendEmail            *Hook
	SendSMS              *Hook
}

func NewHookRecorder() *HookRecorder { _ = "STUB: not implemented"; return nil }

func (o *HookRecorder) Register(
	hookCfg *conf.HookConfiguration,
	baseURL string,
) {
	_ = "STUB: not implemented"
	return
}

func (o *HookRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
