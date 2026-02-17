// Package request defines the Doer interface used by domain packages to perform API requests.
// It is not part of the public API.
package request

import (
	"context"
	"io"
	"net/url"
)

// Doer performs a single API request and decodes the v3 envelope.
// *katapultpro.Client implements this interface.
type Doer interface {
	// Do performs an HTTP request with an optional JSON body and decodes the response into out.
	Do(ctx context.Context, method, path string, query url.Values, body any, out any) error
	// DoWithBody performs a request with a raw body and content type (e.g. image/jpeg).
	DoWithBody(ctx context.Context, method, path string, query url.Values, contentType string, body io.Reader, out any) error
}
