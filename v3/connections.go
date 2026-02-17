package katapultpro

import (
	"context"

	"github.com/romer-pro/katapultpro-go-sdk/v3/connections"
)

// ListConnections returns all connections (and their sections) in the job (v3).
func (c *Client) ListConnections(ctx context.Context, jobID string) ([]Connection, error) {
	return connections.NewClient(c, jobID).List(ctx)
}

// GetConnection returns the specified connection and its sections (v3).
func (c *Client) GetConnection(ctx context.Context, jobID, connectionID string) (*Connection, error) {
	return connections.NewClient(c, jobID).Get(ctx, connectionID)
}

// CreateConnection creates a new connection between two nodes (v3).
func (c *Client) CreateConnection(ctx context.Context, jobID string, req *CreateConnectionRequest) (*Connection, error) {
	return connections.NewClient(c, jobID).Create(ctx, req)
}

// UpdateConnection updates the specified connection, or creates it with the given ID if it does not exist (v3).
func (c *Client) UpdateConnection(ctx context.Context, jobID, connectionID string, req *UpdateConnectionRequest, opts *UpdateConnectionOptions) (*Connection, error) {
	return connections.NewClient(c, jobID).Update(ctx, connectionID, req, opts)
}

// DeleteConnection deletes the connection and all its sections (v3).
func (c *Client) DeleteConnection(ctx context.Context, jobID, connectionID string) error {
	return connections.NewClient(c, jobID).Delete(ctx, connectionID)
}
