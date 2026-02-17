package katapultpro

import (
	"net/http"
	"time"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/ratelimit"
)

// DefaultRateLimitInterval is the Katapult Pro API minimum interval between requests (1 call per 50ms).
const DefaultRateLimitInterval = 50 * time.Millisecond

// WithRateLimit wraps the client's transport with a rate limiter so requests are at least interval apart.
// Use DefaultRateLimitInterval (50ms) to comply with the API's "1 call per 50ms" general rate limit.
// If the client was created with WithHTTPClient, that client's Transport is wrapped; otherwise http.DefaultTransport is used.
func WithRateLimit(interval time.Duration) ClientOption {
	return func(c *Client) {
		if interval <= 0 {
			return
		}
		base := c.httpClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		c.httpClient = &http.Client{
			Transport:     ratelimit.NewTransport(base, interval),
			CheckRedirect: c.httpClient.CheckRedirect,
			Jar:           c.httpClient.Jar,
			Timeout:       c.httpClient.Timeout,
		}
	}
}
