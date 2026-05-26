package provider

type HTTPError struct {
	Code            int    `json:"code"`
	Message         string `json:"msg"`
	InternalError   error  `json:"-"`
	InternalMessage string `json:"-"`
	ErrorID         string `json:"error_id,omitempty"`
}

func (e *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *HTTPError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// Cause returns the root cause error
func (e *HTTPError) Cause() error { _ = "STUB: not implemented"; return nil }

// WithInternalError adds internal error information to the error
func (e *HTTPError) WithInternalError(err error) *HTTPError { _ = "STUB: not implemented"; return nil }

// WithInternalMessage adds internal message information to the error
func (e *HTTPError) WithInternalMessage(fmtString string, args ...interface{}) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func httpError(code int, fmtString string, args ...interface{}) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}
