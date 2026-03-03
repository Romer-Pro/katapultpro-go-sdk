package katapultpro

import (
	"errors"
	"fmt"
)

// ErrMissingAPIKey is returned when an empty API key is passed to NewClient.
var ErrMissingAPIKey = errors.New("katapultpro: API key is required")

// APIError represents an error response from the Katapult Pro API.
type APIError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("katapultpro: API error %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("katapultpro: API error %d", e.StatusCode)
}
