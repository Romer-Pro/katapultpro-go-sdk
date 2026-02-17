package katapultpro_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romer-pro/katapultpro-go-sdk/v3"
)

func TestListJobs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/jobs" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[{"id":"j1","name":"Job 1"}],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	ctx := context.Background()

	jobs, err := client.ListJobs(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 1 || jobs[0].ID != "j1" || jobs[0].Name != "Job 1" {
		t.Errorf("got jobs %+v", jobs)
	}
}

func TestListJobs_withOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("includeArchived") != "true" {
			t.Errorf("expected includeArchived=true, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":[],"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	jobs, err := client.ListJobs(context.Background(), &katapultpro.ListJobsOptions{IncludeArchived: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 0 {
		t.Errorf("expected empty list, got %d", len(jobs))
	}
}

func TestGetJob(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/jobs/job-123" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"job-123","name":"My Job","status":"active"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	job, err := client.GetJob(context.Background(), "job-123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if job.ID != "job-123" || job.Name != "My Job" || job.Status != katapultpro.JobStatusActive {
		t.Errorf("got job %+v", job)
	}
}

func TestCreateJob(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"new-id","name":"New Job"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	job, err := client.CreateJob(context.Background(), &katapultpro.CreateJobRequest{Name: "New Job", Model: "m1"})
	if err != nil {
		t.Fatal(err)
	}
	if job.ID != "new-id" || job.Name != "New Job" {
		t.Errorf("got job %+v", job)
	}
}

func TestUpdateJob(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/job-1" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"id":"job-1","name":"Updated"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	job, err := client.UpdateJob(context.Background(), "job-1", &katapultpro.UpdateJobRequest{Name: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if job.Name != "Updated" {
		t.Errorf("got job %+v", job)
	}
}

func TestGetJobStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/jobs/j1/status" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":{"status":"archived"},"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	status, err := client.GetJobStatus(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if status != katapultpro.JobStatusArchived {
		t.Errorf("got status %q", status)
	}
}

func TestUpdateJobStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/jobs/j1/status" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","data":null,"meta":{"token_count":9999,"last_refill_time":0}}`))
	}))
	defer srv.Close()

	client, _ := katapultpro.NewClient("key", katapultpro.WithBaseURL(srv.URL))
	err := client.UpdateJobStatus(context.Background(), "j1", katapultpro.JobStatusActive)
	if err != nil {
		t.Fatal(err)
	}
}
