package json

import (
	gojson "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	httperr "github.com/omniboost/go-httperr"
)

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// we have an error
	errorResponse := &ErrorResponse{Response: r}

	// if we have no body, return the error response now
	if r.Body != nil {
		errorResponse.Message = r.Status
		// wrap error in http error so we can handle it properly
		return &httperr.Error{StatusCode: r.StatusCode, Err: errorResponse}
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	err = gojson.Unmarshal(data, errorResponse)
	if err == nil {
		return errorResponse
	}

	// failed to unmarshal, set message to status
	errorResponse.Message = r.Status
	// wrap error in http error so we can handle it properly
	return &httperr.Error{StatusCode: r.StatusCode, Err: errorResponse}

}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Fault code
	Details string `json:"Details"`

	// Fault message
	Message string `json:"Message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d (%v %v)",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Details, r.Message)
}
