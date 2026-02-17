package katapultpro

import (
	"context"
	"io"

	"github.com/romer-pro/katapultpro-go-sdk/v3/nodes"
)

// ListNodes returns all nodes in the specified job (v3).
func (c *Client) ListNodes(ctx context.Context, jobID string) ([]Node, error) {
	return nodes.NewClient(c, jobID).List(ctx)
}

// GetNode returns the specified node (v3).
func (c *Client) GetNode(ctx context.Context, jobID, nodeID string) (*Node, error) {
	return nodes.NewClient(c, jobID).Get(ctx, nodeID)
}

// CreateNode creates a new node in the job (v3).
func (c *Client) CreateNode(ctx context.Context, jobID string, req *CreateNodeRequest) (*Node, error) {
	return nodes.NewClient(c, jobID).Create(ctx, req)
}

// UpdateNode updates the specified node, or creates it with the given ID if it does not exist (v3).
func (c *Client) UpdateNode(ctx context.Context, jobID, nodeID string, req *UpdateNodeRequest, opts *UpdateNodeOptions) (*Node, error) {
	return nodes.NewClient(c, jobID).Update(ctx, nodeID, req, opts)
}

// UploadNodePhoto uploads a photo (image/jpeg) and associates it to the node (v3).
func (c *Client) UploadNodePhoto(ctx context.Context, jobID, nodeID string, imageData io.Reader, opts *UploadNodePhotoOptions) (*Photo, error) {
	return nodes.NewClient(c, jobID).UploadPhoto(ctx, nodeID, imageData, opts)
}

// DeleteNode deletes the specified node (v3).
func (c *Client) DeleteNode(ctx context.Context, jobID, nodeID string) error {
	return nodes.NewClient(c, jobID).Delete(ctx, nodeID)
}
