package api

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/conf"
)

// OAuthProviderData contains the userData and token returned by the oauth provider
type OAuthProviderData struct {
	userData     *provider.UserProvidedData
	token        string
	refreshToken string
	code         string
}

// loadFlowState parses the `state` query parameter as a UUID,
// loads the flow state from the database, and extracts the provider requested
func (a *API) loadFlowState(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *API) oAuthCallback(ctx context.Context, r *http.Request, providerType string) (*OAuthProviderData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if there's a non-empty OAuthClientStateID we perform PKCE Flow for the external provider

// apple only returns user info the first time

func (a *API) oAuth1Callback(ctx context.Context, providerType string) (*OAuthProviderData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OAuthProvider returns the corresponding oauth provider as an OAuthProvider interface
func (a *API) OAuthProvider(ctx context.Context, name string) (provider.OAuthProvider, conf.OAuthProviderConfiguration, error) {
	_ = "STUB: not implemented"
	return *new(provider.OAuthProvider), *new(conf.OAuthProviderConfiguration), nil
}
