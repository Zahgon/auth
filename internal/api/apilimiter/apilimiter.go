package apilimiter

import (
	"maps"
	"slices"
	"time"

	"github.com/didip/tollbooth/v5/limiter"
	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/ratelimit"
)

const (
	// GOTRUE_RATE_LIMIT_EMAIL_SENT
	//   -> RateLimitEmailSent
	envRateLimitEmailSent = "GOTRUE_RATE_LIMIT_EMAIL_SENT"
	fieldEmail            = "Email"

	// GOTRUE_RATE_LIMIT_SMS_SENT
	//   -> RateLimitSmsSent
	envRateLimitSmsSent = "GOTRUE_RATE_LIMIT_SMS_SENT"
	fieldPhone          = "Phone"

	// GOTRUE_RATE_LIMIT_ANONYMOUS_USERS
	//  -> RateLimitAnonymousUsers
	envRateLimitAnonymousUsers = "GOTRUE_RATE_LIMIT_ANONYMOUS_USERS"
	fieldAnonymousSignIns      = "AnonymousSignIns"

	// GOTRUE_MFA_RATE_LIMIT_CHALLENGE_AND_VERIFY
	//   -> MFA.RateLimitChallengeAndVerify
	envMFARateLimitChallengeAndVerify = "GOTRUE_MFA_RATE_LIMIT_CHALLENGE_AND_VERIFY"
	fieldFactorChallenge              = "FactorChallenge"
	fieldFactorVerify                 = "FactorVerify"

	// GOTRUE_RATE_LIMIT_OTP
	//   -> RateLimitOtp
	envRateLimitOtp = "GOTRUE_RATE_LIMIT_OTP"
	fieldMagicLink  = "MagicLink"
	fieldOtp        = "Otp"
	fieldRecover    = "Recover"
	fieldResend     = "Resend"
	fieldSignups    = "Signups"
	fieldUser       = "User"

	// GOTRUE_RATE_LIMIT_OAUTH_DYNAMIC_CLIENT_REGISTER
	//   -> RateLimitOAuthDynamicClientRegister
	envRateLimitOAuthDynamicClientRegister = "GOTRUE_RATE_LIMIT_OAUTH_DYNAMIC_CLIENT_REGISTER"
	fieldOAuthClientRegister               = "OAuthClientRegister"

	// GOTRUE_RATE_LIMIT_PASSKEY
	//   -> RateLimitPasskey
	envRateLimitPasskey        = "GOTRUE_RATE_LIMIT_PASSKEY" // #nosec G101
	fieldPasskeyAuthentication = "PasskeyAuthentication"

	// GOTRUE_SAML_RATE_LIMIT_ASSERTION
	//   -> SAML.RateLimitAssertion
	envSAMLRateLimitAssertion = "GOTRUE_SAML_RATE_LIMIT_ASSERTION"
	fieldSAMLAssertion        = "SAMLAssertion"

	// GOTRUE_RATE_LIMIT_SSO
	//   -> RateLimitSso
	envRateLimitSso = "GOTRUE_RATE_LIMIT_SSO"
	fieldSSO        = "SSO"

	// GOTRUE_RATE_LIMIT_TOKEN_REFRESH
	//   -> RateLimitTokenRefresh
	envRateLimitTokenRefresh = "GOTRUE_RATE_LIMIT_TOKEN_REFRESH" // #nosec G101
	fieldToken               = "Token"

	// GOTRUE_RATE_LIMIT_VERIFY
	//   -> RateLimitVerify
	envRateLimitVerify = "GOTRUE_RATE_LIMIT_VERIFY"
	fieldVerify        = "Verify"

	// GOTRUE_RATE_LIMIT_WEB3
	//   -> RateLimitWeb3
	envRateLimitWeb3 = "GOTRUE_RATE_LIMIT_WEB3"
	fieldWeb3        = "Web3"
)

var ratelimitFieldsToEnv = map[string]string{
	fieldEmail: envRateLimitEmailSent,
	fieldPhone: envRateLimitSmsSent,
}

