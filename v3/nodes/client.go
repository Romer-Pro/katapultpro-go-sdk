package nodes

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// Client performs node API operations for a single job. Create with NewClient(do, jobID).
type Client struct {
	do    request.Doer
	jobID string
}

// NewClient returns a nodes client for the given job.
func NewClient(do request.Doer, jobID string) *Client {
	return &Client{do: do, jobID: jobID}
}

// List returns all nodes in the job (v3).
func (c *Client) List(ctx context.Context) ([]Node, error) {
	path := "v3/jobs/" + c.jobID + "/nodes"
	var out []Node
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified node (v3).
func (c *Client) Get(ctx context.Context, nodeID string) (*Node, error) {
	path := "v3/jobs/" + c.jobID + "/nodes/" + nodeID
	var node Node
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &node); err != nil {
		return nil, err
	}
	return &node, nil
}

// Create creates a new node in the job (v3).
func (c *Client) Create(ctx context.Context, req *CreateNodeRequest) (*Node, error) {
	path := "v3/jobs/" + c.jobID + "/nodes"
	var node Node
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &node); err != nil {
		return nil, err
	}
	return &node, nil
}

// Update updates the specified node, or creates it with the given ID if it does not exist (v3).
func (c *Client) Update(ctx context.Context, nodeID string, req *UpdateNodeRequest, opts *UpdateNodeOptions) (*Node, error) {
	path := "v3/jobs/" + c.jobID + "/nodes/" + nodeID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var node Node
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &node); err != nil {
		return nil, err
	}
	return &node, nil
}

// UploadPhoto uploads a photo (image/jpeg) and associates it to the node (v3).
func (c *Client) UploadPhoto(ctx context.Context, nodeID string, imageData io.Reader, opts *UploadNodePhotoOptions) (*photos.Photo, error) {
	path := "v3/jobs/" + c.jobID + "/nodes/" + nodeID + "/photos"
	var q url.Values
	if opts != nil && opts.AssociationValue.IsValid() {
		q = url.Values{}
		q.Set("association_value", opts.AssociationValue.String())
	}
	var photo photos.Photo
	if err := c.do.DoWithBody(ctx, http.MethodPost, path, q, "image/jpeg", imageData, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Delete deletes the specified node (v3).
func (c *Client) Delete(ctx context.Context, nodeID string) error {
	path := "v3/jobs/" + c.jobID + "/nodes/" + nodeID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}
