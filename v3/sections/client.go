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

// Section returns a SectionScope for the given section, enabling drill-down operations.
// Example: client.Job(jobID).Connections().Connection(connID).Sections().Section(sectionID).Get(ctx)
func (c *Client) Section(sectionKey string) *SectionScope {
	return &SectionScope{do: c.do, jobID: c.jobID, connectionID: c.connectionID, sectionKey: sectionKey}
}

// SectionScope scopes operations to a single section.
type SectionScope struct {
	do           request.Doer
	jobID        string
	connectionID string
	sectionKey   string
}

// SectionKey returns the scoped section key.
func (s *SectionScope) SectionKey() string { return s.sectionKey }

// Get returns the section.
func (s *SectionScope) Get(ctx context.Context) (*Section, error) {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID + "/sections/" + s.sectionKey
	var section Section
	if err := s.do.Do(ctx, http.MethodGet, path, nil, nil, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// Update updates the section. Use opts.OnlyIfExists to avoid creating with the given key.
func (s *SectionScope) Update(ctx context.Context, req *UpdateSectionRequest, opts *UpdateSectionOptions) (*Section, error) {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID + "/sections/" + s.sectionKey
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var section Section
	if err := s.do.Do(ctx, http.MethodPost, path, q, req, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// UploadPhoto uploads a photo (image/jpeg) and associates it to the section.
func (s *SectionScope) UploadPhoto(ctx context.Context, imageData io.Reader, opts *UploadSectionPhotoOptions) (*photos.Photo, error) {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID + "/sections/" + s.sectionKey + "/photos"
	var q url.Values
	if opts != nil && opts.AssociationValue.IsValid() {
		q = url.Values{}
		q.Set("association_value", opts.AssociationValue.String())
	}
	var photo photos.Photo
	if err := s.do.DoWithBody(ctx, http.MethodPost, path, q, "image/jpeg", imageData, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Delete deletes the section.
func (s *SectionScope) Delete(ctx context.Context) error {
	path := "v3/jobs/" + s.jobID + "/connections/" + s.connectionID + "/sections/" + s.sectionKey
	return s.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
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
