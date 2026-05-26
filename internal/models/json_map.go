package models

import (
	"database/sql/driver"
)

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (j JSONMap) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }
