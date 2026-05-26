package api

import (
	"time"
)

const APIVersionHeaderName = "X-Supabase-Api-Version"

type APIVersion = time.Time

var (
	APIVersionInitial  = time.Time{}
	APIVersion20240101 = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
)

func DetermineClosestAPIVersion(date string) (APIVersion, error) {
	_ = "STUB: not implemented"
	return *new(APIVersion), nil
}

func FormatAPIVersion(apiVersion APIVersion) string { _ = "STUB: not implemented"; return "" }