var tollboothFieldsToEnv = map[string]string{
	fieldAnonymousSignIns:      envRateLimitAnonymousUsers,
	fieldFactorChallenge:       envMFARateLimitChallengeAndVerify,
	fieldFactorVerify:          envMFARateLimitChallengeAndVerify,
	fieldMagicLink:             envRateLimitOtp,
	fieldOtp:                   envRateLimitOtp,
	fieldRecover:               envRateLimitOtp,
	fieldResend:                envRateLimitOtp,
	fieldSignups:               envRateLimitOtp,
	fieldUser:                  envRateLimitOtp,
	fieldOAuthClientRegister:   envRateLimitOAuthDynamicClientRegister,
	fieldPasskeyAuthentication: envRateLimitPasskey,
	fieldSAMLAssertion:         envSAMLRateLimitAssertion,
	fieldSSO:                   envRateLimitSso,
	fieldToken:                 envRateLimitTokenRefresh,
	fieldVerify:                envRateLimitVerify,
	fieldWeb3:                  envRateLimitWeb3,
}

var fieldsToEnv = func() map[string]string {
	n := len(ratelimitFieldsToEnv) + len(tollboothFieldsToEnv)
	out := make(map[string]string, n)
	maps.Insert(out, maps.All(ratelimitFieldsToEnv))
	maps.Insert(out, maps.All(tollboothFieldsToEnv))
	return out
}()

var envsToFields = func() map[string][]string {
	out := make(map[string][]string)
	for field, env := range fieldsToEnv {
		out[env] = append(out[env], field)
	}
	for _, fields := range out {
		slices.Sort(fields)
	}
	return out
}()

type Limiter struct {
	cfg *conf.GlobalConfiguration

	// GOTRUE_RATE_LIMIT_EMAIL_SENT
	//   -> RateLimitEmailSent
	Email ratelimit.Limiter

	// GOTRUE_RATE_LIMIT_SMS_SENT
	//   -> RateLimitSmsSent
	Phone ratelimit.Limiter

	// GOTRUE_RATE_LIMIT_ANONYMOUS_USERS
	//  -> RateLimitAnonymousUsers
	AnonymousSignIns *limiter.Limiter

	// GOTRUE_MFA_RATE_LIMIT_CHALLENGE_AND_VERIFY
	//   -> MFA.RateLimitChallengeAndVerify
	FactorChallenge *limiter.Limiter
	FactorVerify    *limiter.Limiter

	// GOTRUE_RATE_LIMIT_OTP
	//   -> RateLimitOtp
	MagicLink *limiter.Limiter
	Otp       *limiter.Limiter
	Recover   *limiter.Limiter
	Resend    *limiter.Limiter
	Signups   *limiter.Limiter
	User      *limiter.Limiter

	// GOTRUE_RATE_LIMIT_OAUTH_DYNAMIC_CLIENT_REGISTER
	//   -> RateLimitOAuthDynamicClientRegister
	OAuthClientRegister *limiter.Limiter

	// GOTRUE_RATE_LIMIT_PASSKEY
	//   -> RateLimitPasskey
	PasskeyAuthentication *limiter.Limiter

	// GOTRUE_SAML_RATE_LIMIT_ASSERTION
	//   -> SAML.RateLimitAssertion
	SAMLAssertion *limiter.Limiter

	// GOTRUE_RATE_LIMIT_SSO
	//   -> RateLimitSso
	SSO *limiter.Limiter

	// GOTRUE_RATE_LIMIT_TOKEN_REFRESH
	//   -> RateLimitTokenRefresh
	Token *limiter.Limiter

	// GOTRUE_RATE_LIMIT_VERIFY
	//   -> RateLimitVerify
	Verify *limiter.Limiter

	// GOTRUE_RATE_LIMIT_WEB3
	//   -> RateLimitWeb3
	Web3 *limiter.Limiter
}

func New(gc *conf.GlobalConfiguration) *Limiter { _ = "STUB: not implemented"; return nil }

// These all use the OTP limit per 5 min with 1hour ttl and burst of 30.

func (o *Limiter) Copy() *Limiter { _ = "STUB: not implemented"; return nil }

func (o *Limiter) Update(
	le *logrus.Entry,
	nextCfg *conf.GlobalConfiguration,
) *Limiter {
	_ = "STUB: not implemented"
	return nil
}

func newTollbooth(freq float64, burst int, ttl time.Duration) *limiter.Limiter {
	_ = "STUB: not implemented"
	return nil
}

func newLimiterPer5mOver1h(rate float64) *limiter.Limiter { _ = "STUB: not implemented"; return nil }

func logEnvUpdates(
	le *logrus.Entry,
	env string,
	prevVal, nextVal any,
) {
	_ = "STUB: not implemented"
	return
}

func logUpdate(
	le *logrus.Entry,
	field string,
	prevVal, nextVal any,
) {
	_ = "STUB: not implemented"
	return
}
