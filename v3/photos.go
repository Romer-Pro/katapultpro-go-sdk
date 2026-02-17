package katapultpro

import (
	"context"
	"io"

	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// ListPhotos returns all photo records in the job (v3).
func (c *Client) ListPhotos(ctx context.Context, jobID string) ([]Photo, error) {
	return photos.NewClient(c, jobID).List(ctx)
}

// GetPhoto returns the specified photo record (v3).
func (c *Client) GetPhoto(ctx context.Context, jobID, photoID string) (*Photo, error) {
	return photos.NewClient(c, jobID).Get(ctx, photoID)
}

// UploadJobPhoto uploads a photo (image/jpeg) to the job (v3).
func (c *Client) UploadJobPhoto(ctx context.Context, jobID string, imageData io.Reader) (*Photo, error) {
	return photos.NewClient(c, jobID).Upload(ctx, imageData)
}

// AssociatePhoto associates (or unassociates) the photo to a node or section (v3).
func (c *Client) AssociatePhoto(ctx context.Context, jobID, photoID string, req *AssociatePhotoRequest) error {
	return photos.NewClient(c, jobID).Associate(ctx, photoID, req)
}
