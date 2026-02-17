package sections

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// Client performs section API operations for a single connection. Create with NewClient(do, jobID, connectionID).
type Client struct {
	do           request.Doer
	jobID        string
	connectionID string
}

// NewClient returns a sections client for the given job and connection.
func NewClient(do request.Doer, jobID, connectionID string) *Client {
	return &Client{do: do, jobID: jobID, connectionID: connectionID}
}

// List returns all sections on the connection (v3).
func (c *Client) List(ctx context.Context) ([]Section, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections"
	var out []Section
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified section (v3).
func (c *Client) Get(ctx context.Context, sectionKey string) (*Section, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections/" + sectionKey
	var section Section
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// Create creates a new section on the connection (v3).
func (c *Client) Create(ctx context.Context, req *CreateSectionRequest) (*Section, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections"
	var section Section
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// Update updates the specified section, or creates it with the given key if it does not exist (v3).
func (c *Client) Update(ctx context.Context, sectionKey string, req *UpdateSectionRequest, opts *UpdateSectionOptions) (*Section, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections/" + sectionKey
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var section Section
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// UploadPhoto uploads a photo (image/jpeg) and associates it to the section (v3).
func (c *Client) UploadPhoto(ctx context.Context, sectionID string, imageData io.Reader, opts *UploadSectionPhotoOptions) (*photos.Photo, error) {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections/" + sectionID + "/photos"
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

// Delete deletes the specified section (v3).
func (c *Client) Delete(ctx context.Context, sectionKey string) error {
	path := "v3/jobs/" + c.jobID + "/connections/" + c.connectionID + "/sections/" + sectionKey
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}
