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

// ListPhotoElements returns all elements on the photo (v3).
func (c *Client) ListPhotoElements(ctx context.Context, jobID, photoID string) ([]PhotoElement, error) {
	return photos.NewClient(c, jobID).ListElements(ctx, photoID)
}

// GetPhotoElement returns the specified photo element (v3).
func (c *Client) GetPhotoElement(ctx context.Context, jobID, photoID, elementID string) (*PhotoElement, error) {
	return photos.NewClient(c, jobID).GetElement(ctx, photoID, elementID)
}

// CreatePhotoElement creates a new element on the photo (v3).
func (c *Client) CreatePhotoElement(ctx context.Context, jobID, photoID string, req *CreatePhotoElementRequest) (*PhotoElement, error) {
	return photos.NewClient(c, jobID).CreateElement(ctx, photoID, req)
}

// UpdatePhotoElement updates the specified photo element (v3).
func (c *Client) UpdatePhotoElement(ctx context.Context, jobID, photoID, elementID string, req *UpdatePhotoElementRequest, opts *UpdatePhotoElementOptions) (*PhotoElement, error) {
	return photos.NewClient(c, jobID).UpdateElement(ctx, photoID, elementID, req, opts)
}

// DeletePhotoElement deletes the specified photo element (v3).
func (c *Client) DeletePhotoElement(ctx context.Context, jobID, photoID, elementID string) error {
	return photos.NewClient(c, jobID).DeleteElement(ctx, photoID, elementID)
}

// ListPhotoCalibrationAnchors returns all calibration anchors on the photo (v3).
func (c *Client) ListPhotoCalibrationAnchors(ctx context.Context, jobID, photoID string) ([]PhotoCalibrationAnchor, error) {
	return photos.NewClient(c, jobID).ListCalibrationAnchors(ctx, photoID)
}

// GetPhotoCalibrationAnchor returns the specified calibration anchor (v3).
func (c *Client) GetPhotoCalibrationAnchor(ctx context.Context, jobID, photoID, anchorID string) (*PhotoCalibrationAnchor, error) {
	return photos.NewClient(c, jobID).GetCalibrationAnchor(ctx, photoID, anchorID)
}

// CreatePhotoCalibrationAnchor creates a new calibration anchor on the photo (v3).
func (c *Client) CreatePhotoCalibrationAnchor(ctx context.Context, jobID, photoID string, req *CreatePhotoCalibrationAnchorRequest) (*PhotoCalibrationAnchor, error) {
	return photos.NewClient(c, jobID).CreateCalibrationAnchor(ctx, photoID, req)
}

// UpdatePhotoCalibrationAnchor updates the specified calibration anchor (v3).
func (c *Client) UpdatePhotoCalibrationAnchor(ctx context.Context, jobID, photoID, anchorID string, req *UpdatePhotoCalibrationAnchorRequest, opts *UpdatePhotoCalibrationAnchorOptions) (*PhotoCalibrationAnchor, error) {
	return photos.NewClient(c, jobID).UpdateCalibrationAnchor(ctx, photoID, anchorID, req, opts)
}

// DeletePhotoCalibrationAnchor deletes the specified calibration anchor (v3).
func (c *Client) DeletePhotoCalibrationAnchor(ctx context.Context, jobID, photoID, anchorID string) error {
	return photos.NewClient(c, jobID).DeleteCalibrationAnchor(ctx, photoID, anchorID)
}
