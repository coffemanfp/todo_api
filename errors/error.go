package errors

import "fmt"

// HTTPError represents a error to present to the client.
// Implements the error interface.
type HTTPError struct {
	httpCode int
	message  string
}

func (h HTTPError) Error() string {
	return h.message
}

// HTTPCode gets the http code of the error. Returns 0 if it's not available.
func (h HTTPError) HTTPCode() int {
	return h.httpCode
}

// NewHTTPError initialices a new error with a HTTPError implementation.
//
//	 @param httpCode: represents the http error code for the http response.
//	 @param m string: message to be presented to the client.
//	 @param a ...interface{}: optional arguments for the message.
//		@return $1 error: new HTTPError error implementation instance.
func NewHTTPError(httpCode int, m string, a ...interface{}) error {
	return HTTPError{
		httpCode: httpCode,
		message:  fmt.Sprintf(m, a...),
	}
}
