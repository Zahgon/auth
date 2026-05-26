package api

import (
	"context"
	"net/http"
)

type Web3GrantParams struct {
	Message   string `json:"message,omitempty"`
	Signature string `json:"signature,omitempty"`
	Chain     string `json:"chain,omitempty"`
}

func (a *API) Web3Grant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) web3GrantSolana(ctx context.Context, w http.ResponseWriter, r *http.Request, params *Web3GrantParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Record login for analytics with Web3 context

func (a *API) web3GrantEthereum(ctx context.Context, w http.ResponseWriter, r *http.Request, params *Web3GrantParams) error {
	_ = "STUB: not implemented"
	return nil
}
