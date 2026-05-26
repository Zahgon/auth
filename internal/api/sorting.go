package api

import (
	"net/http"

	"github.com/supabase/auth/internal/models"
)

func sort(r *http.Request, allowedFields map[string]bool, defaultSort []models.SortField) (*models.SortParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
