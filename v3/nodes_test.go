package katapultpro_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestListNodes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/jobs/j1/nodes" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[{"id":"n1","latitude":1.0,"longitude":2.0}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	nodes, err := client.ListNodes(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 1 || nodes[0].ID != "n1" {
		t.Errorf("got nodes %+v", nodes)
	}
}

func TestGetNode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/nodes/n1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"n1","latitude":10,"longitude":20},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	node, err := client.GetNode(context.Background(), "j1", "n1")
	if err != nil {
		t.Fatal(err)
	}
	if node.ID != "n1" || node.Latitude != 10 || node.Longitude != 20 {
		t.Errorf("got node %+v", node)
	}
}

func TestCreateNode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/nodes" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"new-node","latitude":0,"longitude":0},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	node, err := client.CreateNode(context.Background(), "j1", &katapultpro.CreateNodeRequest{Latitude: 0, Longitude: 0})
	if err != nil {
		t.Fatal(err)
	}
	if node.ID != "new-node" {
		t.Errorf("got node %+v", node)
	}
}

func TestUpdateNode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/nodes/n1" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"n1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	node, err := client.UpdateNode(context.Background(), "j1", "n1", &katapultpro.UpdateNodeRequest{Latitude: 5}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if node.ID != "n1" {
		t.Errorf("got node %+v", node)
	}
}

func TestUploadNodePhoto(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/nodes/n1/photos" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "image/jpeg" {
			t.Errorf("expected Content-Type image/jpeg, got %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"photo-1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	photo, err := client.UploadNodePhoto(context.Background(), "j1", "n1", bytes.NewReader([]byte("fake-jpeg")), nil)
	if err != nil {
		t.Fatal(err)
	}
	if photo.ID != "photo-1" {
		t.Errorf("got photo %+v", photo)
	}
}

func TestDeleteNode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/v3/jobs/j1/nodes/n1" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":null,"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	err := client.DeleteNode(context.Background(), "j1", "n1")
	if err != nil {
		t.Fatal(err)
	}
}
