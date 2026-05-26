// Package e2e provides a few utilities for use in unit tests.
package e2e

import (
	"testing"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

var (
	projectRoot string
	configPath  string
)

var isTesting func() bool = testing.Testing

func init() {
	initPackage()
}

func initPackage() { _ = "STUB: not implemented"; return }

// GetProjectRoot returns the path to the root of the project. This may be used
// to locate files without needing the relative path from a given test.
func GetProjectRoot() string {
	_ = "STUB: not implemented"

	// GetConfigPath returns the path for the "/hack/test.env" config file.
	return ""
}

func GetConfigPath() string {
	_ = "STUB: not implemented"

	// Config calls conf.LoadGlobal using GetConfigPath().
	return ""
}

func Config() (*conf.GlobalConfiguration, error) { _ = "STUB: not implemented"; return nil, nil }

// Conn returns a connection for the given config.
func Conn(globalCfg *conf.GlobalConfiguration) (*storage.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Must may be used by Config and Conn, i.e.:
//
//	cfg := e2e.Must(e2e.Config())
//	conn := e2e.Must(e2e.Conn(cfg))
func Must[T any](res T, err error) T { _ = "STUB: not implemented"; return *new(T) }
