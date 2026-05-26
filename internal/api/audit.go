package api

import (
	"net/http"
)

var filterColumnMap = map[string][]string{
	"author": {"actor_username", "actor_name"},
	"action": {"action"},
	"type":   {"log_type"},
}

func (a *API) adminAuditLog(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// aud := a.requestAud(ctx, r)
