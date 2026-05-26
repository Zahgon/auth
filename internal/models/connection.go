package models

import (
	"github.com/supabase/auth/internal/storage"
)

type Pagination struct {
	Page    uint64
	PerPage uint64
	Count   uint64
}

func (p *Pagination) Offset() uint64 { _ = "STUB: not implemented"; return 0 }

type SortDirection string

const Ascending SortDirection = "ASC"
const Descending SortDirection = "DESC"
const CreatedAt = "created_at"

type SortParams struct {
	Fields []SortField
}

type SortField struct {
	Name string
	Dir  SortDirection
}

// TruncateAll deletes all data from the database, as managed by GoTrue. Not
// intended for use outside of tests.
func TruncateAll(conn *storage.Connection) error { _ = "STUB: not implemented"; return nil }
