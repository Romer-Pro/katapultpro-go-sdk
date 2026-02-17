package katapultpro_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestListPhotos(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/photos" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[{"id":"p1"}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	photos, err := client.ListPhotos(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if len(photos) != 1 || photos[0].ID != "p1" {
		t.Errorf("got photos %+v", photos)
	}
}

func TestGetPhoto(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/photos/p1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"p1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	photo, err := client.GetPhoto(context.Background(), "j1", "p1")
	if err != nil {
		t.Fatal(err)
	}
	if photo.ID != "p1" {
		t.Errorf("got photo %+v", photo)
	}
}

func TestUploadJobPhoto(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/photos" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "image/jpeg" {
			t.Errorf("expected Content-Type image/jpeg, got %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"new-photo"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	photo, err := client.UploadJobPhoto(context.Background(), "j1", bytes.NewReader([]byte("jpeg")))
	if err != nil {
		t.Fatal(err)
	}
	if photo.ID != "new-photo" {
		t.Errorf("got photo %+v", photo)
	}
}

func TestAssociatePhoto(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/photos/p1/associate" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":null,"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	err := client.AssociatePhoto(context.Background(), "j1", "p1", &katapultpro.AssociatePhotoRequest{
		NodeID:           "n1",
		AssociationValue: katapultpro.PtrPhotoAssociationMain(),
	})
	if err != nil {
		t.Fatal(err)
	}
}
