package photos

import (
	"context"
	"io"
	"net/http"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
)

// Client performs photo API operations for a single job. Create with NewClient(do, jobID).
type Client struct {
	do    request.Doer
	jobID string
}

// NewClient returns a photos client for the given job.
func NewClient(do request.Doer, jobID string) *Client {
	return &Client{do: do, jobID: jobID}
}

// List returns all photo records in the job (v3).
func (c *Client) List(ctx context.Context) ([]Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos"
	var out []Photo
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified photo record (v3).
func (c *Client) Get(ctx context.Context, photoID string) (*Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID
	var photo Photo
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Upload uploads a photo (image/jpeg) to the job (v3).
func (c *Client) Upload(ctx context.Context, imageData io.Reader) (*Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos"
	var photo Photo
	if err := c.do.DoWithBody(ctx, http.MethodPost, path, nil, "image/jpeg", imageData, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Associate associates (or unassociates) the photo to a node or section (v3).
func (c *Client) Associate(ctx context.Context, photoID string, req *AssociatePhotoRequest) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/associate"
	return c.do.Do(ctx, http.MethodPost, path, nil, req, nil)
}
