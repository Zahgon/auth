package conf

import (
	"net/url"
	"regexp"
	"text/template"
	"time"

	"github.com/gobwas/glob"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

const defaultMinPasswordLength int = 6
const defaultChallengeExpiryDuration float64 = 300
const defaultFactorExpiryDuration time.Duration = 300 * time.Second
const defaultFlowStateExpiryDuration time.Duration = 300 * time.Second

// See: https://www.postgresql.org/docs/7.0/syntax525.htm
var postgresNamesRegexp = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]{0,62}$`)

// See: https://github.com/standard-webhooks/standard-webhooks/blob/main/spec/standard-webhooks.md
// We use 4 * Math.ceil(n/3) to obtain unpadded length in base 64
// So this 4 * Math.ceil(24/3) = 32 and 4 * Math.ceil(64/3) = 88 for symmetric secrets
// Since Ed25519 key is 32 bytes so we have 4 * Math.ceil(32/3) = 44
var symmetricSecretFormat = regexp.MustCompile(`^v1,whsec_[A-Za-z0-9+/=]{32,88}`)
var asymmetricSecretFormat = regexp.MustCompile(`^v1a,whpk_[A-Za-z0-9+/=]{44,}:whsk_[A-Za-z0-9+/=]{44,}$`)

// Time is used to represent timestamps in the configuration, as envconfig has
// trouble parsing empty strings, due to time.Time.UnmarshalText().
type Time struct {
	time.Time
}

func (t *Time) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// OAuthProviderConfiguration holds all config related to external account providers.
type OAuthProviderConfiguration struct {
	ClientID      []string `json:"client_id" split_words:"true"`
	Secret        string   `json:"secret"`
	RedirectURI   string   `json:"redirect_uri" split_words:"true"`
	URL           string   `json:"url"`
	ApiURL        string   `json:"api_url" split_words:"true"`
	Enabled       bool     `json:"enabled"`
	EmailOptional bool     `json:"email_optional" split_words:"true"`
	// SkipNonceCheck bypasses nonce verification during OIDC token validation.
	// Note: Nonce verification helps prevent replay attacks; only disable when necessary.
	SkipNonceCheck bool `json:"skip_nonce_check" split_words:"true"`
}

// OAuthServerConfiguration holds OAuth server configuration
type OAuthServerConfiguration struct {
	Enabled                  bool          `json:"enabled" default:"false"`
	AllowDynamicRegistration bool          `json:"allow_dynamic_registration" split_words:"true"`
	AuthorizationPath        string        `json:"authorization_path" split_words:"true"`
	AuthorizationTTL         time.Duration `json:"authorization_ttl" split_words:"true" default:"10m"`
	// Placeholder for now, for (near) future extensibility
	DefaultScope string `json:"default_scope" split_words:"true" default:"email"`
}

type AnonymousProviderConfiguration struct {
	Enabled bool `json:"enabled" default:"false"`
}

// CustomOAuthConfiguration holds configuration for custom OAuth and OIDC providers
type CustomOAuthConfiguration struct {
	Enabled      bool `json:"enabled" split_words:"true" default:"true"`
	MaxProviders int  `json:"max_providers" split_words:"true" default:"0"`
}

type EmailProviderConfiguration struct {
	Enabled bool `json:"enabled" default:"true"`

	AuthorizedAddresses []string `json:"authorized_addresses" split_words:"true"`

	MagicLinkEnabled bool `json:"magic_link_enabled" default:"true" split_words:"true"`
}

type DBAdvisorConfiguration struct {
	Enabled             bool          `json:"enabled" default:"true"`
	SamplingInterval    time.Duration `json:"sampling_interval" split_words:"true" default:"200ms"`
	ObservationInterval time.Duration `json:"observation_interval" split_words:"true" default:"20s"`
}

// DBConfiguration holds all the database related configuration.
type DBConfiguration struct {
	Driver    string `json:"driver" required:"true"`
	URL       string `json:"url" envconfig:"DATABASE_URL" required:"true"`
	Namespace string `json:"namespace" envconfig:"DB_NAMESPACE" default:"auth"`

	// Percentage of DB conns the auth server may use in
	// integer form i.e.: [1, 100] -> [1%, 100%]
	ConnPercentage int `json:"conn_percentage" split_words:"true"`

	// MaxPoolSize defaults to 0 (unlimited).
	MaxPoolSize       int           `json:"max_pool_size" split_words:"true"`
	MaxIdlePoolSize   int           `json:"max_idle_pool_size" split_words:"true"`
	ConnMaxLifetime   time.Duration `json:"conn_max_lifetime,omitempty" split_words:"true"`
	ConnMaxIdleTime   time.Duration `json:"conn_max_idle_time,omitempty" split_words:"true"`
	HealthCheckPeriod time.Duration `json:"health_check_period" split_words:"true"`
	MigrationsPath    string        `json:"migrations_path" split_words:"true" default:"./migrations"`
	CleanupEnabled    bool          `json:"cleanup_enabled" split_words:"true" default:"false"`

	Advisor DBAdvisorConfiguration `json:"advisor"`
}

func (c *DBConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

// JWTConfiguration holds all the JWT related configuration.
type JWTConfiguration struct {
	Secret           string         `json:"secret" required:"true"`
	Exp              int            `json:"exp"`
	Aud              string         `json:"aud"`
	AdminGroupName   string         `json:"admin_group_name" split_words:"true"`
	AdminRoles       []string       `json:"admin_roles" split_words:"true"`
	DefaultGroupName string         `json:"default_group_name" split_words:"true"`
	Issuer           string         `json:"issuer"`
	KeyID            string         `json:"key_id" split_words:"true"`
	Keys             JwtKeysDecoder `json:"keys"`
	ValidMethods     []string       `json:"-" split_words:"true"`
}

type MFAFactorTypeConfiguration struct {
	EnrollEnabled bool `json:"enroll_enabled" split_words:"true" default:"false"`
	VerifyEnabled bool `json:"verify_enabled" split_words:"true" default:"false"`
}

type TOTPFactorTypeConfiguration struct {
	EnrollEnabled bool `json:"enroll_enabled" split_words:"true" default:"true"`
	VerifyEnabled bool `json:"verify_enabled" split_words:"true" default:"true"`
}

type PhoneFactorTypeConfiguration struct {
	// Default to false in order to ensure Phone MFA is opt-in
	MFAFactorTypeConfiguration
	OtpLength    int                `json:"otp_length" split_words:"true"`
	SMSTemplate  *template.Template `json:"-"`
	MaxFrequency time.Duration      `json:"max_frequency" split_words:"true"`
	Template     string             `json:"template"`
}

// MFAConfiguration holds all the MFA related Configuration
type MFAConfiguration struct {
	ChallengeExpiryDuration     float64                      `json:"challenge_expiry_duration" default:"300" split_words:"true"`
	FactorExpiryDuration        time.Duration                `json:"factor_expiry_duration" default:"300s" split_words:"true"`
	RateLimitChallengeAndVerify float64                      `split_words:"true" default:"15"`
	MaxEnrolledFactors          float64                      `split_words:"true" default:"10"`
	MaxVerifiedFactors          int                          `split_words:"true" default:"10"`
	Phone                       PhoneFactorTypeConfiguration `split_words:"true"`
	TOTP                        TOTPFactorTypeConfiguration  `split_words:"true"`
	WebAuthn                    MFAFactorTypeConfiguration   `split_words:"true"`
}

type WebAuthnConfiguration struct {
	RPID                    string        `json:"rp_id" envconfig:"RP_ID"`
	RPDisplayName           string        `json:"rp_display_name" split_words:"true"`
	RPOrigins               []string      `json:"rp_origins" split_words:"true"`
	ChallengeExpiryDuration time.Duration `json:"challenge_expiry_duration" split_words:"true" default:"5m"`
}

func (w *WebAuthnConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

type PasskeyConfiguration struct {
	Enabled            bool `json:"enabled" default:"false"`
	MaxPasskeysPerUser int  `json:"max_passkeys_per_user" split_words:"true" default:"10"`
}

type APIConfiguration struct {
	Host               string
	Port               string `envconfig:"PORT" default:"8081"`
	Endpoint           string
	RequestIDHeader    string        `envconfig:"REQUEST_ID_HEADER"`
	ExternalURL        string        `json:"external_url" envconfig:"API_EXTERNAL_URL" required:"true"`
	MaxRequestDuration time.Duration `json:"max_request_duration" split_words:"true" default:"10s"`
}

func (a *APIConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

type SessionsConfiguration struct {
	Timebox           *time.Duration `json:"timebox,omitempty"`
	InactivityTimeout *time.Duration `json:"inactivity_timeout,omitempty" split_words:"true"`
	AllowLowAAL       *time.Duration `json:"allow_low_aal,omitempty" split_words:"true"`

	SinglePerUser bool     `json:"single_per_user" split_words:"true"`
	Tags          []string `json:"tags,omitempty"`
}

func (c *SessionsConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

type PasswordRequiredCharacters []string

func (v *PasswordRequiredCharacters) Decode(value string) error {
	_ = "STUB: not implemented"
	return nil
}

// part ended in escape character, so it should be joined with the next one

// HIBPBloomConfiguration configures a bloom cache for pwned passwords. Use
// this tool to gauge the Items and FalsePositives values:
// https://hur.st/bloomfilter
type HIBPBloomConfiguration struct {
	Enabled        bool    `json:"enabled"`
	Items          uint    `json:"items" default:"100000"`
	FalsePositives float64 `json:"false_positives" split_words:"true" default:"0.0000099"`
}

type HIBPConfiguration struct {
	Enabled    bool `json:"enabled"`
	FailClosed bool `json:"fail_closed" split_words:"true"`

	UserAgent string `json:"user_agent" split_words:"true" default:"https://github.com/supabase/gotrue"`

	Bloom HIBPBloomConfiguration `json:"bloom"`
}

type PasswordConfiguration struct {
	MinLength int `json:"min_length" split_words:"true"`

	RequiredCharacters PasswordRequiredCharacters `json:"required_characters" split_words:"true"`

	HIBP HIBPConfiguration `json:"hibp"`
}

type AuditLogConfiguration struct {
	DisablePostgres bool `split_words:"true" default:"false"`
}

type ExperimentalConfiguration struct {
	// Names of providers (e.g. "google") which have their own identity
	// linking domain, meaning that the ones listed here _will not
	// participate_ in email similarity linking with other accounts.
	ProvidersWithOwnLinkingDomain []string `split_words:"true"`
}

// ReloadingConfiguration holds the configuration values for runtime
// configuration reloads. These are startup configuration values meaning
// they do not react to live config reloads.
//
// IMPORTANT:
// * You must provide the --config-dir flag for these settings to take effect.
// * These config values are for startup, they remain static through reloads.
type ReloadingConfiguration struct {

	// If notify reloading is enabled the auth server will attempt to use the
	// filesystems notification support to watch for config updates.
	NotifyEnabled bool `json:"notify_enabled" split_words:"true" default:"true"`

	// When notify reloading fails, fallback to filesystem polling if this
	// setting is enabled.
	PollerEnabled bool `json:"poller_enabled" split_words:"false" default:"false"`

	// This determines how often to poll the filesystem when notify is disabled.
	PollerInterval time.Duration `json:"poller_interval" split_words:"true" default:"10s"`

	// If signal reloading is enabled the auth server will listen for the
	// given SignalNumber and reload the config when received. This may be
	// used to configure `systemd reload` support, by default the SIGUSR1 linux
	// signal number of 10 is used.
	SignalEnabled bool `json:"signal_enabled" split_words:"true" default:"false"`
	SignalNumber  int  `json:"signal_number" split_words:"true" default:"10"`

	// When at least one reloader is enabled this flag determines how much idle
	// time must pass before triggering a reload. This ensures a single
	// auth server config reload operation during a burst of config updates.
	GracePeriodInterval time.Duration `json:"grace_period_interval" split_words:"true" default:"5s"`
}

// GlobalConfiguration holds all the configuration that applies to all instances.
type GlobalConfiguration struct {
	API           APIConfiguration
	DB            DBConfiguration
	External      ProviderConfiguration
	CustomOAuth   CustomOAuthConfiguration `envconfig:"CUSTOM_OAUTH"`
	OAuthServer   OAuthServerConfiguration `envconfig:"OAUTH_SERVER"`
	Logging       LoggingConfig            `envconfig:"LOG"`
	Profiler      ProfilerConfig           `envconfig:"PROFILER"`
	OperatorToken string                   `split_words:"true" required:"false"`
	Tracing       TracingConfig
	Metrics       MetricsConfig
	SMTP          SMTPConfiguration
	AuditLog      AuditLogConfiguration `split_words:"true"`

	RateLimitHeader                     string  `split_words:"true"`
	RateLimitEmailSent                  Rate    `split_words:"true" default:"30"`
	RateLimitSmsSent                    Rate    `split_words:"true" default:"30"`
	RateLimitVerify                     float64 `split_words:"true" default:"30"`
	RateLimitTokenRefresh               float64 `split_words:"true" default:"150"`
	RateLimitSso                        float64 `split_words:"true" default:"30"`
	RateLimitAnonymousUsers             float64 `split_words:"true" default:"30"`
	RateLimitOtp                        float64 `split_words:"true" default:"30"`
	RateLimitWeb3                       float64 `split_words:"true" default:"30"`
	RateLimitPasskey                    float64 `split_words:"true" default:"30"`
	RateLimitOAuthDynamicClientRegister float64 `split_words:"true" default:"10"`

	SiteURL         string   `json:"site_url" split_words:"true" required:"true"`
	URIAllowList    []string `json:"uri_allow_list" split_words:"true"`
	URIAllowListMap map[string]glob.Glob
	Password        PasswordConfiguration    `json:"password"`
	JWT             JWTConfiguration         `json:"jwt"`
	Mailer          MailerConfiguration      `json:"mailer"`
	Sms             SmsProviderConfiguration `json:"sms"`
	DisableSignup   bool                     `json:"disable_signup" split_words:"true"`
	Hook            HookConfiguration        `json:"hook" split_words:"true"`
	Security        SecurityConfiguration    `json:"security"`
	Sessions        SessionsConfiguration    `json:"sessions"`
	MFA             MFAConfiguration         `json:"MFA"`
	SAML            SAMLConfiguration        `json:"saml"`
	WebAuthn        WebAuthnConfiguration    `json:"webauthn"`
	Passkey         PasskeyConfiguration     `json:"passkey"`
	CORS            CORSConfiguration        `json:"cors"`
	IndexWorker     IndexWorkerConfiguration `json:"index_worker" split_words:"true"`

	Experimental ExperimentalConfiguration `json:"experimental"`
	Reloading    ReloadingConfiguration    `json:"reloading"`
}

type CORSConfiguration struct {
	AllowedHeaders []string `json:"allowed_headers" split_words:"true"`
}

func (c *CORSConfiguration) AllAllowedHeaders(defaults []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// EmailContentConfiguration holds the configuration for emails, both subjects and template URLs.
type EmailContentConfiguration struct {
	Invite           string `json:"invite"`
	Confirmation     string `json:"confirmation"`
	Recovery         string `json:"recovery"`
	EmailChange      string `json:"email_change" split_words:"true"`
	MagicLink        string `json:"magic_link" split_words:"true"`
	Reauthentication string `json:"reauthentication"`

	// Account Changes Notifications
	PasswordChangedNotification     string `json:"password_changed_notification" split_words:"true"`
	EmailChangedNotification        string `json:"email_changed_notification" split_words:"true"`
	PhoneChangedNotification        string `json:"phone_changed_notification" split_words:"true"`
	IdentityLinkedNotification      string `json:"identity_linked_notification" split_words:"true"`
	IdentityUnlinkedNotification    string `json:"identity_unlinked_notification" split_words:"true"`
	MFAFactorEnrolledNotification   string `json:"mfa_factor_enrolled_notification" split_words:"true"`
	MFAFactorUnenrolledNotification string `json:"mfa_factor_unenrolled_notification" split_words:"true"`
}

// NotificationsConfiguration holds the configuration for notification email states to indicate whether they are enabled or disabled.
type NotificationsConfiguration struct {
	PasswordChangedEnabled     bool `json:"password_changed_enabled" split_words:"true" default:"false"`
	EmailChangedEnabled        bool `json:"email_changed_enabled" split_words:"true" default:"false"`
	PhoneChangedEnabled        bool `json:"phone_changed_enabled" split_words:"true" default:"false"`
	IdentityLinkedEnabled      bool `json:"identity_linked_enabled" split_words:"true" default:"false"`
	IdentityUnlinkedEnabled    bool `json:"identity_unlinked_enabled" split_words:"true" default:"false"`
	MFAFactorEnrolledEnabled   bool `json:"mfa_factor_enrolled_enabled" split_words:"true" default:"false"`
	MFAFactorUnenrolledEnabled bool `json:"mfa_factor_unenrolled_enabled" split_words:"true" default:"false"`
}

type ProviderConfiguration struct {
	AnonymousUsers          AnonymousProviderConfiguration `json:"anonymous_users" split_words:"true"`
	Apple                   OAuthProviderConfiguration     `json:"apple"`
	Azure                   OAuthProviderConfiguration     `json:"azure"`
	Bitbucket               OAuthProviderConfiguration     `json:"bitbucket"`
	Discord                 OAuthProviderConfiguration     `json:"discord"`
	Facebook                OAuthProviderConfiguration     `json:"facebook"`
	Snapchat                OAuthProviderConfiguration     `json:"snapchat"`
	Figma                   OAuthProviderConfiguration     `json:"figma"`
	Fly                     OAuthProviderConfiguration     `json:"fly"`
	Github                  OAuthProviderConfiguration     `json:"github"`
	Gitlab                  OAuthProviderConfiguration     `json:"gitlab"`
	Google                  OAuthProviderConfiguration     `json:"google"`
	Kakao                   OAuthProviderConfiguration     `json:"kakao"`
	Notion                  OAuthProviderConfiguration     `json:"notion"`
	Keycloak                OAuthProviderConfiguration     `json:"keycloak"`
	Linkedin                OAuthProviderConfiguration     `json:"linkedin"`
	LinkedinOIDC            OAuthProviderConfiguration     `json:"linkedin_oidc" envconfig:"LINKEDIN_OIDC"`
	Spotify                 OAuthProviderConfiguration     `json:"spotify"`
	Slack                   OAuthProviderConfiguration     `json:"slack"`
	SlackOIDC               OAuthProviderConfiguration     `json:"slack_oidc" envconfig:"SLACK_OIDC"`
	Twitter                 OAuthProviderConfiguration     `json:"twitter"`
	Twitch                  OAuthProviderConfiguration     `json:"twitch"`
	VercelMarketplace       OAuthProviderConfiguration     `json:"vercel_marketplace" split_words:"true"`
	WorkOS                  OAuthProviderConfiguration     `json:"workos"`
	Email                   EmailProviderConfiguration     `json:"email"`
	Phone                   PhoneProviderConfiguration     `json:"phone"`
	X                       OAuthProviderConfiguration     `json:"x" envconfig:"X"`
	Zoom                    OAuthProviderConfiguration     `json:"zoom"`
	IosBundleId             string                         `json:"ios_bundle_id" split_words:"true"`
	RedirectURL             string                         `json:"redirect_url"`
	AllowedIdTokenIssuers   []string                       `json:"allowed_id_token_issuers" split_words:"true"`
	FlowStateExpiryDuration time.Duration                  `json:"flow_state_expiry_duration" split_words:"true"`

	// OIDCProviderCacheTTL controls how long OIDC discovery documents are cached.
	OIDCProviderCacheTTL time.Duration `json:"oidc_provider_cache_ttl" split_words:"true" default:"1h"`

	Web3Solana   SolanaConfiguration   `json:"web3_solana" split_words:"true"`
	Web3Ethereum EthereumConfiguration `json:"web3_ethereum" split_words:"true"`
}

type SolanaConfiguration struct {
	Enabled                 bool          `json:"enabled,omitempty" split_words:"true"`
	MaximumValidityDuration time.Duration `json:"maximum_validity_duration,omitempty" default:"10m" split_words:"true"`
}

type EthereumConfiguration struct {
	Enabled                 bool          `json:"enabled,omitempty" split_words:"true"`
	MaximumValidityDuration time.Duration `json:"maximum_validity_duration,omitempty" default:"10m" split_words:"true"`
}

type SMTPConfiguration struct {
	MaxFrequency   time.Duration `json:"max_frequency" split_words:"true"`
	Host           string        `json:"host"`
	Port           int           `json:"port,omitempty" default:"587"`
	User           string        `json:"user"`
	Pass           string        `json:"pass,omitempty"`
	AdminEmail     string        `json:"admin_email" split_words:"true"`
	SenderName     string        `json:"sender_name" split_words:"true"`
	Headers        string        `json:"headers"`
	LoggingEnabled bool          `json:"logging_enabled" split_words:"true" default:"false"`

	fromAddress       string              `json:"-"`
	normalizedHeaders map[string][]string `json:"-"`
}

func (c *SMTPConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *SMTPConfiguration) FromAddress() string { _ = "STUB: not implemented"; return "" }

func (c *SMTPConfiguration) NormalizedHeaders() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

type MailerConfiguration struct {
	Autoconfirm                 bool `json:"autoconfirm"`
	AllowUnverifiedEmailSignIns bool `json:"allow_unverified_email_sign_ins" split_words:"true" default:"false"`

	Subjects      EmailContentConfiguration  `json:"subjects"`
	Templates     EmailContentConfiguration  `json:"templates"`
	URLPaths      EmailContentConfiguration  `json:"url_paths"`
	Notifications NotificationsConfiguration `json:"notifications" split_words:"true"`

	SecureEmailChangeEnabled bool `json:"secure_email_change_enabled" split_words:"true" default:"true"`

	OtpExp    uint `json:"otp_exp" split_words:"true"`
	OtpLength int  `json:"otp_length" split_words:"true"`

	ExternalHosts []string `json:"external_hosts" split_words:"true"`

	// EXPERIMENTAL: All config below here may be removed in a future release.
	EmailBackgroundSending        bool   `json:"email_background_sending" split_words:"true" default:"false"`
	EmailValidationExtended       bool   `json:"email_validation_extended" split_words:"true" default:"false"`
	EmailValidationServiceURL     string `json:"email_validation_service_url" split_words:"true"`
	EmailValidationServiceHeaders string `json:"email_validation_service_headers" split_words:"true"`
	EmailValidationBlockedMX      string `json:"email_validation_blocked_mx" split_words:"true"`

	// Max size in bytes we will read from a template endpoint
	TemplateMaxSize int `json:"template_max_size" split_words:"true" default:"1000000"`

	// The maximum age of a template before we consider it stale.
	TemplateMaxAge time.Duration `json:"template_max_age" split_words:"true" default:"10m"`

	// The time between retrying a failed template reload.
	TemplateRetryInterval time.Duration `json:"template_retry_interval" split_words:"true" default:"10s"`

	// If true enable background reloading of templates to avoid blocking
	// IO in requests.
	TemplateReloadingEnabled bool `json:"template_reloading_enabled" split_words:"true" default:"false"`

	// The maximum time a server may be idle before template reloading stops.
	// Note that even when the server is idle, a config reload will trigger a
	// template reload.
	TemplateReloadingMaxIdle time.Duration `json:"template_reloading_max_idle" split_words:"true" default:"20m"`

	serviceHeaders   map[string][]string `json:"-"`
	blockedMXRecords map[string]bool     `json:"-"`
}

func (c *MailerConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

// EmailValidationBlockedMX is a JSON array in the config string for brevity.

func (c *MailerConfiguration) GetEmailValidationServiceHeaders() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *MailerConfiguration) GetEmailValidationBlockedMXRecords() map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

type PhoneProviderConfiguration struct {
	Enabled bool `json:"enabled" default:"false"`
}

type SmsProviderConfiguration struct {
	Autoconfirm       bool               `json:"autoconfirm"`
	MaxFrequency      time.Duration      `json:"max_frequency" split_words:"true"`
	OtpExp            uint               `json:"otp_exp" split_words:"true"`
	OtpLength         int                `json:"otp_length" split_words:"true"`
	Provider          string             `json:"provider"`
	Template          string             `json:"template"`
	TestOTP           map[string]string  `json:"test_otp" split_words:"true"`
	TestOTPValidUntil Time               `json:"test_otp_valid_until" split_words:"true"`
	SMSTemplate       *template.Template `json:"-"`

	Twilio       TwilioProviderConfiguration       `json:"twilio"`
	TwilioVerify TwilioVerifyProviderConfiguration `json:"twilio_verify" split_words:"true"`
	Messagebird  MessagebirdProviderConfiguration  `json:"messagebird"`
	Textlocal    TextlocalProviderConfiguration    `json:"textlocal"`
	Vonage       VonageProviderConfiguration       `json:"vonage"`
}

func (c *SmsProviderConfiguration) GetTestOTP(phone string, now time.Time) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type TwilioProviderConfiguration struct {
	AccountSid        string `json:"account_sid" split_words:"true"`
	AuthToken         string `json:"auth_token" split_words:"true"`
	MessageServiceSid string `json:"message_service_sid" split_words:"true"`
	ContentSid        string `json:"content_sid" split_words:"true"`
}

type TwilioVerifyProviderConfiguration struct {
	AccountSid        string `json:"account_sid" split_words:"true"`
	AuthToken         string `json:"auth_token" split_words:"true"`
	MessageServiceSid string `json:"message_service_sid" split_words:"true"`
}

type MessagebirdProviderConfiguration struct {
	AccessKey  string `json:"access_key" split_words:"true"`
	Originator string `json:"originator" split_words:"true"`
}

type TextlocalProviderConfiguration struct {
	ApiKey string `json:"api_key" split_words:"true"`
	Sender string `json:"sender" split_words:"true"`
}

type VonageProviderConfiguration struct {
	ApiKey    string `json:"api_key" split_words:"true"`
	ApiSecret string `json:"api_secret" split_words:"true"`
	From      string `json:"from" split_words:"true"`
}

type CaptchaConfiguration struct {
	Enabled  bool          `json:"enabled" default:"false"`
	Provider string        `json:"provider" default:"hcaptcha"`
	Secret   string        `json:"provider_secret"`
	Timeout  time.Duration `json:"timeout" split_words:"true" default:"10s"`
}

func (c *CaptchaConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

// DatabaseEncryptionConfiguration configures Auth to encrypt certain columns.
// Once Encrypt is set to true, data will start getting encrypted with the
// provided encryption key. Setting it to false just stops encryption from
// going on further, but DecryptionKeys would have to contain the same key so
// the encrypted data remains accessible.
type DatabaseEncryptionConfiguration struct {
	Encrypt bool `json:"encrypt"`

	EncryptionKeyID string `json:"encryption_key_id" split_words:"true"`
	EncryptionKey   string `json:"-" split_words:"true"`

	DecryptionKeys map[string]string `json:"-" split_words:"true"`
}

func (c *DatabaseEncryptionConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

type SecurityConfiguration struct {
	Captcha                               CaptchaConfiguration `json:"captcha"`
	RefreshTokenUpgradePercentage         int                  `json:"refresh_token_upgrade_percentage" split_words:"true"`
	RefreshTokenAlgorithmVersion          int                  `json:"refresh_token_algorithm_version" split_words:"true"`
	RefreshTokenRotationEnabled           bool                 `json:"refresh_token_rotation_enabled" split_words:"true" default:"true"`
	RefreshTokenReuseInterval             int                  `json:"refresh_token_reuse_interval" split_words:"true"`
	RefreshTokenAllowReuse                bool                 `json:"refresh_token_allow_reuse" split_words:"true"`
	UpdatePasswordRequireReauthentication bool                 `json:"update_password_require_reauthentication" split_words:"true"`
	UpdatePasswordRequireCurrentPassword  bool                 `json:"update_password_require_current_password" split_words:"true"`
	ManualLinkingEnabled                  bool                 `json:"manual_linking_enabled" split_words:"true" default:"false"`
	SbForwardedForEnabled                 bool                 `json:"sb_forwarded_for_enabled" split_words:"true" default:"false"`

	DBEncryption DatabaseEncryptionConfiguration `json:"database_encryption" split_words:"true"`
}

func (c *SecurityConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func loadEnvironment(filename string) error { _ = "STUB: not implemented"; return nil }

// handle if .env file does not exist, this is OK

// Moving away from the existing HookConfig so we can get a fresh start.
type HookConfiguration struct {
	MFAVerificationAttempt      ExtensibilityPointConfiguration `json:"mfa_verification_attempt" split_words:"true"`
	PasswordVerificationAttempt ExtensibilityPointConfiguration `json:"password_verification_attempt" split_words:"true"`
	CustomAccessToken           ExtensibilityPointConfiguration `json:"custom_access_token" split_words:"true"`
	SendEmail                   ExtensibilityPointConfiguration `json:"send_email" split_words:"true"`
	SendSMS                     ExtensibilityPointConfiguration `json:"send_sms" split_words:"true"`

	BeforeUserCreated ExtensibilityPointConfiguration `json:"before_user_created" split_words:"true"`
	AfterUserCreated  ExtensibilityPointConfiguration `json:"after_user_created" split_words:"true"`
}

type HTTPHookSecrets []string

func (h *HTTPHookSecrets) Decode(value string) error { _ = "STUB: not implemented"; return nil }

type ExtensibilityPointConfiguration struct {
	URI     string `json:"uri"`
	Enabled bool   `json:"enabled"`
	// For internal use together with Postgres Hook. Not publicly exposed.
	HookName string `json:"-"`
	// We use | as a separator for keys and : as a separator for keys within a keypair. For instance: v1,whsec_test|v1a,whpk_myother:v1a,whsk_testkey|v1,whsec_secret3
	HTTPHookSecrets HTTPHookSecrets `json:"secrets" envconfig:"secrets"`
}

func (h *HookConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (e *ExtensibilityPointConfiguration) ValidateExtensibilityPoint() error {
	_ = "STUB: not implemented"
	return nil
}

func validatePostgresPath(u *url.URL) error { _ = "STUB: not implemented"; return nil }

// Validate schema and table names

func isValidSecretFormat(secret string) bool { _ = "STUB: not implemented"; return false }

func validateHTTPHookSecrets(secrets []string) error { _ = "STUB: not implemented"; return nil }

func (e *ExtensibilityPointConfiguration) PopulateExtensibilityPoint() error {
	_ = "STUB: not implemented"
	return nil
}

// LoadFile calls godotenv.Load() when the given filename is empty ignoring any
// errors loading, otherwise it calls godotenv.Overload(filename).
//
// godotenv.Load: preserves env, ".env" path is optional
// godotenv.Overload: overrides env, "filename" path must exist
func LoadFile(filename string) error { _ = "STUB: not implemented"; return nil }

// handle if .env file does not exist, this is OK

// LoadDirectory does nothing when configDir is empty, otherwise it will attempt
// to load a list of configuration files located in configDir by using ReadDir
// to obtain a sorted list of files containing a .env suffix.
//
// When the list is empty it will do nothing, otherwise it passes the file list
// to godotenv.Overload to pull them into the current environment.
func LoadDirectory(configDir string) error { _ = "STUB: not implemented"; return nil }

// Returns entries sorted by filename

// We mimic the behavior of LoadGlobal here, if an explicit path is
// provided we return an error.

// ignore directories

// We only read files ending in .env

// ent.Name() does not include the watch dir.

// If at least one path was found we load the configuration files in the
// directory. We don't call override without config files because it will
// override the env vars previously set with a ".env", if one exists.

func loadDirectoryPaths(p ...string) error {
	_ = "STUB: not implemented"
	// If at least one path was found we load the configuration files in the
	// directory. We don't call override without config files because it will
	// override the env vars previously set with a ".env", if one exists.
	return nil
}

// LoadGlobalFromEnv will return a new *GlobalConfiguration value from the
// currently configured environment.
func LoadGlobalFromEnv() (*GlobalConfiguration, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadGlobal(filename string) (*GlobalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadGlobal(config *GlobalConfiguration) error {
	_ = "STUB: not implemented"
	// although the package is called "auth" it used to be called "gotrue"
	// so environment configs will remain to be called "GOTRUE"
	return nil
}

func populateGlobal(config *GlobalConfiguration) error { _ = "STUB: not implemented"; return nil }

// ApplyDefaults sets defaults for a GlobalConfiguration
func (config *GlobalConfiguration) ApplyDefaults() error { _ = "STUB: not implemented"; return nil }

// transform the secret into a JWK for consistency

// 1 day

// 6-digit otp by default

// 6-digit otp by default

// 6-digit otp by default

func (config *GlobalConfiguration) applyDefaultsJWT(secret []byte) error {
	_ = "STUB: not implemented"
	// transform the secret into a JWK for consistency
	return nil
}

func (config *GlobalConfiguration) applyDefaultsJWTPrivateKey(privKey jwk.Key) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate validates all of configuration.
func (c *GlobalConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (o *OAuthProviderConfiguration) ValidateOAuth() error { _ = "STUB: not implemented"; return nil }

func (t *TwilioProviderConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *TwilioVerifyProviderConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *MessagebirdProviderConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *TextlocalProviderConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *VonageProviderConfiguration) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *SmsProviderConfiguration) IsTwilioVerifyProvider() bool {
	_ = "STUB: not implemented"
	return false
}

// IndexWorkerConfiguration holds the configuration for creating database indexes on the users table.
type IndexWorkerConfiguration struct {
	// user opt-in — when true, always create indexes (threshold is ignored).
	EnsureUserSearchIndexesExist bool `json:"ensure_user_search_indexes_exist" split_words:"true" default:"false"`
	// progressive rollout — when > 0, create indexes only if user count ≤ threshold.
	// A value of 0 means disabled. Has no effect when EnsureUserSearchIndexesExist is true.
	MaxUsersThreshold int64 `json:"max_users_threshold" split_words:"true" default:"0"`
}
