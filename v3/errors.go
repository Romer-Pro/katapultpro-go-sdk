package katapultpro

import "fmt"

// APIError represents an error response from the Katapult Pro API (v3).
// It is returned for non-2xx HTTP responses or when the response envelope has status "error".
// Use errors.As to detect it:
//
//	var apiErr *katapultpro.APIError
//	if errors.As(err, &apiErr) {
//	    log.Printf("API error %d (%s): %s", apiErr.StatusCode, apiErr.Type, apiErr.Message)
//	}
type APIError struct {
	StatusCode int    `json:"-"`                 // HTTP status code (e.g. 404, 429).
	Message    string `json:"message,omitempty"`   // Human-readable error message from the API.
	Type       string `json:"type,omitempty"`     // Error type from the API (e.g. "not_found").
	Meta       *Meta  `json:"meta,omitempty"`    // Token bucket state when the error was returned.
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Type != "" {
		return fmt.Sprintf("katapultpro api error %d (%s): %s", e.StatusCode, e.Type, e.Message)
	}
	return fmt.Sprintf("katapultpro api error %d: %s", e.StatusCode, e.Message)
}

// Ensure APIError is recognized by errors.As.
var _ error = (*APIError)(nil)
