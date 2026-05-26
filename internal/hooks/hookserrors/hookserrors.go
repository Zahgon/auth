// Package hookserrors holds the Error type and some functions to Check
// responses for errors.
package hookserrors

// Error is the type propagated by hook endpoints to communicate failure.
type Error struct {
	HTTPCode int    `json:"http_code,omitempty"`
	Message  string `json:"message,omitempty"`
}

// Error implements the error interface by returning e.Message.
func (e *Error) Error() string {
	_ = "STUB: not implemented"

	// As implements the errors.As interface to allow unwrapping as either an
	// Error or apierrors.HTTPError, depending on the needs of the caller.
	return ""
}

func (e *Error) As(target any) bool { _ = "STUB: not implemented"; return false }

// Check will attempt to extract a hook Error from a byte slice and return a
// non-nil error, otherwise Check returns nil if no error was found.
func Check(b []byte) error { _ = "STUB: not implemented"; return nil }

func check(e *Error) error { _ = "STUB: not implemented"; return nil }

// TODO(cstockton): Changing this would be a BC break, but it also
// doesn't seem to be the best API. For example returning an error object
// with an http_code field set to 500 would not count as an error.

// TODO(cstockton): this really should be a BadRequest as default and
// the returned code should be bounded to 4XX codes, ideally specific
// 4xx codes.
// if httpCode/100 == 4 { httpCode = http.StatusBadRequest }

func fromBytes(b []byte) (*Error, bool) { _ = "STUB: not implemented"; return nil, false }
