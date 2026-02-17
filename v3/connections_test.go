package katapultpro_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestListConnections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[{"id":"c1","node_id_1":"n1","node_id_2":"n2"}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	conns, err := client.ListConnections(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if len(conns) != 1 || conns[0].ID != "c1" {
		t.Errorf("got connections %+v", conns)
	}
}

func TestGetConnection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"c1","node_id_1":"n1","node_id_2":"n2"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	conn, err := client.GetConnection(context.Background(), "j1", "c1")
	if err != nil {
		t.Fatal(err)
	}
	if conn.ID != "c1" || conn.NodeID1 != "n1" || conn.NodeID2 != "n2" {
		t.Errorf("got connection %+v", conn)
	}
}

func TestCreateConnection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/connections" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"new-c","node_id_1":"a","node_id_2":"b"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	conn, err := client.CreateConnection(context.Background(), "j1", &katapultpro.CreateConnectionRequest{NodeID1: "a", NodeID2: "b"})
	if err != nil {
		t.Fatal(err)
	}
	if conn.ID != "new-c" {
		t.Errorf("got connection %+v", conn)
	}
}

func TestUpdateConnection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/connections/c1" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"c1"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	conn, err := client.UpdateConnection(context.Background(), "j1", "c1", &katapultpro.UpdateConnectionRequest{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if conn.ID != "c1" {
		t.Errorf("got connection %+v", conn)
	}
}

func TestDeleteConnection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/v3/jobs/j1/connections/c1" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":null,"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	err := client.DeleteConnection(context.Background(), "j1", "c1")
	if err != nil {
		t.Fatal(err)
	}
}
