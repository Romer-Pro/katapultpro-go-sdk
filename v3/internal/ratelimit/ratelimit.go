// Package ratelimit provides an http.RoundTripper that enforces a minimum interval between requests.
// It is used by the v3 client's WithRateLimit option and is not part of the public API.
package ratelimit

import (
	"net/http"
	"sync"
	"time"
)

// Transport wraps an http.RoundTripper and enforces a minimum interval between requests.
type Transport struct {
	Base     http.RoundTripper
	Interval time.Duration
	mu       sync.Mutex
	last     time.Time
}

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.mu.Lock()
	elapsed := time.Since(t.last)
	if elapsed < t.Interval {
		time.Sleep(t.Interval - elapsed)
	}
	t.last = time.Now()
	t.mu.Unlock()
	return t.Base.RoundTrip(req)
}

// NewTransport returns an http.RoundTripper that limits requests to at least interval apart.
func NewTransport(base http.RoundTripper, interval time.Duration) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Transport{Base: base, Interval: interval}
}
