package katapultpro_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestListSections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1/sections" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[{"key":"s1","latitude":1}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	sections, err := client.ListSections(context.Background(), "j1", "c1")
	if err != nil {
		t.Fatal(err)
	}
	if len(sections) != 1 || sections[0].Key != "s1" {
		t.Errorf("got sections %+v", sections)
	}
}

func TestGetSection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1/sections/s1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"key":"s1","latitude":2},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	section, err := client.GetSection(context.Background(), "j1", "c1", "s1")
	if err != nil {
		t.Fatal(err)
	}
	if section.Key != "s1" || section.Latitude != 2 {
		t.Errorf("got section %+v", section)
	}
}

func TestCreateSection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/connections/c1/sections" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"key":"new-s"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	section, err := client.CreateSection(context.Background(), "j1", "c1", &katapultpro.CreateSectionRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if section.Key != "new-s" {
		t.Errorf("got section %+v", section)
	}
}

func TestUpdateSection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1/sections/s1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"key":"s1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	section, err := client.UpdateSection(context.Background(), "j1", "c1", "s1", &katapultpro.UpdateSectionRequest{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if section.Key != "s1" {
		t.Errorf("got section %+v", section)
	}
}

func TestUploadSectionPhoto(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1/sections/s1/photos" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"photo-1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	photo, err := client.UploadSectionPhoto(context.Background(), "j1", "c1", "s1", bytes.NewReader([]byte("jpeg")), nil)
	if err != nil {
		t.Fatal(err)
	}
	if photo.ID != "photo-1" {
		t.Errorf("got photo %+v", photo)
	}
}

func TestDeleteSection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/v3/jobs/j1/connections/c1/sections/s1" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":null,"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	err := client.DeleteSection(context.Background(), "j1", "c1", "s1")
	if err != nil {
		t.Fatal(err)
	}
}
