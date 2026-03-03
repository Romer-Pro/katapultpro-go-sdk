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

// Connection returns a ConnectionScope for the given connection, enabling drill-down to sections.
// Example: client.Job(jobID).Connections().Connection(connID).Sections().List(ctx)
func (c *Client) Connection(connectionID string) *ConnectionScope {
	return &ConnectionScope{do: c.do, jobID: c.jobID, connectionID: connectionID}
}

// Sections returns a sections client for the given connection.
// Deprecated: Use Connection(connectionID).Sections() instead for consistent scope pattern.
func (c *Client) Sections(connectionID string) *sections.Client {
	return sections.NewClient(c.do, c.jobID, connectionID)
}

// ConnectionScope scopes operations to a single connection. Use Connections().Connection(connectionID) to create one.
type ConnectionScope struct {
	do           request.Doer
	jobID        string
	connectionID string
}

// ConnectionID returns the scoped connection ID.
func (s *ConnectionScope) ConnectionID() string { return s.connectionID }

// Get returns the connection and its sections.
func (s *ConnectionScope) Get(ctx context.Context) (*Connection, error) {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID
	var conn Connection
	if err := s.do.Do(ctx, http.MethodGet, path, nil, nil, &conn); err != nil {
		return nil, err
	}
	return &conn, nil
}

// Update updates the connection. Use opts.OnlyIfExists to avoid creating with the given id.
func (s *ConnectionScope) Update(ctx context.Context, req *UpdateConnectionRequest, opts *UpdateConnectionOptions) (*Connection, error) {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var conn Connection
	if err := s.do.Do(ctx, http.MethodPost, path, q, req, &conn); err != nil {
		return nil, err
	}
	return &conn, nil
}

// Delete deletes the connection and all its sections.
func (s *ConnectionScope) Delete(ctx context.Context) error {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID
	return s.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// Sections returns a sections client for this connection.
func (s *ConnectionScope) Sections() *sections.Client {
	return sections.NewClient(s.do, s.jobID, s.connectionID)
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
