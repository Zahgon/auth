package apierrors

// OAuthError is the JSON handler for OAuth2 error responses
type OAuthError struct {
	Err             string `json:"error"`
	Description     string `json:"error_description,omitempty"`
	InternalError   error  `json:"-"`
	InternalMessage string `json:"-"`
}

func NewOAuthError(err string, description string) *OAuthError {
	_ = "STUB: not implemented"
	return nil
}

func (e *OAuthError) Error() string { _ = "STUB: not implemented"; return "" }

// WithInternalError adds internal error information to the error
func (e *OAuthError) WithInternalError(err error) *OAuthError {
	_ = "STUB: not implemented"
	return nil
}

// WithInternalMessage adds internal message information to the error
func (e *OAuthError) WithInternalMessage(fmtString string, args ...any) *OAuthError {
	_ = "STUB: not implemented"
	return nil
}

// Cause returns the root cause error
func (e *OAuthError) Cause() error { _ = "STUB: not implemented"; return nil }

// HTTPError is an error with a message and an HTTP status code.
type HTTPError struct {
	HTTPStatus      int    `json:"code"`                 // do not rename the JSON tags!
	ErrorCode       string `json:"error_code,omitempty"` // do not rename the JSON tags!
	Message         string `json:"msg"`                  // do not rename the JSON tags!
	InternalError   error  `json:"-"`
	InternalMessage string `json:"-"`
	ErrorID         string `json:"error_id,omitempty"`
}

func NewHTTPError(httpStatus int, errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewBadRequestError(errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewNotFoundError(errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewForbiddenError(errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewUnprocessableEntityError(errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewTooManyRequestsError(errorCode ErrorCode, fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewInternalServerError(fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewConflictError(fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func (e *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *HTTPError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// Cause returns the root cause error
func (e *HTTPError) Cause() error { _ = "STUB: not implemented"; return nil }

// WithInternalError adds internal error information to the error
func (e *HTTPError) WithInternalError(err error) *HTTPError { _ = "STUB: not implemented"; return nil }

// WithInternalMessage adds internal message information to the error
func (e *HTTPError) WithInternalMessage(fmtString string, args ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}
