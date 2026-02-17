package jobs

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
)

// Client performs jobs API operations. Create with NewClient(do).
// Methods do not take the doer; it is held by the Client.
type Client struct {
	do request.Doer
}

// NewClient returns a jobs client that uses do for requests.
func NewClient(do request.Doer) *Client {
	return &Client{do: do}
}

// List returns all jobs accessible to the requester (v3).
func (c *Client) List(ctx context.Context, opts *ListJobsOptions) ([]Job, error) {
	var q url.Values
	if opts != nil {
		if opts.IncludeArchived {
			q = url.Values{}
			q.Set("includeArchived", "true")
		}
		if opts.MetadataFilter != "" {
			if q == nil {
				q = url.Values{}
			}
			q.Set("metadataFilter", opts.MetadataFilter)
		}
	}
	var out []Job
	if err := c.do.Do(ctx, http.MethodGet, "v3/jobs", q, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns partial or full job data for the given job ID (v3).
func (c *Client) Get(ctx context.Context, jobID string, opts *GetJobOptions) (*Job, error) {
	path := "v3/jobs/" + jobID
	var q url.Values
	if opts != nil && len(opts.Paths) > 0 {
		q = url.Values{}
		parts := make([]string, len(opts.Paths))
		for i, p := range opts.Paths {
			parts[i] = string(p)
		}
		q.Set("paths", strings.Join(parts, ","))
	}
	var job Job
	if err := c.do.Do(ctx, http.MethodGet, path, q, nil, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// Create creates a new job (v3).
func (c *Client) Create(ctx context.Context, req *CreateJobRequest) (*Job, error) {
	var job Job
	if err := c.do.Do(ctx, http.MethodPost, "v3/jobs", nil, req, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// Update updates the specified job (v3).
func (c *Client) Update(ctx context.Context, jobID string, req *UpdateJobRequest) (*Job, error) {
	path := "v3/jobs/" + jobID
	var job Job
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// GetStatus returns the status of the job (v3).
func (c *Client) GetStatus(ctx context.Context, jobID string) (JobStatus, error) {
	path := "v3/jobs/" + jobID + "/status"
	var out JobStatusResponse
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return "", err
	}
	return out.Status, nil
}

// UpdateStatus sets the job status (v3).
func (c *Client) UpdateStatus(ctx context.Context, jobID string, status JobStatus) error {
	path := "v3/jobs/" + jobID + "/status"
	req := &UpdateJobStatusRequest{Status: status}
	return c.do.Do(ctx, http.MethodPost, path, nil, req, nil)
}
