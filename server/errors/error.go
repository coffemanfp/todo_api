package errors

import "fmt"

// HTTPError represents a error to present to the client.
// Implements the error interface.
type HTTPError struct {
	Code    int
	Message string
}

func (h HTTPError) Error() string {
	return h.Message
}

// NewHTTPError initialices a new error with a HTTPError implementation.
//
//	 @param code: represents the http error code for the http response.
//	 @param m string: message to be presented to the client.
//	 @param a ...interface{}: optional arguments for the message.
//		@return $1 error: new HTTPError error implementation instance.
func NewHTTPError(code int, m string, a ...interface{}) error {
	return HTTPError{
		Code:    code,
		Message: fmt.Sprintf(m, a...),
	}
}
