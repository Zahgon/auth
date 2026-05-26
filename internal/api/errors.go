package api

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/supabase/auth/internal/api/apierrors"
)

// Common error messages during signup flow
const (
	DuplicateEmailMsg = "A user with this email address has already been registered"
	DuplicatePhoneMsg = "A user with this phone number has already been registered"
)

var (
	UserExistsError error = errors.New("user already exists")
)

const InvalidChannelError = "Invalid channel, supported values are 'sms' or 'whatsapp'. 'whatsapp' is only supported if Twilio or Twilio Verify is used as the provider."

var oauthErrorMap = map[int]string{
	http.StatusBadRequest:          "invalid_request",
	http.StatusUnauthorized:        "unauthorized_client",
	http.StatusForbidden:           "access_denied",
	http.StatusInternalServerError: "server_error",
	http.StatusServiceUnavailable:  "temporarily_unavailable",
}

// Type aliases while we slowly refactor api errors.
type (
	HTTPError  = apierrors.HTTPError
	OAuthError = apierrors.OAuthError
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
func recoverer(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ErrorCause is an error interface that contains the method Cause() for returning root cause errors
type ErrorCause interface {
	Cause() error
}

type HTTPErrorResponse20240101 struct {
	Code    apierrors.ErrorCode `json:"code"`
	Message string              `json:"message"`
}

func HandleResponseError(err error, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Echo back the determined API version from the request

// this will get us the stack trace too

// Provide better error messages for certain user-triggered Postgres errors.

func generateFrequencyLimitErrorMessage(timeStamp *time.Time, maxFrequency time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}
