package katapultpro_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	katapultpro "github.com/romer-pro/katapultpro-go-sdk/v2"
)

func TestNewClient_MissingAPIKey(t *testing.T) {
	_, err := katapultpro.NewClient("")
	if err != katapultpro.ErrMissingAPIKey {
		t.Errorf("expected ErrMissingAPIKey, got %v", err)
	}
}

func TestNewClient_WithAPIKey(t *testing.T) {
	client, err := katapultpro.NewClient("test-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestGetPhotoURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify path
		expectedPath := "/v2/jobs/job123/photoURL/photo456"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify query params
		if r.URL.Query().Get("api_key") != "test-key" {
			t.Errorf("expected api_key=test-key, got %s", r.URL.Query().Get("api_key"))
		}
		if r.URL.Query().Get("file_size") != "extra_large" {
			t.Errorf("expected file_size=extra_large, got %s", r.URL.Query().Get("file_size"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"url":"https://firebasestorage.googleapis.com/v0/b/test/photo.webp?token=abc"}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("test-key", katapultpro.WithBaseURL(srv.URL))
	resp, err := client.GetPhotoURL(context.Background(), "job123", "photo456", katapultpro.PhotoSizeExtraLarge)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.URL != "https://firebasestorage.googleapis.com/v0/b/test/photo.webp?token=abc" {
		t.Errorf("unexpected URL: %s", resp.URL)
	}
}

func TestGetPhotoURL_DefaultSize(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// When no size is specified, file_size should not be in query
		if r.URL.Query().Get("file_size") != "" {
			t.Errorf("expected no file_size param, got %s", r.URL.Query().Get("file_size"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"url":"https://example.com/photo.jpg"}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("test-key", katapultpro.WithBaseURL(srv.URL))
	resp, err := client.GetPhotoURL(context.Background(), "job123", "photo456", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.URL != "https://example.com/photo.jpg" {
		t.Errorf("unexpected URL: %s", resp.URL)
	}
}

func TestGetPhotoURL_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`Photo not found`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("test-key", katapultpro.WithBaseURL(srv.URL))
	_, err := client.GetPhotoURL(context.Background(), "job123", "invalid", katapultpro.PhotoSizeFull)
	if err == nil {
		t.Fatal("expected error")
	}

	apiErr, ok := err.(*katapultpro.APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", apiErr.StatusCode)
	}
}

func TestPhotoSize_IsValid(t *testing.T) {
	tests := []struct {
		size  katapultpro.PhotoSize
		valid bool
	}{
		{katapultpro.PhotoSizeFull, true},
		{katapultpro.PhotoSizeExtraLarge, true},
		{katapultpro.PhotoSizeLarge, true},
		{katapultpro.PhotoSizeSmall, true},
		{katapultpro.PhotoSizeTiny, true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := tt.size.IsValid(); got != tt.valid {
			t.Errorf("PhotoSize(%q).IsValid() = %v, want %v", tt.size, got, tt.valid)
		}
	}
}
