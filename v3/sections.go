package katapultpro

import (
	"context"
	"io"

	"github.com/romer-pro/katapultpro-go-sdk/v3/sections"
)

// ListSections returns all sections on the specified connection (v3).
func (c *Client) ListSections(ctx context.Context, jobID, connectionID string) ([]Section, error) {
	return sections.NewClient(c, jobID, connectionID).List(ctx)
}

// GetSection returns the specified section (v3).
func (c *Client) GetSection(ctx context.Context, jobID, connectionID, sectionKey string) (*Section, error) {
	return sections.NewClient(c, jobID, connectionID).Get(ctx, sectionKey)
}

// CreateSection creates a new section on the connection (v3).
func (c *Client) CreateSection(ctx context.Context, jobID, connectionID string, req *CreateSectionRequest) (*Section, error) {
	return sections.NewClient(c, jobID, connectionID).Create(ctx, req)
}

// UpdateSection updates the specified section, or creates it with the given key if it does not exist (v3).
func (c *Client) UpdateSection(ctx context.Context, jobID, connectionID, sectionKey string, req *UpdateSectionRequest, opts *UpdateSectionOptions) (*Section, error) {
	return sections.NewClient(c, jobID, connectionID).Update(ctx, sectionKey, req, opts)
}

// UploadSectionPhoto uploads a photo (image/jpeg) and associates it to the section (v3).
func (c *Client) UploadSectionPhoto(ctx context.Context, jobID, connectionID, sectionID string, imageData io.Reader, opts *UploadSectionPhotoOptions) (*Photo, error) {
	return sections.NewClient(c, jobID, connectionID).UploadPhoto(ctx, sectionID, imageData, opts)
}

// DeleteSection deletes the specified section (v3).
func (c *Client) DeleteSection(ctx context.Context, jobID, connectionID, sectionKey string) error {
	return sections.NewClient(c, jobID, connectionID).Delete(ctx, sectionKey)
}
