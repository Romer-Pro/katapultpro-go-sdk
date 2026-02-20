// Package transport performs HTTP requests for the Katapult Pro API client.
// It is used by the v3 client and is not part of the public API.
package transport

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Do executes an HTTP request and returns the status code and response body.
// The caller is responsible for parsing the response (e.g. with envelope.Parse).
func Do(ctx context.Context, method string, baseURL *url.URL, path string, query url.Values, body io.Reader, contentType string, apiKey string, client *http.Client) (statusCode int, respBody []byte, err error) {
	path = strings.TrimPrefix(path, "/")
	segments := strings.Split(path, "/")
	u := baseURL.JoinPath(segments...)
	// Katapult Pro uses query parameter authentication: ?api_key=<key>
	if query == nil {
		query = url.Values{}
	}
	if apiKey != "" {
		query.Set("api_key", apiKey)
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return 0, nil, fmt.Errorf("create request: %w", err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("read response: %w", err)
	}
	return resp.StatusCode, respBody, nil
}
