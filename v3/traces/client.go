package traces

import (
	"context"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
)

// Client performs trace API operations for a single job. Create with NewClient(do, jobID).
type Client struct {
	do    request.Doer
	jobID string
}

// NewClient returns a traces client for the given job.
func NewClient(do request.Doer, jobID string) *Client {
	return &Client{do: do, jobID: jobID}
}

// List returns all traces in the job (v3).
func (c *Client) List(ctx context.Context) ([]Trace, error) {
	path := "v3/jobs/" + c.jobID + "/traces"
	var out []Trace
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified trace (v3).
func (c *Client) Get(ctx context.Context, traceID string) (*Trace, error) {
	path := "v3/jobs/" + c.jobID + "/traces/" + traceID
	var trace Trace
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &trace); err != nil {
		return nil, err
	}
	return &trace, nil
}

// Create creates a new trace in the job (v3).
func (c *Client) Create(ctx context.Context, req *CreateTraceRequest) (*Trace, error) {
	path := "v3/jobs/" + c.jobID + "/traces"
	var trace Trace
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &trace); err != nil {
		return nil, err
	}
	return &trace, nil
}

// Update updates the specified trace, or creates it with the given ID if it does not exist (v3).
func (c *Client) Update(ctx context.Context, traceID string, req *UpdateTraceRequest, opts *UpdateTraceOptions) (*Trace, error) {
	path := "v3/jobs/" + c.jobID + "/traces/" + traceID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var trace Trace
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &trace); err != nil {
		return nil, err
	}
	return &trace, nil
}

// Delete deletes the specified trace (v3).
func (c *Client) Delete(ctx context.Context, traceID string) error {
	path := "v3/jobs/" + c.jobID + "/traces/" + traceID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}
