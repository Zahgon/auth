package storage

import (
	"database/sql/driver"
)

type NullString string

func (s *NullString) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (s NullString) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	// if nil or empty string
	return *new(driver.Value), nil
}

func (s NullString) String() string { _ = "STUB: not implemented"; return "" }
