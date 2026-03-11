package katapultpro

import (
	"net/http"
	"net/url"
	"strings"
)

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the API.
// Use this for private Katapult Pro servers (e.g., "https://yourcompany.katapultpro.com/api/").
// A trailing slash is automatically appended if not present.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		if u, err := url.Parse(baseURL); err == nil {
			c.baseURL = u
		}
	}
}

// WithHTTPClient sets a custom HTTP client.
// Use this to configure timeouts, proxies, or other transport settings.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}
