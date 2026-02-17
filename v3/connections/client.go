package connections

import (
	"context"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
	"github.com/romer-pro/katapultpro-go-sdk/v3/sections"
)

// Client performs connection API operations for a single job. Create with NewClient(do, jobID).
type Client struct {
	do    request.Doer
	jobID string
}

// NewClient returns a connections client for the given job.
func NewClient(do request.Doer, jobID string) *Client {
	return &Client{do: do, jobID: jobID}
}

// Sections returns a sections client for the given connection.
func (c *Client) Sections(connectionID string) *sections.Client {
	return sections.NewClient(c.do, c.jobID, connectionID)
}

// List returns all connections (and their sections) in the job (v3).
func (c *Client) List(ctx context.Context) ([]Connection, error) {
	path := "v3/jobs/" + c.jobID + "/connections"
	var out []Connection
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified connection and its sections (v3).
func (c *Client) Get(ctx context.Context, connectionID string) (*Connection, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + connectionID
	var conn Connection
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &conn); err != nil {
		return nil, err
	}
	return &conn, nil
}

// Create creates a new connection between two nodes (v3).
func (c *Client) Create(ctx context.Context, req *CreateConnectionRequest) (*Connection, error) {
	path := "v3/jobs/" + c.jobID + "/connections"
	var conn Connection
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &conn); err != nil {
		return nil, err
	}
	return &conn, nil
}

// Update updates the specified connection, or creates it with the given ID if it does not exist (v3).
func (c *Client) Update(ctx context.Context, connectionID string, req *UpdateConnectionRequest, opts *UpdateConnectionOptions) (*Connection, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + connectionID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var conn Connection
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &conn); err != nil {
		return nil, err
	}
	return &conn, nil
}

// Delete deletes the connection and all its sections (v3).
func (c *Client) Delete(ctx context.Context, connectionID string) error {
	path := "v3/jobs/" + c.jobID + "/connections/" + connectionID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}
