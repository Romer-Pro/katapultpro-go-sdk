package katapultpro_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestWithRateLimit(t *testing.T) {
	var callCount int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	// Use 20ms interval so two calls are at least 20ms apart
	client, err := katapultpro.NewClient("key",
		katapultpro.WithBaseURL(srv.URL),
		katapultpro.WithRateLimit(20*time.Millisecond),
	)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	start := time.Now()
	_, _ = client.ListJobs(ctx, nil)
	_, _ = client.ListJobs(ctx, nil)
	elapsed := time.Since(start)

	if callCount != 2 {
		t.Errorf("expected 2 calls, got %d", callCount)
	}
	// Second request should have been delayed by at least 20ms
	if elapsed < 20*time.Millisecond {
		t.Errorf("rate limit should have delayed second request (elapsed %v)", elapsed)
	}
}
