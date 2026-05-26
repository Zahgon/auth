package api

import (
	"net/http"
	"regexp"
	"time"

	"github.com/supabase/auth/internal/api/apilimiter"
	"github.com/supabase/auth/internal/api/oauthserver"
	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/hooks/v0hooks"
	"github.com/supabase/auth/internal/mailer"
	"github.com/supabase/auth/internal/security"
	"github.com/supabase/auth/internal/storage"
	"github.com/supabase/auth/internal/tokens"
	"github.com/supabase/hibp"
)

const (
	audHeaderName  = "X-JWT-AUD"
	defaultVersion = "unknown version"
)

var bearerRegexp = regexp.MustCompile(`(?i)^bearer (\S+$)`)

// API is the main REST API
type API struct {
	handler http.Handler
	db      *storage.Connection
	config  *conf.GlobalConfiguration
	version string

	hooksMgr     *v0hooks.Manager
	hibpClient   *hibp.PwnedClient
	oauthServer  *oauthserver.Server
	tokenService *tokens.Service
	mailer       mailer.Mailer
	oidcCache    *provider.OIDCProviderCache

	captchaVerifier security.CaptchaVerifier

	// overrideTime can be used to override the clock used by handlers. Should only be used in tests!
	overrideTime func() time.Time

	limiterOpts *apilimiter.Limiter
}

func (a *API) GetConfig() *conf.GlobalConfiguration { _ = "STUB: not implemented"; return nil }
func (a *API) GetDB() *storage.Connection           { _ = "STUB: not implemented"; return nil }
func (a *API) GetTokenService() *tokens.Service     { _ = "STUB: not implemented"; return nil }
func (a *API) Mailer() mailer.Mailer                { _ = "STUB: not implemented"; return *new(mailer.Mailer) }

func (a *API) Version() string { _ = "STUB: not implemented"; return "" }

func (a *API) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewAPI instantiates a new REST API
func NewAPI(globalConfig *conf.GlobalConfiguration, db *storage.Connection, opt ...Option) *API {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) deprecationNotices() { _ = "STUB: not implemented"; return }

// NewAPIWithVersion creates a new REST API using the specified version
func NewAPIWithVersion(globalConfig *conf.GlobalConfiguration, db *storage.Connection, version string, opt ...Option) *API {
	_ = "STUB: not implemented"
	return nil
}

// Initialize token service if not provided via options

// Connect token service to API's time function (supports test overrides)

// Initialize OAuth server (only if enabled)

// all HIBP API requests should finish quickly to avoid
// unnecessary slowdowns

// 1MB

// request tracing should be added only when tracing or metrics is enabled

// Both OIDC Discovery and OAuth Authorization Server Metadata use the same unified handler
// OIDC Discovery is an extension of RFC 8414, so one response satisfies both specs

// `/authorize` to initiate OAuth2 authorization flow with the external providers
// where Supabase Auth is an OAuth2 Client

// rate limit per hour

// apply ip-based rate limiting on otps

// rate limiting applied in handler

// OAuth grant management endpoints

// Admin only oauth client management endpoints

// Manual client registration

// Custom OAuth/OIDC provider management endpoints

// supports both OAuth2 and OIDC via provider_type)
// Optional ?type=oauth2 or ?type=oidc filter
// provider_type in request body

// OAuth Dynamic Client Registration endpoint (public, rate limited)

// OAuth Token endpoint (public, with client authentication)

// OIDC UserInfo endpoint (requires user authentication via Bearer token)

// OAuth 2.1 Authorization endpoints
// `/authorize` to initiate OAuth2 authorization code flow where Supabase Auth is the OAuth2 provider

type HealthCheckResponse struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// HealthCheck endpoint indicates if the gotrue api service is available
func (a *API) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// ServeHTTP implements the http.Handler interface by passing the request along
// to its underlying Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
