package freshclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	http "net/http"
)

type Client struct {
	// The http.Client to use for requests
	HTTPClient *http.Client
	// The API key to use for requests
	APIKey *string
	// The API endpoint to use for requests
	APIEndpoint *string
}

// NewClient creates a new FreshClient.
func NewClient(apiKey string, apiEndpoint string) *Client {
	// Check if the API key and endpoint are set
	if apiKey == "" || apiEndpoint == "" {
		panic("apiKey or apiEndpoint not set")
	}

	return &Client{
		HTTPClient:  http.DefaultClient,
		APIKey:      &apiKey,
		APIEndpoint: &apiEndpoint,
	}
}

// MakeRequest makes a request to the FreshService API.
func (client *Client) MakeRequest(method string, url string, body interface{}) (*http.Response, error) {
	req := &http.Request{}
	if body != nil {
		marshalledBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		// Create a new request
		req, err = http.NewRequest(method, url, bytes.NewReader(marshalledBody))
		if err != nil {
			return nil, err
		}
	} else {
		err := error(nil)
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	// Add the API key to the request using basic auth
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(*client.APIKey, "x")

	// Make the request
	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	// Check for errors
	if _, errorExist := ErrorMessages[resp.StatusCode]; errorExist {
		return nil, NewErrorByCode(resp.StatusCode, ErrorMessages[resp.StatusCode])
	}

	return resp, nil
}

// Error handeling
// APIError represents an error in the API with additional details.
type APIError struct {
	Code        int    // HTTP status code
	Text        string // Textual description
	Description string // Additional details about the error
}

// NewAPIError creates a new APIError instance.
func NewAPIError(code int, text, description string) *APIError {
	return &APIError{
		Code:        code,
		Text:        text,
		Description: description,
	}
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("HTTP %d - %s: %s", e.Code, e.Text, e.Description)
}

// ErrorCode constants for common API errors.
const (
	ErrClientValidation        = 400
	ErrAuthenticationFailure   = 401
	ErrAccessDenied            = 403
	ErrResourceNotFound        = 404
	ErrMethodNotAllowed        = 405
	ErrUnsupportedAcceptHeader = 406
	ErrInconsistentState       = 409
	ErrUnsupportedContentType  = 415
	ErrRateLimitExceeded       = 429
	ErrUnexpectedServerError   = 500
)

// ErrorMessages map error codes to their respective error messages.
var ErrorMessages = map[int]string{
	ErrClientValidation:        "Client or Validation Error",
	ErrAuthenticationFailure:   "Authentication Failure",
	ErrAccessDenied:            "Access Denied",
	ErrResourceNotFound:        "Requested Resource not Found",
	ErrMethodNotAllowed:        "Method not allowed",
	ErrUnsupportedAcceptHeader: "Unsupported Accept Header",
	ErrInconsistentState:       "Inconsistent/Conflicting State",
	ErrUnsupportedContentType:  "Unsupported Content-type",
	ErrRateLimitExceeded:       "Rate Limit Exceeded",
	ErrUnexpectedServerError:   "Unexpected Server Error",
}

// NewErrorByCode creates a new APIError based on the provided error code.
func NewErrorByCode(code int, description string) *APIError {
	text, exists := ErrorMessages[code]
	if !exists {
		text = "Unknown Error"
	}
	return NewAPIError(code, text, description)
}
