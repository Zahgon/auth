// Package e2eapi provides utilities for end-to-end testing the api.
package e2eapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/supabase/auth/internal/api"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
	"github.com/supabase/auth/internal/storage/test"
	"github.com/supabase/auth/internal/utilities"
)

type Instance struct {
	Config    *conf.GlobalConfiguration
	Conn      *storage.Connection
	APIServer *httptest.Server
	APIClient *http.Client
	apiURL    *url.URL

	closers []func()
}

func New(globalCfg *conf.GlobalConfiguration) (*Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Instance) init() error {
	conn, err := test.SetupDBConnection(o.Config)
	if err != nil {
		return fmt.Errorf("error setting up db connection: %w", err)
	}
	o.addCloser(func() {
		if conn.Store != nil {
			_ = conn.Close()
		}
	})
	o.Conn = conn

	apiVer := utilities.Version
	if apiVer == "" {
		apiVer = "1"
	}

	a := api.NewAPIWithVersion(o.Config, conn, apiVer)
	apiSrv := httptest.NewServer(a)
	o.addCloser(apiSrv)
	o.APIServer = apiSrv
	o.APIClient = apiSrv.Client()

	return o.initURL()
}

func (o *Instance) initURL() error { _ = "STUB: not implemented"; return nil }

func (o *Instance) Close() error { _ = "STUB: not implemented"; return nil }

func (o *Instance) addCloser(v any) { _ = "STUB: not implemented"; return }

func (o *Instance) Do(
	req *http.Request,
) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Instance) DoAuth(
	req *http.Request,
	jwt string,
) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Instance) DoAdmin(
	req *http.Request,
) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Instance) doAdmin(
	req *http.Request,
	key any,
) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Do(
	ctx context.Context,
	method string,
	url string,
	req, res any,
) error {
	_ = "STUB: not implemented"
	return nil
}

const responseLimit = 1e6

var defaultClient = http.DefaultClient

func do(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
