package katapultpro

import (
	"net/http"
	"net/url"
)

// ClientOption configures a Client. Options are applied in order;
// later options override earlier ones where applicable.
type ClientOption func(*Client)

// WithBaseURL sets the API base URL. The default is "https://katapultpro.com/api" (v3 paths are e.g. /v3/jobs).
// The given raw URL must be valid; invalid URLs are ignored.
func WithBaseURL(raw string) ClientOption {
	return func(c *Client) {
		if u, err := url.Parse(raw); err == nil {
			c.baseURL = u
		}
	}
}

// WithHTTPClient sets the http.Client used for all requests.
// The default is http.DefaultClient. Use this to customize timeouts,
// redirect policy, or to inject a custom transport (e.g. for testing).
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}
