package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func newRouter() *router { _ = "STUB: not implemented"; return nil }

type router struct {
	chi chi.Router
}

func (r *router) Route(pattern string, fn func(*router)) { _ = "STUB: not implemented"; return }

func (r *router) Get(pattern string, fn apiHandler) { _ = "STUB: not implemented"; return }

func (r *router) Post(pattern string, fn apiHandler) { _ = "STUB: not implemented"; return }

func (r *router) Put(pattern string, fn apiHandler) { _ = "STUB: not implemented"; return }

func (r *router) Patch(pattern string, fn apiHandler) { _ = "STUB: not implemented"; return }

func (r *router) Delete(pattern string, fn apiHandler) { _ = "STUB: not implemented"; return }

func (r *router) With(fn middlewareHandler) *router { _ = "STUB: not implemented"; return nil }

func (r *router) WithBypass(fn func(next http.Handler) http.Handler) *router {
	_ = "STUB: not implemented"
	return nil
}

func (r *router) Use(fn middlewareHandler) { _ = "STUB: not implemented"; return }

func (r *router) UseBypass(fn func(next http.Handler) http.Handler) {
	_ = "STUB: not implemented"
	return
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

type apiHandler func(w http.ResponseWriter, r *http.Request) error

func handler(fn apiHandler) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (h apiHandler) serve(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type middlewareHandler func(w http.ResponseWriter, r *http.Request) (context.Context, error)

func (m middlewareHandler) handler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (m middlewareHandler) serve(next http.Handler, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func middleware(fn middlewareHandler) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
