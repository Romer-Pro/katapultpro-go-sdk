package katapultpro_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestNewClient_MissingAPIKey(t *testing.T) {
	_, err := katapultpro.NewClient("")
	if err == nil {
		t.Fatal("expected error for empty API key, got nil")
	}
	if !errors.Is(err, katapultpro.ErrMissingAPIKey) {
		t.Fatalf("expected ErrMissingAPIKey, got %v", err)
	}
}

func TestNewClient_WithAPIKey(t *testing.T) {
	client, err := katapultpro.NewClient("test-api-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestClient_APIKeyInQueryParam(t *testing.T) {
	const testAPIKey = "my-secret-api-key"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify API key is in query parameter, not header
		if got := r.URL.Query().Get("api_key"); got != testAPIKey {
			t.Errorf("expected api_key=%q in query, got %q", testAPIKey, got)
		}
		if auth := r.Header.Get("Authorization"); auth != "" {
			t.Errorf("expected no Authorization header, got %q", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient(testAPIKey, katapultpro.WithBaseURL(srv.URL))
	_, err := client.ListJobs(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
}

func ExampleNewClient() {
	client, err := katapultpro.NewClient("your-api-key")
	if err != nil {
		panic(err)
	}

	// Customize with options
	_ = client
}

func ExampleNewClient_withOptions() {
	client, err := katapultpro.NewClient("api-key",
		katapultpro.WithBaseURL("https://api.example.com"),
		katapultpro.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	)
	if err != nil {
		panic(err)
	}
	_ = client
}

func ExampleClient_Get() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// v3 envelope: { status, data, meta }
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"123","name":"test"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	ctx := context.Background()

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := client.Get(ctx, "/v3/jobs/123", &result); err != nil {
		panic(err)
	}
	fmt.Println(result.ID, result.Name)
	// Output: 123 test
}

func ExampleClient_ListJobs() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// v3 list response: data is an array
		_, _ = w.Write([]byte(`{"status":"success","data":[{"id":"job1","name":"Job One"},{"id":"job2","name":"Job Two"}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	ctx := context.Background()

	jobs, err := client.ListJobs(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(jobs), jobs[0].Name)
	// Output: 2 Job One
}

func ExampleAPIError() {
	var err error = &katapultpro.APIError{StatusCode: 404, Type: "not_found", Message: "not found"}
	fmt.Println(err.Error())
	// Output: katapultpro api error 404 (not_found): not found
}
