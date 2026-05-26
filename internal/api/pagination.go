package api

import (
	"net/http"

	"github.com/supabase/auth/internal/models"
)

const defaultPerPage = 50

func calculateTotalPages(perPage, total uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func addPaginationHeaders(w http.ResponseWriter, r *http.Request, p *models.Pagination) {
	_ = "STUB: not implemented"
	return
}

func paginate(r *http.Request) (*models.Pagination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
